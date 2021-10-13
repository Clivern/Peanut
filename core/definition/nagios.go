// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package definition

import (
	"fmt"
)

const (
	// NagiosService const
	NagiosService = "nagios"

	// NagiosPort const
	NagiosPort = "80"

	// NagiosDockerImage const
	NagiosDockerImage = "jasonrivers/nagios"

	// NagiosDockerImageVersion const
	NagiosDockerImageVersion = "latest"

	// NagiosRestartPolicy const
	NagiosRestartPolicy = "unless-stopped"

	// NagiosRootUser const
	NagiosRootUser = "nagiosadmin"

	// NagiosRootPassword const
	NagiosRootPassword = "nagios"
)

// GetNagiosConfig gets yaml definition object
func GetNagiosConfig(name, version string) DockerComposeConfig {
	services := make(map[string]Service)

	if version == "" {
		version = NagiosDockerImageVersion
	}

	services[name] = Service{
		Image:   fmt.Sprintf("%s:%s", NagiosDockerImage, version),
		Restart: NagiosRestartPolicy,
		Ports:   []string{NagiosPort},
	}

	return DockerComposeConfig{
		Version:  "3",
		Services: services,
	}
}
