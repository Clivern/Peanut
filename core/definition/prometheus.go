// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package definition

import (
	"fmt"
)

const (
	// PrometheusService const
	PrometheusService = "prometheus"

	// PrometheusPort const
	PrometheusPort = "9090"

	// PrometheusDockerImage const
	PrometheusDockerImage = "prom/prometheus"

	// PrometheusDockerImageVersion const
	PrometheusDockerImageVersion = "v2.28.1"

	// PrometheusRestartPolicy const
	PrometheusRestartPolicy = "unless-stopped"

	// PrometheusDefaultConfig const
	PrometheusDefaultConfig = "Z2xvYmFsOgogIGV2YWx1YXRpb25faW50ZXJ2YWw6IDE1cwogIHNjcmFwZV9pbnRlcnZhbDogMTVzCnJ1bGVfZmlsZXM6IH4Kc2NyYXBlX2NvbmZpZ3M6CiAgLQogICAgam9iX25hbWU6IHByb21ldGhldXMKICAgIHNjcmFwZV9pbnRlcnZhbDogNXMKICAgIHN0YXRpY19jb25maWdzOgogICAgICAtCiAgICAgICAgdGFyZ2V0czoKICAgICAgICAgIC0gImxvY2FsaG9zdDo5MDkwIg=="
)

// GetPrometheusConfig gets yaml definition object
func GetPrometheusConfig(name, version, configPath string) DockerComposeConfig {
	services := make(map[string]Service)

	if version == "" {
		version = PrometheusDockerImageVersion
	}

	services[name] = Service{
		Image:   fmt.Sprintf("%s:%s", PrometheusDockerImage, version),
		Restart: PrometheusRestartPolicy,
		Ports:   []string{PrometheusPort},
		Volumes: []string{
			fmt.Sprintf("%s:/etc/prometheus/prometheus.yml", configPath),
		},
	}

	return DockerComposeConfig{
		Version:  "3",
		Services: services,
	}
}
