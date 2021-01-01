// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package runtime

import (
	"fmt"
	"strconv"

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
func (d *DockerCompose) Deploy(service model.ServiceRecord) (map[string]string, error) {
	data := make(map[string]string)

	// Deploy redis service
	if model.RedisService == service.Template {
		return d.deployRedis(service)
	}

	return data, fmt.Errorf("Error! Undefined service")
}

// Destroy destroys services
func (d *DockerCompose) Destroy(service model.ServiceRecord) error {
	if model.RedisService == service.Template {
		return d.destroyRedis(service)
	}

	return nil
}

// deployRedis deploys a redis
func (d *DockerCompose) deployRedis(service model.ServiceRecord) (map[string]string, error) {
	data := make(map[string]string)

	if service.Configs != nil {
		data = service.Configs
	}

	data["address"] = "[NodeIp]"
	data["port"] = service.GetConfig("port", strconv.Itoa(util.Rand(
		definition.RedisPortStart,
		definition.RedisPortEnd,
	)))

	data["password"] = service.GetConfig("password", "")

	redis := definition.GetRedisConfig(
		service.GetConfig("image", RedisDockerImage),
		data["port"],
		service.GetConfig("restartPolicy", "unless-stopped"),
		data["password"],
	)

	result, err := redis.ToString()

	if err != nil {
		return data, err
	}

	err = util.StoreFile(
		fmt.Sprintf("%s/%s.yml", viper.GetString("app.storage.path"), service.ID),
		result,
	)

	if err != nil {
		return data, err
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
		return data, err
	}

	// Store runtime verbose logs only in dev environment
	if viper.GetString("app.mode") == "dev" {
		err = util.StoreFile(
			fmt.Sprintf("%s/%s.stdout.log", viper.GetString("app.storage.path"), service.ID),
			stdout,
		)

		if err != nil {
			return data, err
		}

		err = util.StoreFile(
			fmt.Sprintf("%s/%s.stderr.log", viper.GetString("app.storage.path"), service.ID),
			stderr,
		)

		if err != nil {
			return data, err
		}
	}

	return data, nil
}

// destroyRedis destroys redis
func (d *DockerCompose) destroyRedis(service model.ServiceRecord) error {
	redis := definition.GetRedisConfig(
		service.GetConfig("image", RedisDockerImage),
		service.GetConfig("port", ""),
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
		"docker-compose -f %s/%s.yml down --rmi all --remove-orphans",
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
