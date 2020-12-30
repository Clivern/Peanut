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

// Message type
type Message struct {
	JobID         string            `json:"jobId"`
	ServiceID     string            `json:"serviceId"`
	Template      string            `json:"template"`
	Configs       map[string]string `json:"configs"`
	DeleteAfter   string            `json:"deleteAfter"`
	CorrelationID string            `json:"correlationID"`
}

// Workers type
type Workers struct {
	job              *model.Job
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
	result.broadcast = make(chan Message, viper.GetInt("app.workers.buffer"))

	if viper.GetString("app.containerization") == "docker_compose" {
		result.containerization = runtime.NewDockerCompose()
	} else {
		panic("Invalid containerization runtime!")
	}

	return result
}

// BroadcastRequest sends a request to workers
func (w *Workers) BroadcastRequest(c *gin.Context, rawBody []byte) {
	message := &Message{}

	err := util.LoadFromJSON(message, rawBody)

	message.CorrelationID = c.GetHeader("x-correlation-id")

	if err != nil {
		log.WithFields(log.Fields{
			"error": err.Error(),
		}).Error(`Invalid message`)

		c.JSON(http.StatusBadRequest, gin.H{
			"correlationID": c.GetHeader("x-correlation-id"),
			"errorMessage":  "Error! Invalid request",
		})
		return
	}

	message.JobID = util.GenerateUUID4()
	message.ServiceID = util.GenerateUUID4()

	log.WithFields(log.Fields{
		"correlation_id": message.CorrelationID,
		"message":        message,
	}).Info(`Incoming request`)

	// Create a async job
	err = w.job.CreateRecord(model.JobRecord{
		ID: message.JobID,
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
		"id":        message.JobID,
		"type":      "service.provision",
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
			go w.DeployService(notifyChannel, wg)
		}

		wg.Wait()

		close(notifyChannel)
	}()

	return notifyChannel
}

// DeployService process incoming request
func (w *Workers) DeployService(notifyChannel chan<- Message, wg *sync.WaitGroup) {
	for message := range w.broadcast {
		log.WithFields(log.Fields{
			"correlation_id": message.CorrelationID,
			"message":        message,
		}).Info(`Worker received a new message`)

		// Deploy the service
		w.containerization.Deploy(model.ServiceRecord{
			ID:          message.ServiceID,
			Template:    message.Template,
			Configs:     message.Configs,
			DeleteAfter: message.DeleteAfter,
		})

		log.WithFields(log.Fields{
			"correlation_id": message.CorrelationID,
			"message":        message,
		}).Info(`Worker finished deploying the service`)

		notifyChannel <- message
	}

	wg.Done()
}

// Finalize finalizes a request
func (w *Workers) Finalize(notifyChannel <-chan Message) {
	for message := range notifyChannel {
		// Store the service data
		log.WithFields(log.Fields{
			"correlation_id": message.CorrelationID,
			"message":        message,
		}).Info(`Worker finalize processing`)
	}
}
