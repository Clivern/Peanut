// Copyright 2020 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package controller

import (
	"net/http"
	"time"

	"github.com/clivern/peanut/core/driver"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// ServicePayload type
type ServicePayload struct {
	Ident string `json:"ident"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// GetServices controller
func GetServices(c *gin.Context) {
	var services []ServicePayload

	db := driver.NewEtcdDriver()

	err := db.Connect()

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

	defer db.Close()

	c.JSON(http.StatusOK, gin.H{
		"services": services,
	})
}

// GetService controller
func GetService(c *gin.Context) {

	db := driver.NewEtcdDriver()

	err := db.Connect()

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

	defer db.Close()

	c.JSON(http.StatusOK, gin.H{
		"id": 1,
	})
}

// DeleteService controller
func DeleteService(c *gin.Context) {
	db := driver.NewEtcdDriver()

	err := db.Connect()

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

	defer db.Close()

	c.Status(http.StatusNoContent)
	return
}
