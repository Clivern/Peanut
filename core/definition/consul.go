// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package definition

import (
	"fmt"
)

const (
	// ConsulService const
	ConsulService = "consul"

	// ConsulHTTPPort const
	ConsulHTTPPort = "8500"

	// ConsulDockerImage const
	ConsulDockerImage = "consul"

	// ConsulDockerImageVersion const
	ConsulDockerImageVersion = "1.9.7"

	// ConsulRestartPolicy const
	ConsulRestartPolicy = "unless-stopped"
)

// GetConsulConfig gets yaml definition object
func GetConsulConfig(name, version string) DockerComposeConfig {
	services := make(map[string]Service)

	if version == "" {
		version = ConsulDockerImageVersion
	}

	services[name] = Service{
		Image:   fmt.Sprintf("%s:%s", ConsulDockerImage, version),
		Restart: ConsulRestartPolicy,
		Ports:   []string{ConsulHTTPPort},
		Command: "consul agent -dev -client=0.0.0.0",
	}

	return DockerComposeConfig{
		Version:  "3",
		Services: services,
	}
}
