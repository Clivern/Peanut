// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package definition

import (
	"fmt"
)

const (
	// GrafanaService const
	GrafanaService = "grafana"

	// GrafanaPort const
	GrafanaPort = "3000"

	// GrafanaDockerImage const
	GrafanaDockerImage = "grafana/grafana:8.0.4"

	// GrafanaRestartPolicy const
	GrafanaRestartPolicy = "unless-stopped"

	// GrafanaDefaultUsername const
	GrafanaDefaultUsername = "admin"

	// GrafanaDefaultPassword const
	GrafanaDefaultPassword = "admin"
)

// GetGrafanaConfig gets yaml definition object
func GetGrafanaConfig(name, username, password string) DockerComposeConfig {
	services := make(map[string]Service)

	if username == "" {
		username = GrafanaDefaultUsername
	}

	if password == "" {
		password = GrafanaDefaultPassword
	}

	services[name] = Service{
		Image:   GrafanaDockerImage,
		Restart: GrafanaRestartPolicy,
		Ports:   []string{GrafanaPort},
		Environment: []string{
			fmt.Sprintf("GF_SECURITY_ADMIN_USER=%s", username),
			fmt.Sprintf("GF_SECURITY_ADMIN_PASSWORD=%s", password),
			"GF_USERS_ALLOW_SIGN_UP=false",
		},
	}

	return DockerComposeConfig{
		Version:  "3",
		Services: services,
	}
}
