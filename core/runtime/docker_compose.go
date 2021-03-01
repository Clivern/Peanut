// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package runtime

import (
	"fmt"

	"github.com/clivern/peanut/core/definition"
	"github.com/clivern/peanut/core/model"
	"github.com/clivern/peanut/core/util"
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
	// If redis
	if model.RedisService == service.Template {
		redis := definition.GetRedisConfig(
			"redis:5.0.10-alpine",
			definition.RedisPort,
			"unless-stopped",
		)

		result, _ := redis.ToString()
		util.StoreFile(fmt.Sprintf("/tmp/%s.yml", service.ID), result)

		stdout, stderr, _ := util.Exec(fmt.Sprintf("docker-compose -f /tmp/%s.yml", service.ID))
		util.StoreFile(fmt.Sprintf("/tmp/%s.stdout.log", service.ID), stdout)
		util.StoreFile(fmt.Sprintf("/tmp/%s.stderr.log", service.ID), stderr)
	}

	return nil
}

// Destroy destroys services
func (d *DockerCompose) Destroy(service *model.ServiceRecord) error {
	return nil
}
