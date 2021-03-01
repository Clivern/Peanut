// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package definition

import (
	"fmt"
)

const (
	// MemcachedService const
	MemcachedService = "memcached"

	// MemcachedPort const
	MemcachedPort = "11211"

	// MemcachedDockerImage const
	MemcachedDockerImage = "memcached"

	// MemcachedDockerImageVersion const
	MemcachedDockerImageVersion = "1.6.9"

	// MemcachedRestartPolicy const
	MemcachedRestartPolicy = "unless-stopped"
)

// GetMemcachedConfig gets yaml definition object
func GetMemcachedConfig(name, version string) DockerComposeConfig {
	services := make(map[string]Service)

	if version == "" {
		version = MemcachedDockerImageVersion
	}

	services[name] = Service{
		Image:   fmt.Sprintf("%s:%s", MemcachedDockerImage, version),
		Restart: MemcachedRestartPolicy,
		Ports:   []string{MemcachedPort},
	}

	return DockerComposeConfig{
		Version:  "3",
		Services: services,
	}
}
