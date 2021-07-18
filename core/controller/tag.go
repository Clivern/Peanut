// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package controller

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/clivern/peanut/core/driver"
	"github.com/clivern/peanut/core/model"
	"github.com/clivern/peanut/core/service"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// GetTags controller
func GetTags(c *gin.Context) {
	var err error

	image := c.Param("image")
	org := c.Param("org")
	fromCache := c.Param("fromCache")

	gitHub := service.NewDockerHub(service.NewHTTPClient(20))

	// If caching is deactivated, fetch tags from docker hub
	if viper.GetInt("app.containerization.cacheTagsTimeInMinutes") == 0 {
		log.WithFields(log.Fields{
			"correlation_id": c.GetHeader("x-correlation-id"),
		}).Info("Image tags cache is disabled")

		tags, err := gitHub.GetTags(context.Background(), org, image)

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
			"image":     image,
			"fromCache": fromCache,
			"tags":      tags,
		})
		return
	}

	db := driver.NewEtcdDriver()

	err = db.Connect()

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

	optionStore := model.NewOptionStore(db)

	optionData, err := optionStore.GetOptionByKey(fmt.Sprintf(
		"%s_%s_tags",
		org,
		image,
	))

	// If no cached tags
	if err != nil && strings.Contains(err.Error(), "Unable to find") {
		log.WithFields(log.Fields{
			"correlation_id": c.GetHeader("x-correlation-id"),
		}).Info("Store image tags")

		tags, err := gitHub.GetTags(context.Background(), org, image)

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

		err = optionStore.CreateOption(model.OptionData{
			Key:   fmt.Sprintf("%s_%s_tags", org, image),
			Value: strings.Join(tags, ";;"),
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

		c.JSON(http.StatusOK, gin.H{
			"image":     image,
			"fromCache": fromCache,
			"tags":      tags,
		})
		return
	}

	checkValue := optionData.UpdatedAt + int64(viper.GetInt("app.containerization.cacheTagsTimeInMinutes")*60)

	// If cache still valid
	if (time.Now().Unix() < checkValue) && fromCache == "true" {
		log.WithFields(log.Fields{
			"correlation_id": c.GetHeader("x-correlation-id"),
		}).Info("Cached image tags still valid")

		c.JSON(http.StatusOK, gin.H{
			"image":     image,
			"fromCache": fromCache,
			"tags":      strings.Split(optionData.Value, ";;"),
		})
		return
	}

	log.WithFields(log.Fields{
		"correlation_id": c.GetHeader("x-correlation-id"),
	}).Info("Cached image tags expired")

	tags, err := gitHub.GetTags(context.Background(), org, image)

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

	err = optionStore.UpdateOptionByKey(model.OptionData{
		Key:   fmt.Sprintf("%s_%s_tags", org, image),
		Value: strings.Join(tags, ";;"),
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

	c.JSON(http.StatusOK, gin.H{
		"image":     image,
		"fromCache": fromCache,
		"tags":      tags,
	})
}
