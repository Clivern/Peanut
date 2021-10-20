// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package definition

import (
	"fmt"
)

const (
	// GhostService const
	GhostService = "ghost"

	// GhostPort const
	GhostPort = "2368"

	// GhostDockerImage const
	GhostDockerImage = "ghost"

	// GhostDockerImageVersion const
	GhostDockerImageVersion = "4.19.1"

	// GhostRestartPolicy const
	GhostRestartPolicy = "unless-stopped"
)

// GetGhostConfig gets yaml definition object
func GetGhostConfig(name, version string) DockerComposeConfig {
	services := make(map[string]Service)

	if version == "" {
		version = GhostDockerImageVersion
	}

	services[name] = Service{
		Image:   fmt.Sprintf("%s:%s", GhostDockerImage, version),
		Restart: GhostRestartPolicy,
		Ports:   []string{GhostPort},
	}

	return DockerComposeConfig{
		Version:  "3",
		Services: services,
	}
}
