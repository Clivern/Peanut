// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package definition

import (
	"fmt"
)

const (
	// KumaService const
	KumaService = "kuma"

	// KumaPort const
	KumaPort = "3001"

	// KumaDockerImage const
	KumaDockerImage = "louislam/uptime-kuma"

	// KumaDockerImageVersion const
	KumaDockerImageVersion = "1"

	// KumaRestartPolicy const
	KumaRestartPolicy = "unless-stopped"
)

// GetKumaConfig gets yaml definition object
func GetKumaConfig(name, version string) DockerComposeConfig {
	services := make(map[string]Service)

	if version == "" {
		version = KumaDockerImageVersion
	}

	services[name] = Service{
		Image:   fmt.Sprintf("%s:%s", KumaDockerImage, version),
		Restart: KumaRestartPolicy,
		Ports:   []string{KumaPort},
	}

	return DockerComposeConfig{
		Version:  "3",
		Services: services,
	}
}
