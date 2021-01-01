// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package controller

import (
	"net/http"
	"sync"
	"time"

	"github.com/clivern/peanut/core/driver"
	"github.com/clivern/peanut/core/model"
	"github.com/clivern/peanut/core/runtime"
	"github.com/clivern/peanut/core/util"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// Workers type
type Workers struct {
	job              *model.Job
	service          *model.Service
	broadcast        chan Message
	containerization runtime.Containerization
}

// NewWorkers get a new workers instance
func NewWorkers() *Workers {
	result := new(Workers)

	db := driver.NewEtcdDriver()

	err := db.Connect()

	if err != nil || !db.IsConnected() {
		panic("Error while connecting to DB")
	}

	result.job = model.NewJobStore(db)
	result.service = model.NewServiceStore(db)
	result.broadcast = make(chan Message, viper.GetInt("app.workers.buffer"))

	if viper.GetString("app.containerization") == "docker" {
		result.containerization = runtime.NewDockerCompose()
	} else {
		panic("Invalid containerization runtime!")
	}

	return result
}

// DeployRequest sends a deploy request to workers
func (w *Workers) DeployRequest(c *gin.Context, rawBody []byte) {
	message := &DeployRequest{}

	err := util.LoadFromJSON(message, rawBody)

	if err != nil {
		log.WithFields(log.Fields{
			"error": err.Error(),
		}).Debug(`Invalid message`)

		c.JSON(http.StatusBadRequest, gin.H{
			"correlationID": c.GetHeader("x-correlation-id"),
			"errorMessage":  "Error! Invalid request",
		})
		return
	}

	message.CorrelationID = c.GetHeader("x-correlation-id")
	message.JobID = util.GenerateUUID4()
	message.ServiceID = util.GenerateUUID4()
	message.Type = "DeployRequest"

	log.WithFields(log.Fields{
		"correlation_id": message.GetCorrelation(),
		"message":        message,
	}).Info(`Incoming request`)

	// Create a async job
	err = w.job.CreateRecord(model.JobRecord{
		ID:     message.JobID,
		Action: model.DeployJob,
		Service: model.ServiceRecord{
			ID:          message.ServiceID,
			Template:    message.Template,
			Configs:     message.Configs,
			DeleteAfter: message.DeleteAfter,
		},
		Status: model.PendingStatus,
	})

	if err != nil {
		log.WithFields(log.Fields{
			"correlation_id": c.GetHeader("x-correlation-id"),
			"error":          err.Error(),
		}).Error("Internal server error")

		c.JSON(http.StatusInternalServerError, gin.H{
			"correlationID": c.GetHeader("x-correlation-id"),
			"errorMessage":  "Internal server error",
		})
		return
	}

	w.broadcast <- *message

	c.JSON(http.StatusAccepted, gin.H{
		"id":        message.GetJob(),
		"type":      "service.deploy",
		"status":    "PENDING",
		"createdAt": time.Now().UTC().Format("2006-01-02T15:04:05.000Z"),
	})
}

// DestroyRequest sends a destroy request to workers
func (w *Workers) DestroyRequest(c *gin.Context, rawBody []byte) {
	message := DestroyRequest{
		CorrelationID: c.GetHeader("x-correlation-id"),
		ServiceID:     c.Param("serviceId"),
		JobID:         util.GenerateUUID4(),
		Type:          "DestroyRequest",
	}

	service, err := w.service.GetRecord(c.Param("serviceId"))

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"correlationID": c.GetHeader("x-correlation-id"),
			"errorMessage":  "Error! service not found",
		})
		return
	}

	message.Template = service.Template
	message.Configs = service.Configs
	message.DeleteAfter = service.DeleteAfter

	log.WithFields(log.Fields{
		"correlation_id": message.GetCorrelation(),
		"message":        message,
	}).Info(`Incoming request`)

	// Create a async job
	err = w.job.CreateRecord(model.JobRecord{
		ID:     message.JobID,
		Action: model.DestroyJob,
		Service: model.ServiceRecord{
			ID: message.ServiceID,
		},
		Status: model.PendingStatus,
	})

	if err != nil {
		log.WithFields(log.Fields{
			"correlation_id": c.GetHeader("x-correlation-id"),
			"error":          err.Error(),
		}).Error("Internal server error")

		c.JSON(http.StatusInternalServerError, gin.H{
			"correlationID": c.GetHeader("x-correlation-id"),
			"errorMessage":  "Internal server error",
		})
		return
	}

	w.broadcast <- message

	c.JSON(http.StatusAccepted, gin.H{
		"id":        message.GetJob(),
		"type":      "service.destroy",
		"status":    "PENDING",
		"createdAt": time.Now().UTC().Format("2006-01-02T15:04:05.000Z"),
	})
}

