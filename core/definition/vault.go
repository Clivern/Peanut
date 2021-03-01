// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package definition

import (
	"fmt"
)

const (
	// VaultService const
	VaultService = "vault"

	// VaultHTTPPort const
	VaultHTTPPort = "8200"

	// VaultDockerImage const
	VaultDockerImage = "vault"

	// VaultDockerImageVersion const
	VaultDockerImageVersion = "1.7.3"

	// VaultRestartPolicy const
	VaultRestartPolicy = "unless-stopped"

	// VaultDefaultToken const
	VaultDefaultToken = "peanut"
)

// GetVaultConfig gets yaml definition object
func GetVaultConfig(name, version, token string) DockerComposeConfig {
	services := make(map[string]Service)

	if token == "" {
		token = VaultDefaultToken
	}

	if version == "" {
		version = VaultDockerImageVersion
	}

	services[name] = Service{
		Image:       fmt.Sprintf("%s:%s", VaultDockerImage, version),
		Restart:     VaultRestartPolicy,
		Ports:       []string{VaultHTTPPort},
		Environment: []string{fmt.Sprintf("VAULT_DEV_LISTEN_ADDRESS=0.0.0.0:%s", VaultHTTPPort)},
		Command:     fmt.Sprintf("vault server -dev -dev-root-token-id=%s", token),
	}

	return DockerComposeConfig{
		Version:  "3",
		Services: services,
	}
}
