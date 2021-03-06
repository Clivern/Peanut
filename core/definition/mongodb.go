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
	MongoDBDockerImage = "mongo"

	// MongoDBDockerImageVersion const
	MongoDBDockerImageVersion = "5.0.0-rc7"

	// MongoDBRestartPolicy const
	MongoDBRestartPolicy = "unless-stopped"

	// MongoDBDefaultDatabase const
	MongoDBDefaultDatabase = "peanut"

	// MongoDBDefaultUsername const
	MongoDBDefaultUsername = "peanut"

	// MongoDBDefaultPassword const
	MongoDBDefaultPassword = "secret"
)

// GetMongoDBConfig gets yaml definition object
func GetMongoDBConfig(name, version, database, username, password string) DockerComposeConfig {
	services := make(map[string]Service)

	if database == "" {
		database = MongoDBDefaultDatabase
	}

	if username == "" {
		username = MongoDBDefaultUsername
	}

	if password == "" {
		password = MongoDBDefaultPassword
	}

	if version == "" {
		version = MongoDBDockerImageVersion
	}

	services[name] = Service{
		Image:   fmt.Sprintf("%s:%s", MongoDBDockerImage, version),
		Restart: MongoDBRestartPolicy,
		Ports:   []string{MongoDBPort},
		Environment: []string{
			fmt.Sprintf("MONGO_INITDB_DATABASE=%s", database),
			fmt.Sprintf("MONGO_INITDB_ROOT_USERNAME=%s", username),
			fmt.Sprintf("MONGO_INITDB_ROOT_PASSWORD=%s", password),
		},
	}

	return DockerComposeConfig{
		Version:  "3",
		Services: services,
	}
}
