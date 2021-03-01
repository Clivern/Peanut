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
func GetRedisConfig(image, port, restart string) *DockerComposeConfig {
	services := make(map[string]Service)
	volumes := make(map[string]string)

	services["redis"] = Service{
		Image:   image,
		Restart: restart,
		Ports: []string{
			fmt.Sprintf("%s:%s", RedisPort, port),
		},
	}

	return &DockerComposeConfig{
		Version:  "3",
		Services: services,
		Volumes:  volumes,
	}
}
