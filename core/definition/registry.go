// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package definition

import (
	"fmt"
)

const (
	// RegistryService const
	RegistryService = "registry"

	// RegistryPort const
	RegistryPort = "5000"

	// RegistryDockerImage const
	RegistryDockerImage = "registry"

	// RegistryDockerImageVersion const
	RegistryDockerImageVersion = "2"

	// RegistryRestartPolicy const
	RegistryRestartPolicy = "unless-stopped"
)

// GetRegistryConfig gets yaml definition object
func GetRegistryConfig(name, version string) DockerComposeConfig {
	services := make(map[string]Service)

	if version == "" {
		version = RegistryDockerImageVersion
	}

	services[name] = Service{
		Image:   fmt.Sprintf("%s:%s", RegistryDockerImage, version),
		Restart: RegistryRestartPolicy,
		Ports:   []string{RegistryPort},
	}

	return DockerComposeConfig{
		Version:  "3",
		Services: services,
	}
}
