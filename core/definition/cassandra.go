// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package definition

import (
	"fmt"
)

const (
	// CassandraService const
	CassandraService = "cassandra"

	// CassandraPort const
	CassandraPort = "9042"

	// CassandraDockerImage const
	CassandraDockerImage = "cassandra"

	// CassandraDockerImageVersion const
	CassandraDockerImageVersion = "4.0"

	// CassandraRestartPolicy const
	CassandraRestartPolicy = "unless-stopped"
)

// GetCassandraConfig gets yaml definition object
func GetCassandraConfig(name, version string) DockerComposeConfig {
	services := make(map[string]Service)

	if version == "" {
		version = CassandraDockerImageVersion
	}

	services[name] = Service{
		Image:   fmt.Sprintf("%s:%s", CassandraDockerImage, version),
		Restart: CassandraRestartPolicy,
		Ports:   []string{CassandraPort},
	}

	return DockerComposeConfig{
		Version:  "3",
		Services: services,
	}
}
