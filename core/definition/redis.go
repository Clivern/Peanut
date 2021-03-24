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

	// RedisData const
	RedisData = "/data"
)

// GetRedis gets yaml definition object
func GetRedis(image, volume, port, restart string) *DockerComposeConfig {
	services := make(map[string]Service)
	volumes := make(map[string]string)

	services["redis"] = Service{
		Image:   image,
		Restart: restart,
		Volumes: []string{
			fmt.Sprintf("%s:%s", volume, RedisData),
		},
		Ports: []string{
			fmt.Sprintf("%s:%s", RedisPort, port),
		},
	}

	volumes[volume] = ""

	return &DockerComposeConfig{
		Version:  "3",
		Services: services,
		Volumes:  volumes,
	}
}
