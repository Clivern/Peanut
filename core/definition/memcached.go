// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package definition

const (
	// MemcachedService const
	MemcachedService = "memcached"

	// MemcachedPort const
	MemcachedPort = "11211"

	// MemcachedDockerImage const
	MemcachedDockerImage = "memcached:1.6.9"

	// MemcachedRestartPolicy const
	MemcachedRestartPolicy = "unless-stopped"
)

// GetMemcachedConfig gets yaml definition object
func GetMemcachedConfig(name string) DockerComposeConfig {
	services := make(map[string]Service)

	services[name] = Service{
		Image:   MemcachedDockerImage,
		Restart: MemcachedRestartPolicy,
		Ports:   []string{MemcachedPort},
	}

	return DockerComposeConfig{
		Version:  "3",
		Services: services,
	}
}
