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
)

// GetJob controller
func GetJob(c *gin.Context) {
	jobID := c.Param("jobId")
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

	jobStore := model.NewJobStore(db)

	jobData, err := jobStore.GetRecord(serviceID, jobID)

	if err != nil && strings.Contains(err.Error(), "Unable to find") {
		c.JSON(http.StatusNotFound, gin.H{
			"correlationID": c.GetHeader("x-correlation-id"),
			"errorMessage":  fmt.Sprintf("Unable to find job: %s", jobID),
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

	c.JSON(http.StatusOK, gin.H{
		"id":        jobData.ID,
		"action":    jobData.Action,
		"status":    jobData.Status,
		"createdAt": time.Unix(jobData.CreatedAt, 0),
		"updatedAt": time.Unix(jobData.UpdatedAt, 0),
	})
}
