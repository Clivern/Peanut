// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package controller

import (
	"net/http"
	"sync"

	"github.com/clivern/peanut/core/util"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// Message type
type Message struct {
	Payload       string `json:"payload"`
	CorrelationID string `json:"CorrelationID"`
}

// Workers type
type Workers struct {
	broadcast chan Message
}

// NewWorkers get a new workers instance
func NewWorkers() *Workers {
	result := new(Workers)
	result.broadcast = make(chan Message, viper.GetInt("app.workers.buffer"))

	return result
}

// BroadcastRequest sends a request to workers
func (w *Workers) BroadcastRequest(c *gin.Context, rawBody []byte) {
	message := &Message{}

	err := util.LoadFromJSON(message, rawBody)

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

	log.WithFields(log.Fields{
		"correlation_id": message.CorrelationID,
		"message":        message,
	}).Info(`Incoming request`)

	w.broadcast <- *message

	c.Status(http.StatusAccepted)
	return
}

// HandleWorkload handles all incoming requests
func (w *Workers) HandleWorkload() <-chan Message {
	notifyChannel := make(chan Message)

	go func() {
		wg := &sync.WaitGroup{}

		for t := 0; t < viper.GetInt("app.workers.count"); t++ {
			wg.Add(1)
			go w.ProcessAction(notifyChannel, wg)
		}

		wg.Wait()

		close(notifyChannel)
	}()

	return notifyChannel
}

// ProcessAction process incoming request
func (w *Workers) ProcessAction(notifyChannel chan<- Message, wg *sync.WaitGroup) {
	for message := range w.broadcast {
		log.WithFields(log.Fields{
			"correlation_id": message.CorrelationID,
			"message":        message,
		}).Info(`Worker received a new message`)

		// Process message

		log.WithFields(log.Fields{
			"correlation_id": message.CorrelationID,
			"message":        message,
		}).Info(`Worker finished processing`)

		notifyChannel <- message
	}

	wg.Done()
}

// Finalize finalizes a request
func (w *Workers) Finalize(notifyChannel <-chan Message) {
	for message := range notifyChannel {
		log.WithFields(log.Fields{
			"correlation_id": message.CorrelationID,
			"message":        message,
		}).Info(`Worker finalize processing`)
	}
}
