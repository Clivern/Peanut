// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package controller

import (
	"net/http"

	"github.com/clivern/peanut/core/driver"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// Ready controller
func Ready(c *gin.Context) {
	db := driver.NewEtcdDriver()

	err := db.Connect()

	if err != nil || !db.IsConnected() {
		log.WithFields(log.Fields{
			"correlation_id": c.Request.Header.Get("X-Correlation-ID"),
			"status":         "NotOk",
		}).Info(`Ready check`)

		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "NotOk",
		})

		return
	}

	defer db.Close()

	log.WithFields(log.Fields{
		"correlation_id": c.Request.Header.Get("X-Correlation-ID"),
		"status":         "ok",
	}).Info(`Ready check`)

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}
