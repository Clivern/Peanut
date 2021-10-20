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

	"github.com/clivern/peanut/core/definition"
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
	serviceType := c.Param("serviceType")
	fromCache := c.Param("fromCache")

	dockerImagesOrgs := map[string][]string{
		definition.RedisService:         []string{"bitnami", "redis"},
		definition.EtcdService:          []string{"bitnami", "etcd"},
		definition.GrafanaService:       []string{"grafana", "grafana"},
		definition.MariaDBService:       []string{"library", "mariadb"},
		definition.MySQLService:         []string{"library", "mysql"},
		definition.ElasticSearchService: []string{"library", "elasticsearch"},
		definition.GraphiteService:      []string{"graphiteapp", "graphite-statsd"},
		definition.PrometheusService:    []string{"prom", "prometheus"},
		definition.ZipkinService:        []string{"openzipkin", "zipkin"},
		definition.MemcachedService:     []string{"library", "memcached"},
		definition.MailhogService:       []string{"mailhog", "mailhog"},
		definition.JaegerService:        []string{"jaegertracing", "all-in-one"},
		definition.PostgreSQLService:    []string{"library", "postgres"},
		definition.MongoDBService:       []string{"library", "mongo"},
		definition.RabbitMQService:      []string{"library", "rabbitmq"},
		definition.ConsulService:        []string{"library", "consul"},
		definition.VaultService:         []string{"library", "vault"},
		definition.CassandraService:     []string{"library", "cassandra"},
		definition.MinioService:         []string{"minio", "minio"},
		definition.RegistryService:      []string{"library", "registry"},
		definition.GhostService:         []string{"library", "ghost"},
	}

	if _, ok := dockerImagesOrgs[serviceType]; !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"correlationID": c.GetHeader("x-correlation-id"),
			"errorMessage":  fmt.Sprintf("Error! Invalid service provided: %s", serviceType),
		})
		return
	}

	gitHub := service.NewDockerHub(service.NewHTTPClient(20))

	// If caching is deactivated, fetch tags from docker hub
	if viper.GetInt("app.containerization.cacheTagsTimeInMinutes") == 0 {
		log.WithFields(log.Fields{
			"correlation_id": c.GetHeader("x-correlation-id"),
		}).Info("Image tags cache is disabled")

		tags, err := gitHub.GetTags(
			context.Background(),
			dockerImagesOrgs[serviceType][0],
			dockerImagesOrgs[serviceType][1],
		)

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
			"serviceType": serviceType,
			"fromCache":   fromCache,
			"tags":        tags,
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
		dockerImagesOrgs[serviceType][0],
		dockerImagesOrgs[serviceType][1],
	))

	// If no cached tags
	if err != nil && strings.Contains(err.Error(), "Unable to find") {
		log.WithFields(log.Fields{
			"correlation_id": c.GetHeader("x-correlation-id"),
		}).Info("Store image tags")

		tags, err := gitHub.GetTags(
			context.Background(),
			dockerImagesOrgs[serviceType][0],
			dockerImagesOrgs[serviceType][1],
		)

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
			Key: fmt.Sprintf(
				"%s_%s_tags",
				dockerImagesOrgs[serviceType][0],
				dockerImagesOrgs[serviceType][1],
			),
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
			"serviceType": serviceType,
			"fromCache":   fromCache,
			"tags":        tags,
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
			"serviceType": serviceType,
			"fromCache":   fromCache,
			"tags":        strings.Split(optionData.Value, ";;"),
		})
		return
	}

	log.WithFields(log.Fields{
		"correlation_id": c.GetHeader("x-correlation-id"),
	}).Info("Cached image tags expired")

	tags, err := gitHub.GetTags(
		context.Background(),
		dockerImagesOrgs[serviceType][0],
		dockerImagesOrgs[serviceType][1],
	)

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
		Key: fmt.Sprintf(
			"%s_%s_tags",
			dockerImagesOrgs[serviceType][0],
			dockerImagesOrgs[serviceType][1],
		),
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
		"serviceType": serviceType,
		"fromCache":   fromCache,
		"tags":        tags,
	})
}
