// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package definition

import (
	"fmt"
)

const (
	// MariaDBService const
	MariaDBService = "mariadb"

	// MariaDBPort const
	MariaDBPort = "3306"

	// MariaDBDockerImage const
	MariaDBDockerImage = "mariadb:10.6.2"

	// MariaDBRestartPolicy const
	MariaDBRestartPolicy = "unless-stopped"

	// MariaDBDefaultRootPassword const
	MariaDBDefaultRootPassword = "root"

	// MariaDBDefaultDatabase const
	MariaDBDefaultDatabase = "peanut"

	// MariaDBDefaultUsername const
	MariaDBDefaultUsername = "peanut"

	// MariaDBDefaultPassword const
	MariaDBDefaultPassword = "secret"
)

// GetMariaDBConfig gets yaml definition object
func GetMariaDBConfig(name, rootPassword, database, username, password string) DockerComposeConfig {
	services := make(map[string]Service)

	if rootPassword == "" {
		rootPassword = MariaDBDefaultRootPassword
	}

	if database == "" {
		database = MariaDBDefaultDatabase
	}

	if username == "" {
		username = MariaDBDefaultUsername
	}

	if password == "" {
		password = MariaDBDefaultPassword
	}

	services[name] = Service{
		Image:   MariaDBDockerImage,
		Restart: MariaDBRestartPolicy,
		Ports:   []string{MariaDBPort},
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
