// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package definition

import (
	"fmt"
)

const (
	// EtherpadService const
	EtherpadService = "etherpad"

	// EtherpadPort const
	EtherpadPort = "9001"

	// EtherpadDockerImage const
	EtherpadDockerImage = "etherpad/etherpad"

	// EtherpadDockerImageVersion const
	EtherpadDockerImageVersion = "1.8.14"

	// EtherpadRestartPolicy const
	EtherpadRestartPolicy = "unless-stopped"
)

// GetEtherpadConfig gets yaml definition object
func GetEtherpadConfig(name, version string) DockerComposeConfig {
	services := make(map[string]Service)

	if version == "" {
		version = EtherpadDockerImageVersion
	}

	services[name] = Service{
		Image:   fmt.Sprintf("%s:%s", EtherpadDockerImage, version),
		Restart: EtherpadRestartPolicy,
		Ports:   []string{EtherpadPort},
	}

	return DockerComposeConfig{
		Version:  "3",
		Services: services,
	}
}
