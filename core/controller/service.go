// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package controller

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/clivern/peanut/core/driver"
	"github.com/clivern/peanut/core/model"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// ServicePayload type
type ServicePayload struct {
	ID          string            `json:"id"`
	Service     string            `json:"service"`
	Configs     map[string]string `json:"configs"`
	DeleteAfter string            `json:"deleteAfter"`
	CreatedAt   time.Time         `json:"createdAt"`
	UpdatedAt   time.Time         `json:"updatedAt"`
}

// GetServices controller
func GetServices(c *gin.Context) {
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

	serviceStore := model.NewServiceStore(db)

	data, err := serviceStore.GetRecords()

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

	var services []ServicePayload

	for _, v := range data {
		v.Configs["address"] = viper.GetString("app.hostname")

		services = append(services, ServicePayload{
			ID:          v.ID,
			Service:     v.Service,
			Configs:     v.Configs,
			DeleteAfter: v.DeleteAfter,
			CreatedAt:   time.Unix(v.CreatedAt, 0),
			UpdatedAt:   time.Unix(v.UpdatedAt, 0),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"services": services,
	})
}

// GetService controller
func GetService(c *gin.Context) {
	serviceID := c.Param("serviceId")

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

	serviceStore := model.NewServiceStore(db)

	serviceData, err := serviceStore.GetRecord(serviceID)

	if err != nil && strings.Contains(err.Error(), "Unable to find") {
		c.JSON(http.StatusNotFound, gin.H{
			"correlationID": c.GetHeader("x-correlation-id"),
			"errorMessage":  fmt.Sprintf("Unable to find job: %s", serviceID),
		})
		return
	}

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

	serviceData.Configs["address"] = viper.GetString("app.hostname")

	c.JSON(http.StatusOK, gin.H{
		"id":          serviceData.ID,
		"service":     serviceData.Service,
		"configs":     serviceData.Configs,
		"deleteAfter": serviceData.DeleteAfter,
		"createdAt":   time.Unix(serviceData.CreatedAt, 0),
		"updatedAt":   time.Unix(serviceData.UpdatedAt, 0),
	})
}
