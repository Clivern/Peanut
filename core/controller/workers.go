// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package controller

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/clivern/peanut/core/definition"
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
	Service       string            `json:"service"`
	Version       string            `json:"version"`
	Configs       map[string]string `json:"configs"`
	DeleteAfter   string            `json:"deleteAfter"`
	Type          string            `json:"type"`
	CorrelationID string            `json:"correlationID"`
}

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

	if viper.GetString("app.containerization.driver") == "docker" {
		result.containerization = runtime.NewDockerCompose()
	} else {
		panic("Invalid containerization runtime!")
	}

	return result
}

// DeployRequest sends a deploy request to workers
func (w *Workers) DeployRequest(c *gin.Context, rawBody []byte) {
	message := &Message{}

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

	allowed := []string{
		definition.RedisService,
		definition.EtcdService,
		definition.GrafanaService,
		definition.MariaDBService,
		definition.MySQLService,
		definition.ElasticSearchService,
		definition.GraphiteService,
		definition.PrometheusService,
		definition.ZipkinService,
		definition.MemcachedService,
		definition.MailhogService,
		definition.JaegerService,
		definition.PostgreSQLService,
		definition.MongoDBService,
		definition.RabbitMQService,
		definition.ConsulService,
		definition.VaultService,
	}

	defaultTags := map[string]string{
		definition.RedisService:         definition.RedisDockerImageVersion,
		definition.EtcdService:          definition.EtcdDockerImageVersion,
		definition.GrafanaService:       definition.GrafanaDockerImageVersion,
		definition.MariaDBService:       definition.MariaDBDockerImageVersion,
		definition.MySQLService:         definition.MySQLDockerImageVersion,
		definition.ElasticSearchService: definition.ElasticSearchDockerImageVersion,
		definition.GraphiteService:      definition.GraphiteDockerImageVersion,
		definition.PrometheusService:    definition.PrometheusDockerImageVersion,
		definition.ZipkinService:        definition.ZipkinDockerImageVersion,
		definition.MemcachedService:     definition.MemcachedDockerImageVersion,
		definition.MailhogService:       definition.MailhogDockerImageVersion,
		definition.JaegerService:        definition.JaegerDockerImageVersion,
		definition.PostgreSQLService:    definition.PostgreSQLDockerImageVersion,
		definition.MongoDBService:       definition.MongoDBDockerImageVersion,
		definition.RabbitMQService:      definition.RabbitMQDockerImageVersion,
		definition.ConsulService:        definition.VaultDockerImageVersion,
		definition.VaultService:         definition.ConsulDockerImageVersion,
	}

	if !util.InArray(message.Service, allowed) {
		c.JSON(http.StatusBadRequest, gin.H{
			"correlationID": c.GetHeader("x-correlation-id"),
			"errorMessage":  fmt.Sprintf("Error! Invalid service provided: %s", message.Service),
		})
		return
	}

	// Override version with the default one if not provided
	if message.Version == "" {
		message.Version = defaultTags[message.Service]
	}

	message.CorrelationID = c.GetHeader("x-correlation-id")
	message.JobID = util.GenerateUUID4()
	message.ServiceID = util.GenerateUUID4()
	message.Type = "DeployRequest"

	if message.Configs == nil {
		message.Configs = map[string]string{}
	}

	log.WithFields(log.Fields{
		"correlation_id": message.CorrelationID,
		"message":        message,
	}).Info(`Incoming request`)

	// Create a async job
	err = w.job.CreateRecord(model.JobRecord{
		ID:     message.JobID,
		Action: model.DeployJob,
		Service: model.ServiceRecord{
			ID:          message.ServiceID,
			Service:     message.Service,
			Configs:     message.Configs,
			DeleteAfter: message.DeleteAfter,
			Version:     message.Version,
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
		"service":   message.ServiceID,
		"type":      "service.deploy",
		"status":    model.PendingStatus,
		"createdAt": time.Now().UTC().Format("2006-01-02T15:04:05.000Z"),
	})
}

// DestroyRequest sends a destroy request to workers
func (w *Workers) DestroyRequest(c *gin.Context) {
	message := Message{
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

	message.Service = service.Service
	message.Configs = service.Configs
	message.DeleteAfter = service.DeleteAfter

	log.WithFields(log.Fields{
		"correlation_id": message.CorrelationID,
		"message":        message,
	}).Info(`Incoming request`)

	// Create a async job
	err = w.job.CreateRecord(model.JobRecord{
		ID:     message.JobID,
		Action: model.DestroyJob,
		Service: model.ServiceRecord{
			ID:          message.ServiceID,
			Service:     message.Service,
			Configs:     message.Configs,
			DeleteAfter: message.DeleteAfter,
			Version:     message.Version,
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
		"id":        message.JobID,
		"service":   message.ServiceID,
		"type":      "service.destroy",
		"status":    model.PendingStatus,
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
			"correlation_id": message.CorrelationID,
			"message":        message,
		}).Info(`Worker received a new message`)

		switch message.Type {
		case "DeployRequest":
			// Deploy the service
			result := make(map[string]string)
			result, err = w.containerization.Deploy(message.ServiceID, message.Service, message.Version, message.Configs)
			message.Configs = util.MergeMaps(message.Configs, result)
		case "DestroyRequest":
			// Destroy the service
			err = w.containerization.Destroy(message.ServiceID, message.Service, message.Version, message.Configs)
		default:
			log.WithFields(log.Fields{
				"correlation_id": message.CorrelationID,
				"type":           message.Type,
			}).Error(`Failed to find message type `)
		}

		// Update Job Status
		job, errr := w.job.GetRecord(message.ServiceID, message.JobID)

		if errr != nil {
			log.WithFields(log.Fields{
				"correlation_id": message.CorrelationID,
				"error":          errr.Error(),
			}).Error(`Worker failed to find the async job`)
			continue
		}

		if err == nil {
			job.Status = model.SuccessStatus
		} else {
			job.Status = model.FailedStatus

			log.WithFields(log.Fields{
				"correlation_id": message.CorrelationID,
				"message":        message,
				"error":          err.Error(),
			}).Error(`Worker failed to process message`)
		}

		w.job.UpdateRecord(job)

		log.WithFields(log.Fields{
			"correlation_id": message.CorrelationID,
			"message":        message,
		}).Info(`Worker finished processing the message`)

		notifyChannel <- message
	}

	wg.Done()
}

// Finalize finalizes a request
func (w *Workers) Finalize(notifyChannel <-chan Message) {
	for message := range notifyChannel {
		switch message.Type {
		case "DeployRequest":
			w.service.CreateRecord(model.ServiceRecord{
				ID:          message.ServiceID,
				Service:     message.Service,
				Configs:     message.Configs,
				DeleteAfter: message.DeleteAfter,
				Version:     message.Version,
			})
		case "DestroyRequest":
			w.service.DeleteRecord(message.ServiceID)
		}

		log.WithFields(log.Fields{
			"correlation_id": message.CorrelationID,
			"message":        message,
		}).Info(`Worker finalize processing`)
	}
}

// Watch watches for a pending jobs
func (w *Workers) Watch() {
	for {
		time.Sleep(20 * time.Second)

		data, err := w.service.GetRecords()

		if err != nil {
			continue
		}

		for _, v := range data {
			if v.DeleteAfter == "" {
				continue
			}

			if v.CreatedAt+int64(util.TimeInSec(v.DeleteAfter)) > time.Now().Unix() {
				continue
			}

			message := Message{
				ServiceID:     v.ID,
				JobID:         util.GenerateUUID4(),
				Type:          "DestroyRequest",
				Service:       v.Service,
				Configs:       v.Configs,
				DeleteAfter:   v.DeleteAfter,
				Version:       v.Version,
				CorrelationID: "",
			}

			w.job.CreateRecord(model.JobRecord{
				ID:     message.JobID,
				Action: model.DestroyJob,
				Service: model.ServiceRecord{
					ID:          message.ServiceID,
					Service:     message.Service,
					Configs:     message.Configs,
					DeleteAfter: message.DeleteAfter,
					Version:     message.Version,
				},
				Status: model.PendingStatus,
			})

			log.WithFields(log.Fields{
				"message": message,
			}).Info(`Destroy a service since due time reached`)

			w.broadcast <- message
		}
	}
}
