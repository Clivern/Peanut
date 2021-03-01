// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package definition

import (
	"fmt"
)

const (
	// RedisService const
	RedisService = "redis"

	// RedisPort const
	RedisPort = "6379"

	// RedisDockerImage const
	RedisDockerImage = "bitnami/redis"

	// RedisDockerImageVersion const
	RedisDockerImageVersion = "6.2.4"

	// RedisRestartPolicy const
	RedisRestartPolicy = "unless-stopped"

	// RedisDefaultPassword const
	RedisDefaultPassword = ""
)

// GetRedisConfig gets yaml definition object
func GetRedisConfig(name, version, password string) DockerComposeConfig {
	services := make(map[string]Service)

	envVar1 := "ALLOW_EMPTY_PASSWORD=yes"

	if password != "" {
		envVar1 = fmt.Sprintf("REDIS_PASSWORD=%s", password)
	}

	if version == "" {
		version = RedisDockerImageVersion
	}

	services[name] = Service{
		Image:   fmt.Sprintf("%s:%s", RedisDockerImage, version),
		Restart: RedisRestartPolicy,
		Ports:   []string{RedisPort},
		Environment: []string{
			envVar1,
		},
	}

	return DockerComposeConfig{
		Version:  "3",
		Services: services,
	}
}
