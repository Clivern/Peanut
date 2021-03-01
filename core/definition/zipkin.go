// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package definition

import (
	"fmt"
)

const (
	// ZipkinService const
	ZipkinService = "zipkin"

	// ZipkinPort const
	ZipkinPort = "9411"

	// ZipkinDockerImage const
	ZipkinDockerImage = "openzipkin/zipkin"

	// ZipkinDockerImageVersion const
	ZipkinDockerImageVersion = "2.23"

	// ZipkinRestartPolicy const
	ZipkinRestartPolicy = "unless-stopped"
)

// GetZipkinConfig gets yaml definition object
func GetZipkinConfig(name, version string) DockerComposeConfig {
	services := make(map[string]Service)

	if version == "" {
		version = ZipkinDockerImageVersion
	}

	services[name] = Service{
		Image:   fmt.Sprintf("%s:%s", ZipkinDockerImage, version),
		Restart: ZipkinRestartPolicy,
		Ports:   []string{ZipkinPort},
	}

	return DockerComposeConfig{
		Version:  "3",
		Services: services,
	}
}
