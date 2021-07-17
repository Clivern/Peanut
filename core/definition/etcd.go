// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package definition

import (
	"fmt"
)

const (
	// EtcdService const
	EtcdService = "etcd"

	// EtcdPort const
	EtcdPort = "2379"

	// EtcdDockerImage const
	EtcdDockerImage = "bitnami/etcd"

	// EtcdDockerImageVersion const
	EtcdDockerImageVersion = "3.5.0"

	// EtcdRestartPolicy const
	EtcdRestartPolicy = "unless-stopped"
)

// GetEtcdConfig gets yaml definition object
func GetEtcdConfig(name, version string) DockerComposeConfig {
	services := make(map[string]Service)

	if version == "" {
		version = EtcdDockerImageVersion
	}

	services[name] = Service{
		Image:       fmt.Sprintf("%s:%s", EtcdDockerImage, version),
		Restart:     EtcdRestartPolicy,
		Ports:       []string{EtcdPort},
		Environment: []string{"ALLOW_NONE_AUTHENTICATION=yes"},
	}

	return DockerComposeConfig{
		Version:  "3",
		Services: services,
	}
}
