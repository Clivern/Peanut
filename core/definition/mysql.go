// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package definition

import (
	"fmt"
)

const (
	// MySQLService const
	MySQLService = "mysql"

	// MySQLPort const
	MySQLPort = "3306"

	// MySQLDockerImage const
	MySQLDockerImage = "mysql:8.0"

	// MySQLRestartPolicy const
	MySQLRestartPolicy = "unless-stopped"

	// MySQLDefaultRootPassword const
	MySQLDefaultRootPassword = "root"

	// MySQLDefaultDatabase const
	MySQLDefaultDatabase = "peanut"

	// MySQLDefaultUsername const
	MySQLDefaultUsername = "peanut"

	// MySQLDefaultPassword const
	MySQLDefaultPassword = "secret"
)

// GetMySQLConfig gets yaml definition object
func GetMySQLConfig(name, rootPassword, database, username, password string) DockerComposeConfig {
	services := make(map[string]Service)

	if rootPassword == "" {
		rootPassword = MySQLDefaultRootPassword
	}

	if database == "" {
		database = MySQLDefaultDatabase
	}

	if username == "" {
		username = MySQLDefaultUsername
	}

	if password == "" {
		password = MySQLDefaultPassword
	}

	services[name] = Service{
		Image:   MySQLDockerImage,
		Restart: MySQLRestartPolicy,
		Ports:   []string{MySQLPort},
		Environment: []string{
			fmt.Sprintf("MYSQL_ROOT_PASSWORD=%s", rootPassword),
			fmt.Sprintf("MYSQL_DATABASE=%s", database),
			fmt.Sprintf("MYSQL_USER=%s", username),
			fmt.Sprintf("MYSQL_PASSWORD=%s", password),
			"MYSQL_ALLOW_EMPTY_PASSWORD=no",
		},
	}

	return DockerComposeConfig{
		Version:  "3",
		Services: services,
	}
}
