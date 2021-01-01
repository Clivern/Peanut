// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package definition

import (
	"fmt"
)

const (
	// RedisPort const
	RedisPort = "6379"
)

// GetRedisConfig gets yaml definition object
func GetRedisConfig(name, image, port, restart, password string) *DockerComposeConfig {
	services := make(map[string]Service)
	volumes := make(map[string]string)

	envVar1 := "ALLOW_EMPTY_PASSWORD=yes"

	if password != "" {
		envVar1 = fmt.Sprintf("REDIS_PASSWORD=%s", password)
	}

	services[name] = Service{
		Image:   image,
		Restart: restart,
		Ports:   []string{port},
		Environment: []string{
			envVar1,
		},
	}

	return &DockerComposeConfig{
		Version:  "3",
		Services: services,
		Volumes:  volumes,
	}
}
