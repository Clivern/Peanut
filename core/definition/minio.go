// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package definition

import (
	"fmt"
)

const (
	// MinioService const
	MinioService = "minio"

	// MinioAPIPort const
	MinioAPIPort = "9000"

	// MinioConsolePort const
	MinioConsolePort = "9001"

	// MinioDockerImage const
	MinioDockerImage = "minio/minio"

	// MinioDockerImageVersion const
	MinioDockerImageVersion = "RELEASE.2021-07-27T02-40-15Z"

	// MinioRestartPolicy const
	MinioRestartPolicy = "unless-stopped"

	// MinioRootUser const
	MinioRootUser = "admin12345678"

	// MinioRootPassword const
	MinioRootPassword = "admin12345678"
)

// GetMinioConfig gets yaml definition object
func GetMinioConfig(name, version, username, password string) DockerComposeConfig {
	services := make(map[string]Service)

	if version == "" {
		version = MinioDockerImageVersion
	}

	if username == "" {
		username = MinioRootUser
	}

	if password == "" {
		password = MinioRootPassword
	}

	services[name] = Service{
		Image:   fmt.Sprintf("%s:%s", MinioDockerImage, version),
		Command: "server /data --console-address :9001",
		Restart: MinioRestartPolicy,
		Ports: []string{
			MinioAPIPort,
			MinioConsolePort,
		},
		Environment: []string{
			fmt.Sprintf("MINIO_ROOT_USER=%s", username),
			fmt.Sprintf("MINIO_ROOT_PASSWORD=%s", password),
		},
	}

	return DockerComposeConfig{
		Version:  "3",
		Services: services,
	}
}
