// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package runtime

import (
	"fmt"
	"strings"

	"github.com/clivern/peanut/core/definition"
	"github.com/clivern/peanut/core/util"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// DockerCompose type
type DockerCompose struct {
}

// NewDockerCompose creates a new instance
func NewDockerCompose() *DockerCompose {
	instance := new(DockerCompose)

	return instance
}

// Deploy deploys services
func (d *DockerCompose) Deploy(serviceID, service string, configs map[string]string) (map[string]string, error) {
	var def definition.DockerComposeConfig
	var err error

	dynamicConfigs := make(map[string]string)

	// Deploy Redis
	if definition.RedisService == service {
		dynamicConfigs["password"] = util.GetVal(configs, "password", definition.RedisDefaultPassword)

		def = definition.GetRedisConfig(serviceID, dynamicConfigs["password"])

		err = d.deployService(serviceID, def)

		if err != nil {
			return dynamicConfigs, err
		}

		dynamicConfigs["port"], err = d.fetchServicePort(serviceID, definition.RedisPort, def)

		if err != nil {
			return dynamicConfigs, err
		}

		// Deploy Etcd
	} else if definition.EtcdService == service {
		def = definition.GetEtcdConfig(serviceID)

		err = d.deployService(serviceID, def)

		if err != nil {
			return dynamicConfigs, err
		}

		dynamicConfigs["port"], err = d.fetchServicePort(serviceID, definition.EtcdPort, def)

		if err != nil {
			return dynamicConfigs, err
		}

		// Deploy Grafana
	} else if definition.GrafanaService == service {
		dynamicConfigs["username"] = util.GetVal(configs, "username", definition.GrafanaDefaultUsername)
		dynamicConfigs["password"] = util.GetVal(configs, "password", definition.GrafanaDefaultPassword)

		def = definition.GetGrafanaConfig(serviceID, dynamicConfigs["username"], dynamicConfigs["password"])

		err = d.deployService(serviceID, def)

		if err != nil {
			return dynamicConfigs, err
		}

		dynamicConfigs["port"], err = d.fetchServicePort(serviceID, definition.GrafanaPort, def)

		if err != nil {
			return dynamicConfigs, err
		}

		// Deploy MariaDB
	} else if definition.MariaDBService == service {
		dynamicConfigs["rootPassword"] = util.GetVal(configs, "rootPassword", definition.MariaDBDefaultRootPassword)
		dynamicConfigs["database"] = util.GetVal(configs, "database", definition.MariaDBDefaultDatabase)
		dynamicConfigs["username"] = util.GetVal(configs, "username", definition.MariaDBDefaultUsername)
		dynamicConfigs["password"] = util.GetVal(configs, "password", definition.MariaDBDefaultPassword)

		def = definition.GetMariaDBConfig(
			serviceID,
			dynamicConfigs["rootPassword"],
			dynamicConfigs["database"],
			dynamicConfigs["username"],
			dynamicConfigs["password"],
		)

		err = d.deployService(serviceID, def)

		if err != nil {
			return dynamicConfigs, err
		}

		dynamicConfigs["port"], err = d.fetchServicePort(serviceID, definition.MariaDBPort, def)

		if err != nil {
			return dynamicConfigs, err
		}

		// Deploy MySQL
	} else if definition.MySQLService == service {
		dynamicConfigs["rootPassword"] = util.GetVal(configs, "rootPassword", definition.MySQLDefaultRootPassword)
		dynamicConfigs["database"] = util.GetVal(configs, "database", definition.MySQLDefaultDatabase)
		dynamicConfigs["username"] = util.GetVal(configs, "username", definition.MySQLDefaultUsername)
		dynamicConfigs["password"] = util.GetVal(configs, "password", definition.MySQLDefaultPassword)

		def = definition.GetMySQLConfig(
			serviceID,
			dynamicConfigs["rootPassword"],
			dynamicConfigs["database"],
			dynamicConfigs["username"],
			dynamicConfigs["password"],
		)

		err = d.deployService(serviceID, def)

		if err != nil {
			return dynamicConfigs, err
		}

		dynamicConfigs["port"], err = d.fetchServicePort(serviceID, definition.MySQLPort, def)

		if err != nil {
			return dynamicConfigs, err
		}
	}

	return dynamicConfigs, nil
}

// Destroy destroys services
func (d *DockerCompose) Destroy(serviceID, service string, configs map[string]string) error {
	var def definition.DockerComposeConfig

	if definition.RedisService == service {

		// Get Redis Definition
		def = definition.GetRedisConfig(
			serviceID,
			util.GetVal(configs, "password", definition.RedisDefaultPassword),
		)

	} else if definition.EtcdService == service {

		// Get Etcd Definition
		def = definition.GetEtcdConfig(serviceID)

	} else if definition.GrafanaService == service {

		// Get Grafana Definition
		def = definition.GetGrafanaConfig(
			serviceID,
			util.GetVal(configs, "username", definition.GrafanaDefaultUsername),
			util.GetVal(configs, "password", definition.GrafanaDefaultPassword),
		)

	} else if definition.MariaDBService == service {

		// Get MariaDB Definition
		def = definition.GetMariaDBConfig(
			serviceID,
			util.GetVal(configs, "rootPassword", definition.MariaDBDefaultRootPassword),
			util.GetVal(configs, "database", definition.MariaDBDefaultDatabase),
			util.GetVal(configs, "username", definition.MariaDBDefaultUsername),
			util.GetVal(configs, "password", definition.MariaDBDefaultPassword),
		)

	} else if definition.MySQLService == service {

		// Get MySQL Definition
		def = definition.GetMySQLConfig(
			serviceID,
			util.GetVal(configs, "rootPassword", definition.MySQLDefaultRootPassword),
			util.GetVal(configs, "database", definition.MySQLDefaultDatabase),
			util.GetVal(configs, "username", definition.MySQLDefaultUsername),
			util.GetVal(configs, "password", definition.MySQLDefaultPassword),
		)

	}

	err := d.destroyService(serviceID, def)

	if err != nil {
		return err
	}

	return d.Prune()
}

// Prune remove all unused containers, networks, images
func (d *DockerCompose) Prune() error {
	command := "docker system prune -a -f --volumes"

	stdout, stderr, err := util.Exec(command)

	log.WithFields(log.Fields{
		"command": command,
	}).Info("Run a shell command")

	if err != nil {
		return err
	}

	// Store runtime verbose logs only in dev environment
	if viper.GetString("app.mode") == "dev" {
		err = util.StoreFile(
			fmt.Sprintf("%s/prune.stdout.log", viper.GetString("app.storage.path")),
			stdout,
		)

		if err != nil {
			return err
		}

		err = util.StoreFile(
			fmt.Sprintf("%s/prune.stderr.log", viper.GetString("app.storage.path")),
			stderr,
		)

		if err != nil {
			return err
		}
	}

	return nil
}

// deployService deploys a service
func (d *DockerCompose) deployService(serviceID string, definition definition.DockerComposeConfig) error {
	result, err := definition.ToString()

	if err != nil {
		return err
	}

	err = util.StoreFile(
		fmt.Sprintf("%s/%s.yml", viper.GetString("app.storage.path"), serviceID),
		result,
	)

	if err != nil {
		return err
	}

	command := fmt.Sprintf(
		"docker-compose -f %s/%s.yml -p %s up -d --force-recreate",
		viper.GetString("app.storage.path"),
		serviceID,
		serviceID,
	)

	stdout, stderr, err := util.Exec(command)

	log.WithFields(log.Fields{
		"command": command,
	}).Info("Run a shell command")

	if err != nil {
		return err
	}

	// Store runtime verbose logs only in dev environment
	if viper.GetString("app.mode") == "dev" {
		err = util.StoreFile(
			fmt.Sprintf("%s/%s.deploy.stdout.log", viper.GetString("app.storage.path"), serviceID),
			stdout,
		)

		if err != nil {
			return err
		}

		err = util.StoreFile(
			fmt.Sprintf("%s/%s.deploy.stderr.log", viper.GetString("app.storage.path"), serviceID),
			stderr,
		)

		if err != nil {
			return err
		}
	}

	return nil
}

// destroyService destroys a service
func (d *DockerCompose) destroyService(serviceID string, definition definition.DockerComposeConfig) error {
	result, err := definition.ToString()

	if err != nil {
		return err
	}

	err = util.StoreFile(
		fmt.Sprintf("%s/%s.yml", viper.GetString("app.storage.path"), serviceID),
		result,
	)

	if err != nil {
		return err
	}

	command := fmt.Sprintf(
		"docker-compose -f %s/%s.yml -p %s down -v --remove-orphans",
		viper.GetString("app.storage.path"),
		serviceID,
		serviceID,
	)

	stdout, stderr, err := util.Exec(command)

	log.WithFields(log.Fields{
		"command": command,
	}).Info("Run a shell command")

	if err != nil {
		return err
	}

	// Store runtime verbose logs only in dev environment
	if viper.GetString("app.mode") == "dev" {
		err = util.StoreFile(
			fmt.Sprintf("%s/%s.destroy.stdout.log", viper.GetString("app.storage.path"), serviceID),
			stdout,
		)

		if err != nil {
			return err
		}

		err = util.StoreFile(
			fmt.Sprintf("%s/%s.destroy.stderr.log", viper.GetString("app.storage.path"), serviceID),
			stderr,
		)

		if err != nil {
			return err
		}
	}

	return nil
}

// fetchServicePort get assigned port
func (d *DockerCompose) fetchServicePort(serviceID string, port string, definition definition.DockerComposeConfig) (string, error) {
	result, err := definition.ToString()

	if err != nil {
		return "", err
	}

	err = util.StoreFile(
		fmt.Sprintf("%s/%s.yml", viper.GetString("app.storage.path"), serviceID),
		result,
	)

	if err != nil {
		return "", err
	}

	command := fmt.Sprintf(
		"docker-compose -f %s/%s.yml -p %s port %s %s",
		viper.GetString("app.storage.path"),
		serviceID,
		serviceID,
		serviceID,
		port,
	)

	stdout, stderr, err := util.Exec(command)

	log.WithFields(log.Fields{
		"command": command,
	}).Info("Run a shell command")

	if err != nil {
		return "", err
	}

	// Store runtime verbose logs only in dev environment
	if viper.GetString("app.mode") == "dev" {
		err = util.StoreFile(
			fmt.Sprintf("%s/%s.port_%s.stdout.log", viper.GetString("app.storage.path"), serviceID, port),
			stdout,
		)

		if err != nil {
			return "", err
		}

		err = util.StoreFile(
			fmt.Sprintf("%s/%s.port_%s.stderr.log", viper.GetString("app.storage.path"), serviceID, port),
			stderr,
		)

		if err != nil {
			return "", err
		}
	}

	return strings.TrimSuffix(strings.Replace(stdout, "0.0.0.0:", "", -1), "\n"), nil
}
