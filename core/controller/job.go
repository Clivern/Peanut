// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetJobs controller
func GetJobs(c *gin.Context) {
	c.Status(http.StatusOK)
	return
}

// GetJob controller
func GetJob(c *gin.Context) {
	c.Status(http.StatusOK)
	return
}

// DeleteJob controller
func DeleteJob(c *gin.Context) {
	c.Status(http.StatusOK)
	return
}
