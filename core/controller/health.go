// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// Health controller
func Health(c *gin.Context) {

	log.WithFields(log.Fields{
		"correlation_id": c.Request.Header.Get("X-Correlation-ID"),
		"status":         "ok",
	}).Info(`Health check`)

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}
