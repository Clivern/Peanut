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
	GrafanaDockerImage = "grafana/grafana"

	// GrafanaDockerImageVersion const
	GrafanaDockerImageVersion = "8.0.4"

	// GrafanaRestartPolicy const
	GrafanaRestartPolicy = "unless-stopped"

	// GrafanaDefaultUsername const
	GrafanaDefaultUsername = "admin"

	// GrafanaDefaultPassword const
	GrafanaDefaultPassword = "admin"

	// GrafanaDefaultAnonymousAccess const
	GrafanaDefaultAnonymousAccess = "true"

	// GrafanaDefaultAllowSignup const
	GrafanaDefaultAllowSignup = "false"
)

// GetGrafanaConfig gets yaml definition object
func GetGrafanaConfig(name, version, username, password, allowSignup, anonymousAccess string) DockerComposeConfig {
	services := make(map[string]Service)

	if username == "" {
		username = GrafanaDefaultUsername
	}

	if password == "" {
		password = GrafanaDefaultPassword
	}

	if allowSignup == "" {
		allowSignup = GrafanaDefaultAllowSignup
	}

	if anonymousAccess == "" {
		anonymousAccess = GrafanaDefaultAnonymousAccess
	}

	if version == "" {
		version = GrafanaDockerImageVersion
	}

	services[name] = Service{
		Image:   fmt.Sprintf("%s:%s", GrafanaDockerImage, version),
		Restart: GrafanaRestartPolicy,
		Ports:   []string{GrafanaPort},
		Environment: []string{
			fmt.Sprintf("GF_SECURITY_ADMIN_USER=%s", username),
			fmt.Sprintf("GF_SECURITY_ADMIN_PASSWORD=%s", password),
			fmt.Sprintf("GF_USERS_ALLOW_SIGN_UP=%s", allowSignup),
			fmt.Sprintf("GF_AUTH_ANONYMOUS_ENABLED=%s", anonymousAccess),
		},
	}

	return DockerComposeConfig{
		Version:  "3",
		Services: services,
	}
}