// HandleWorkload handles all incoming requests
func (w *Workers) HandleWorkload() <-chan Message {
	notifyChannel := make(chan Message)

	go func() {
		wg := &sync.WaitGroup{}

		for t := 0; t < viper.GetInt("app.workers.count"); t++ {
			wg.Add(1)
			go w.ProcessRequest(notifyChannel, wg)
		}

		wg.Wait()

		close(notifyChannel)
	}()

	return notifyChannel
}

// ProcessRequest process incoming request
func (w *Workers) ProcessRequest(notifyChannel chan<- Message, wg *sync.WaitGroup) {
	var err error

	for message := range w.broadcast {
		log.WithFields(log.Fields{
			"correlation_id": message.GetCorrelation(),
			"message":        message,
		}).Info(`Worker received a new message`)

		switch message.GetType() {
		case "DeployRequest":
			// Deploy the service
			depr := DeployRequest{}
			result := make(map[string]string)

			result, err = w.containerization.Deploy(model.ServiceRecord{
				ID:          message.(DeployRequest).ServiceID,
				Template:    message.(DeployRequest).Template,
				Configs:     message.(DeployRequest).Configs,
				DeleteAfter: message.(DeployRequest).DeleteAfter,
			})

			// Override configs
			depr = message.(DeployRequest)
			depr.Configs = result
			message = depr
		case "DestroyRequest":
			// Destroy the service
			err = w.containerization.Destroy(model.ServiceRecord{
				ID:          message.(DestroyRequest).ServiceID,
				Template:    message.(DestroyRequest).Template,
				Configs:     message.(DestroyRequest).Configs,
				DeleteAfter: message.(DestroyRequest).DeleteAfter,
			})
		default:
			log.WithFields(log.Fields{
				"correlation_id": message.GetCorrelation(),
				"type":           message.GetType(),
			}).Error(`Failed to find message type `)
		}

		// Update Job Status
		job, errr := w.job.GetRecord(message.GetService(), message.GetJob())

		if errr != nil {
			log.WithFields(log.Fields{
				"correlation_id": message.GetCorrelation(),
				"error":          errr.Error(),
			}).Error(`Worker failed to find the async job`)
			continue
		}

		if err == nil {
			job.Status = model.SuccessStatus
		} else {
			job.Status = model.FailedStatus

			log.WithFields(log.Fields{
				"correlation_id": message.GetCorrelation(),
				"message":        message,
				"error":          err.Error(),
			}).Error(`Worker failed to process message`)
		}

		w.job.UpdateRecord(*job)

		log.WithFields(log.Fields{
			"correlation_id": message.GetCorrelation(),
			"message":        message,
		}).Info(`Worker finished processing the message`)

		notifyChannel <- message
	}

	wg.Done()
}

// Finalize finalizes a request
func (w *Workers) Finalize(notifyChannel <-chan Message) {
	for message := range notifyChannel {
		switch message.GetType() {
		case "DeployRequest":
			w.service.CreateRecord(model.ServiceRecord{
				ID:          message.GetService(),
				Template:    message.(DeployRequest).Template,
				Configs:     message.(DeployRequest).Configs,
				DeleteAfter: message.(DeployRequest).DeleteAfter,
			})
		case "DestroyRequest":
			// Delete service if no error raised
			w.service.DeleteRecord(message.GetService())
		}

		log.WithFields(log.Fields{
			"correlation_id": message.GetCorrelation(),
			"message":        message,
		}).Info(`Worker finalize processing`)
	}
}

// Watch watches for a pending jobs
func (w *Workers) Watch() {
	for {
		time.Sleep(5 * time.Second)
	}
}