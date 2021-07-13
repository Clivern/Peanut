// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package definition

import (
	"fmt"
)

const (
	// PostgreSQLService const
	PostgreSQLService = "postgresql"

	// PostgreSQLPort const
	PostgreSQLPort = "5432"

	// PostgreSQLDockerImage const
	PostgreSQLDockerImage = "postgres:13.3"

	// PostgreSQLRestartPolicy const
	PostgreSQLRestartPolicy = "unless-stopped"

	// PostgreSQLDefaultDatabase const
	PostgreSQLDefaultDatabase = "peanut"

	// PostgreSQLDefaultUsername const
	PostgreSQLDefaultUsername = "peanut"

	// PostgreSQLDefaultPassword const
	PostgreSQLDefaultPassword = "secret"
)

// GetPostgreSQLConfig gets yaml definition object
func GetPostgreSQLConfig(name, database, username, password string) DockerComposeConfig {
	services := make(map[string]Service)

	if database == "" {
		database = PostgreSQLDefaultDatabase
	}

	if username == "" {
		username = PostgreSQLDefaultUsername
	}

	if password == "" {
		password = PostgreSQLDefaultPassword
	}

	services[name] = Service{
		Image:   PostgreSQLDockerImage,
		Restart: PostgreSQLRestartPolicy,
		Ports:   []string{PostgreSQLPort},
		Environment: []string{
			fmt.Sprintf("POSTGRES_DB=%s", database),
			fmt.Sprintf("POSTGRES_USER=%s", username),
			fmt.Sprintf("POSTGRES_PASSWORD=%s", password),
		},
	}

	return DockerComposeConfig{
		Version:  "3",
		Services: services,
	}
}
