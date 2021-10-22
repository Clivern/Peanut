// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package definition

import (
	"fmt"
)

const (
	// HttpbinService const
	HttpbinService = "httpbin"

	// HttpbinPort const
	HttpbinPort = "80"

	// HttpbinDockerImage const
	HttpbinDockerImage = "kennethreitz/httpbin"

	// HttpbinDockerImageVersion const
	HttpbinDockerImageVersion = "latest"

	// HttpbinRestartPolicy const
	HttpbinRestartPolicy = "unless-stopped"
)

// GetHttpbinConfig gets yaml definition object
func GetHttpbinConfig(name, version string) DockerComposeConfig {
	services := make(map[string]Service)

	if version == "" {
		version = HttpbinDockerImageVersion
	}

	services[name] = Service{
		Image:   fmt.Sprintf("%s:%s", HttpbinDockerImage, version),
		Restart: HttpbinRestartPolicy,
		Ports:   []string{HttpbinPort},
	}

	return DockerComposeConfig{
		Version:  "3",
		Services: services,
	}
}
