// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package runtime

import (
	"fmt"

	"github.com/clivern/peanut/core/definition"
	"github.com/clivern/peanut/core/model"
	"github.com/clivern/peanut/core/util"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// DockerCompose type
type DockerCompose struct {
}

// NewDockerCompose creates a new instance
func NewDockerCompose() *DockerCompose {
	instance := new(DockerCompose)
	return instance
}

// Deploy deploys services
func (d *DockerCompose) Deploy(service model.ServiceRecord) error {
	// Deploy redis service
	if model.RedisService == service.Template {
		return d.deployRedis(service)
	}

	return nil
}

// Destroy destroys services
func (d *DockerCompose) Destroy(service model.ServiceRecord) error {
	return nil
}

// deployRedis deploys a redis
func (d *DockerCompose) deployRedis(service model.ServiceRecord) error {
	redis := definition.GetRedisConfig(
		service.GetConfig("image", RedisDockerImage),
		service.GetConfig("port", definition.RedisPort),
		service.GetConfig("restartPolicy", "unless-stopped"),
		service.GetConfig("password", ""),
	)

	result, err := redis.ToString()

	if err != nil {
		return err
	}

	err = util.StoreFile(
		fmt.Sprintf("%s/%s.yml", viper.GetString("app.storage.path"), service.ID),
		result,
	)

	if err != nil {
		return err
	}

	command := fmt.Sprintf(
		"docker-compose -f %s/%s.yml up -d --force-recreate",
		viper.GetString("app.storage.path"),
		service.ID,
	)

	stdout, stderr, err := util.Exec(command)

	log.WithFields(log.Fields{
		"command": command,
	}).Info("Run a shell command")

	if err != nil {
		return err
	}

	// Store runtime verbose logs only in dev environment
	if viper.GetString("app.mode") == "dev" {
		err = util.StoreFile(
			fmt.Sprintf("%s/%s.stdout.log", viper.GetString("app.storage.path"), service.ID),
			stdout,
		)

		if err != nil {
			return err
		}

		err = util.StoreFile(
			fmt.Sprintf("%s/%s.stderr.log", viper.GetString("app.storage.path"), service.ID),
			stderr,
		)

		if err != nil {
			return err
		}
	}

	return nil
}
