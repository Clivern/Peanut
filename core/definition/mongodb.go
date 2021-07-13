// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package definition

import (
	"fmt"
)

const (
	// MongoDBService const
	MongoDBService = "mongodb"

	// MongoDBPort const
	MongoDBPort = "27017"

	// MongoDBDockerImage const
	MongoDBDockerImage = "mongo:5.0.0-rc7"

	// MongoDBRestartPolicy const
	MongoDBRestartPolicy = "unless-stopped"

	// MongoDBRootUsername const
	MongoDBRootUsername = "peanut"

	// MongoDBRootPassword const
	MongoDBRootPassword = "secret"
)

// GetMongoDBConfig gets yaml definition object
func GetMongoDBConfig(name, username, password string) DockerComposeConfig {
	services := make(map[string]Service)

	if username == "" {
		username = MongoDBRootUsername
	}

	if password == "" {
		password = MongoDBRootPassword
	}

	services[name] = Service{
		Image:   MongoDBDockerImage,
		Restart: MongoDBRestartPolicy,
		Ports:   []string{MongoDBPort},
		Environment: []string{
			fmt.Sprintf("MONGO_INITDB_ROOT_USERNAME=%s", username),
			fmt.Sprintf("MONGO_INITDB_ROOT_PASSWORD=%s", password),
		},
	}

	return DockerComposeConfig{
		Version:  "3",
		Services: services,
	}
}
