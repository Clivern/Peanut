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
func (d *DockerCompose) Deploy(serviceID, service, version string, configs map[string]string) (map[string]string, error) {
	var def definition.DockerComposeConfig
	var err error

	dynamicConfigs := make(map[string]string)

	if definition.RedisService == service {
		// Deploy Redis
		dynamicConfigs["password"] = util.GetVal(configs, "password", definition.RedisDefaultPassword)

		def = definition.GetRedisConfig(serviceID, version, dynamicConfigs["password"])

		err = d.deployService(serviceID, def)

		if err != nil {
			return dynamicConfigs, err
		}

		dynamicConfigs["port"], err = d.fetchServicePort(serviceID, definition.RedisPort, def)

		if err != nil {
			return dynamicConfigs, err
		}

	} else if definition.EtcdService == service {
		// Deploy Etcd
		def = definition.GetEtcdConfig(serviceID, version)

		err = d.deployService(serviceID, def)

		if err != nil {
			return dynamicConfigs, err
		}

		dynamicConfigs["port"], err = d.fetchServicePort(serviceID, definition.EtcdPort, def)

		if err != nil {
			return dynamicConfigs, err
		}

	} else if definition.GrafanaService == service {
		// Deploy Grafana
		dynamicConfigs["username"] = util.GetVal(configs, "username", definition.GrafanaDefaultUsername)
		dynamicConfigs["password"] = util.GetVal(configs, "password", definition.GrafanaDefaultPassword)
		dynamicConfigs["anonymousAccess"] = util.GetVal(configs, "anonymousAccess", definition.GrafanaDefaultAnonymousAccess)
		dynamicConfigs["allowSignup"] = util.GetVal(configs, "allowSignup", definition.GrafanaDefaultAllowSignup)

		def = definition.GetGrafanaConfig(
			serviceID,
			version,
			dynamicConfigs["username"],
			dynamicConfigs["password"],
			dynamicConfigs["allowSignup"],
			dynamicConfigs["anonymousAccess"],
		)

		err = d.deployService(serviceID, def)

		if err != nil {
			return dynamicConfigs, err
		}

		dynamicConfigs["port"], err = d.fetchServicePort(serviceID, definition.GrafanaPort, def)

		if err != nil {
			return dynamicConfigs, err
		}

	} else if definition.MariaDBService == service {
		// Deploy MariaDB
		dynamicConfigs["rootPassword"] = util.GetVal(configs, "rootPassword", definition.MariaDBDefaultRootPassword)
		dynamicConfigs["database"] = util.GetVal(configs, "database", definition.MariaDBDefaultDatabase)
		dynamicConfigs["username"] = util.GetVal(configs, "username", definition.MariaDBDefaultUsername)
		dynamicConfigs["password"] = util.GetVal(configs, "password", definition.MariaDBDefaultPassword)

		def = definition.GetMariaDBConfig(
			serviceID,
			version,
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

	} else if definition.MySQLService == service {
		// Deploy MySQL
		dynamicConfigs["rootPassword"] = util.GetVal(configs, "rootPassword", definition.MySQLDefaultRootPassword)
		dynamicConfigs["database"] = util.GetVal(configs, "database", definition.MySQLDefaultDatabase)
		dynamicConfigs["username"] = util.GetVal(configs, "username", definition.MySQLDefaultUsername)
		dynamicConfigs["password"] = util.GetVal(configs, "password", definition.MySQLDefaultPassword)

		def = definition.GetMySQLConfig(
			serviceID,
			version,
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
	} else if definition.PostgreSQLService == service {
		// Deploy Postgresql
		dynamicConfigs["database"] = util.GetVal(configs, "database", definition.PostgreSQLDefaultDatabase)
		dynamicConfigs["username"] = util.GetVal(configs, "username", definition.PostgreSQLDefaultUsername)
		dynamicConfigs["password"] = util.GetVal(configs, "password", definition.PostgreSQLDefaultPassword)

		def = definition.GetPostgreSQLConfig(
			serviceID,
			version,
			dynamicConfigs["database"],
			dynamicConfigs["username"],
			dynamicConfigs["password"],
		)

		err = d.deployService(serviceID, def)

		if err != nil {
			return dynamicConfigs, err
		}

		dynamicConfigs["port"], err = d.fetchServicePort(serviceID, definition.PostgreSQLPort, def)

		if err != nil {
			return dynamicConfigs, err
		}
	} else if definition.MongoDBService == service {
		// Deploy MongoDB
		dynamicConfigs["database"] = util.GetVal(configs, "database", definition.MongoDBDefaultDatabase)
		dynamicConfigs["username"] = util.GetVal(configs, "username", definition.MongoDBDefaultUsername)
		dynamicConfigs["password"] = util.GetVal(configs, "password", definition.MongoDBDefaultPassword)

		def = definition.GetMongoDBConfig(
			serviceID,
			version,
			dynamicConfigs["database"],
			dynamicConfigs["username"],
			dynamicConfigs["password"],
		)

		err = d.deployService(serviceID, def)

		if err != nil {
			return dynamicConfigs, err
		}

		dynamicConfigs["port"], err = d.fetchServicePort(serviceID, definition.MongoDBPort, def)

		if err != nil {
			return dynamicConfigs, err
		}
	} else if definition.ElasticSearchService == service {
		// Deploy ElasticSearch
		def = definition.GetElasticSearchConfig(serviceID, version)

		err = d.deployService(serviceID, def)

		if err != nil {
			return dynamicConfigs, err
		}

		dynamicConfigs["requestsPort"], err = d.fetchServicePort(
			serviceID,
			definition.ElasticSearchRequestsPort,
			def,
		)

		if err != nil {
			return dynamicConfigs, err
		}

		dynamicConfigs["communicationPort"], err = d.fetchServicePort(
			serviceID,
			definition.ElasticSearchCommunicationPort,
			def,
		)

		if err != nil {
			return dynamicConfigs, err
		}
	} else if definition.GraphiteService == service {
		// Deploy Graphite
		def = definition.GetGraphiteConfig(serviceID, version)

		err = d.deployService(serviceID, def)

		if err != nil {
			return dynamicConfigs, err
		}

		dynamicConfigs["webPort"], err = d.fetchServicePort(
			serviceID,
			definition.GraphiteWebPort,
			def,
		)

		if err != nil {
			return dynamicConfigs, err
		}

		dynamicConfigs["carbonPort"], err = d.fetchServicePort(
			serviceID,
			definition.GraphiteCarbonPort,
			def,
		)

		if err != nil {
			return dynamicConfigs, err
		}

		dynamicConfigs["carbonPicklePort"], err = d.fetchServicePort(
			serviceID,
			definition.GraphiteCarbonPicklePort,
			def,
		)

		if err != nil {
			return dynamicConfigs, err
		}

		dynamicConfigs["carbonAggregatorPort"], err = d.fetchServicePort(
			serviceID,
			definition.GraphiteCarbonAggregatorPort,
			def,
		)

		if err != nil {
			return dynamicConfigs, err
		}

		dynamicConfigs["carbonAggregatorPicklePort"], err = d.fetchServicePort(
			serviceID,
			definition.GraphiteCarbonAggregatorPicklePort,
			def,
		)

		if err != nil {
			return dynamicConfigs, err
		}

		dynamicConfigs["statsdPort"], err = d.fetchServicePort(
			serviceID,
			definition.GraphiteStatsdPort,
			def,
		)

		if err != nil {
			return dynamicConfigs, err
		}

		dynamicConfigs["statsdAdminPort"], err = d.fetchServicePort(
			serviceID,
			definition.GraphiteStatsdAdminPort,
			def,
		)

		if err != nil {
			return dynamicConfigs, err
		}
	} else if definition.PrometheusService == service {
		// Get Prometheus Definition
		dynamicConfigs["configsBase64Encoded"] = util.GetVal(configs, "configsBase64Encoded", definition.PrometheusDefaultConfig)

		// Base64 decode configs
		plainConfigs, err := util.Base64Decode(dynamicConfigs["configsBase64Encoded"])

		if err != nil {
			return dynamicConfigs, err
		}

		// Store Prometheus config file in peanut local storage & mount later
		err = util.StoreFile(
			fmt.Sprintf("%s/%s.prometheus.yml", util.RemoveTrailingSlash(viper.GetString("app.storage.path")), serviceID),
			plainConfigs,
		)

		if err != nil {
			return dynamicConfigs, err
		}

		def = definition.GetPrometheusConfig(
			serviceID,
			version,
			fmt.Sprintf("%s/%s.prometheus.yml", util.RemoveTrailingSlash(viper.GetString("app.storage.path")), serviceID),
		)

		err = d.deployService(serviceID, def)

		if err != nil {
			return dynamicConfigs, err
		}

		dynamicConfigs["port"], err = d.fetchServicePort(
			serviceID,
			definition.PrometheusPort,
			def,
		)
	} else if definition.ZipkinService == service {
		// Deploy Zipkin
		def = definition.GetZipkinConfig(serviceID, version)

		err = d.deployService(serviceID, def)

		if err != nil {
			return dynamicConfigs, err
		}

		dynamicConfigs["port"], err = d.fetchServicePort(serviceID, definition.ZipkinPort, def)

		if err != nil {
			return dynamicConfigs, err
		}

	} else if definition.MemcachedService == service {
		// Deploy Memcached
		def = definition.GetMemcachedConfig(serviceID, version)

		err = d.deployService(serviceID, def)

		if err != nil {
			return dynamicConfigs, err
		}

		dynamicConfigs["port"], err = d.fetchServicePort(serviceID, definition.MemcachedPort, def)

		if err != nil {
			return dynamicConfigs, err
		}
	} else if definition.MailhogService == service {
		// Deploy Mailhog
		def = definition.GetMailhogConfig(serviceID, version)

		err = d.deployService(serviceID, def)

		if err != nil {
			return dynamicConfigs, err
		}

		dynamicConfigs["smtpPort"], err = d.fetchServicePort(serviceID, definition.MailhogSMTPPort, def)

		if err != nil {
			return dynamicConfigs, err
		}

		dynamicConfigs["httpPort"], err = d.fetchServicePort(serviceID, definition.MailhogHTTPPort, def)

		if err != nil {
			return dynamicConfigs, err
		}

	} else if definition.JaegerService == service {
		// Deploy Jaeger
		def = definition.GetJaegerConfig(serviceID, version)

		err = d.deployService(serviceID, def)

		if err != nil {
			return dynamicConfigs, err
		}

		dynamicConfigs["udpPort1"], err = d.fetchServicePort(serviceID, definition.JaegerUDPPort1, def)

		if err != nil {
			return dynamicConfigs, err
		}

		dynamicConfigs["udpPort2"], err = d.fetchServicePort(serviceID, definition.JaegerUDPPort2, def)

		if err != nil {
			return dynamicConfigs, err
		}

		dynamicConfigs["udpPort3"], err = d.fetchServicePort(serviceID, definition.JaegerUDPPort3, def)

		if err != nil {
			return dynamicConfigs, err
		}

		dynamicConfigs["httpPort1"], err = d.fetchServicePort(serviceID, definition.JaegerHTTPPort1, def)

		if err != nil {
			return dynamicConfigs, err
		}

		dynamicConfigs["httpPort2"], err = d.fetchServicePort(serviceID, definition.JaegerHTTPPort2, def)

		if err != nil {
			return dynamicConfigs, err
		}

		dynamicConfigs["httpPort3"], err = d.fetchServicePort(serviceID, definition.JaegerHTTPPort3, def)

		if err != nil {
			return dynamicConfigs, err
		}

		dynamicConfigs["httpPort4"], err = d.fetchServicePort(serviceID, definition.JaegerHTTPPort4, def)

		if err != nil {
			return dynamicConfigs, err
		}

		dynamicConfigs["httpPort5"], err = d.fetchServicePort(serviceID, definition.JaegerHTTPPort5, def)

		if err != nil {
			return dynamicConfigs, err
		}
	} else if definition.RabbitMQService == service {
		// Deploy RabbitMQ
		def = definition.GetRabbitMQConfig(serviceID, version)

		err = d.deployService(serviceID, def)

		if err != nil {
			return dynamicConfigs, err
		}

		dynamicConfigs["username"] = definition.RabbitMQDefaultUsername

		dynamicConfigs["password"] = definition.RabbitMQDefaultPassword

		dynamicConfigs["amqpPort"], err = d.fetchServicePort(serviceID, definition.RabbitMQAMQPPort, def)

		if err != nil {
			return dynamicConfigs, err
		}

		dynamicConfigs["dashboardPort"], err = d.fetchServicePort(serviceID, definition.RabbitMQDashboardPort, def)

		if err != nil {
			return dynamicConfigs, err
		}

	} else if definition.ConsulService == service {
		// Deploy Consul
		def = definition.GetConsulConfig(
			serviceID,
			version,
		)

		err = d.deployService(serviceID, def)

		if err != nil {
			return dynamicConfigs, err
		}

		dynamicConfigs["httpPort"], err = d.fetchServicePort(serviceID, definition.ConsulHTTPPort, def)

		if err != nil {
			return dynamicConfigs, err
		}

	} else if definition.VaultService == service {
		// Deploy Vault
		dynamicConfigs["token"] = util.GetVal(configs, "token", definition.VaultDefaultToken)

		def = definition.GetVaultConfig(
			serviceID,
			version,
			dynamicConfigs["token"],
		)

		err = d.deployService(serviceID, def)

		if err != nil {
			return dynamicConfigs, err
		}

		dynamicConfigs["httpPort"], err = d.fetchServicePort(serviceID, definition.VaultHTTPPort, def)

		if err != nil {
			return dynamicConfigs, err
		}
	} else if definition.CassandraService == service {
		// Deploy Cassandra
		def = definition.GetCassandraConfig(
			serviceID,
			version,
		)

		err = d.deployService(serviceID, def)

		if err != nil {
			return dynamicConfigs, err
		}

		dynamicConfigs["port"], err = d.fetchServicePort(serviceID, definition.CassandraPort, def)

		if err != nil {
			return dynamicConfigs, err
		}
	} else if definition.MinioService == service {
		// Deploy Minio
		dynamicConfigs["username"] = util.GetVal(configs, "username", definition.MinioRootUser)
		dynamicConfigs["password"] = util.GetVal(configs, "password", definition.MinioRootPassword)

		def = definition.GetMinioConfig(
			serviceID,
			version,
			dynamicConfigs["username"],
			dynamicConfigs["password"],
		)

		err = d.deployService(serviceID, def)

		if err != nil {
			return dynamicConfigs, err
		}

		dynamicConfigs["apiPort"], err = d.fetchServicePort(serviceID, definition.MinioAPIPort, def)

		if err != nil {
			return dynamicConfigs, err
		}

		dynamicConfigs["consolePort"], err = d.fetchServicePort(serviceID, definition.MinioConsolePort, def)

		if err != nil {
			return dynamicConfigs, err
		}
	} else if definition.RegistryService == service {
		// Deploy Registry
		def = definition.GetRegistryConfig(serviceID, version)

		err = d.deployService(serviceID, def)

		if err != nil {
			return dynamicConfigs, err
		}

		dynamicConfigs["port"], err = d.fetchServicePort(serviceID, definition.RegistryPort, def)

		if err != nil {
			return dynamicConfigs, err
		}
	} else if definition.GhostService == service {
		// Deploy Ghost
		def = definition.GetGhostConfig(serviceID, version)

		err = d.deployService(serviceID, def)

		if err != nil {
			return dynamicConfigs, err
		}

		dynamicConfigs["port"], err = d.fetchServicePort(serviceID, definition.GhostPort, def)

		if err != nil {
			return dynamicConfigs, err
		}
	} else if definition.HttpbinService == service {
		// Deploy Httpbin
		def = definition.GetHttpbinConfig(serviceID, version)

		err = d.deployService(serviceID, def)

		if err != nil {
			return dynamicConfigs, err
		}

		dynamicConfigs["port"], err = d.fetchServicePort(serviceID, definition.HttpbinPort, def)

		if err != nil {
			return dynamicConfigs, err
		}
	} else if definition.EtherpadService == service {
		// Deploy Etherpad
		def = definition.GetEtherpadConfig(serviceID, version)

		err = d.deployService(serviceID, def)

		if err != nil {
			return dynamicConfigs, err
		}

		dynamicConfigs["port"], err = d.fetchServicePort(serviceID, definition.EtherpadPort, def)

		if err != nil {
			return dynamicConfigs, err
		}
	} else if definition.NagiosService == service {
		dynamicConfigs["username"] = definition.NagiosRootUser
		dynamicConfigs["password"] = definition.NagiosRootPassword

		// Deploy Nagios
		def = definition.GetNagiosConfig(serviceID, version)

		err = d.deployService(serviceID, def)

		if err != nil {
			return dynamicConfigs, err
		}

		dynamicConfigs["port"], err = d.fetchServicePort(serviceID, definition.NagiosPort, def)

		if err != nil {
			return dynamicConfigs, err
		}
	}

	return dynamicConfigs, nil
}

// Destroy destroys services
func (d *DockerCompose) Destroy(serviceID, service, version string, configs map[string]string) error {
	var def definition.DockerComposeConfig

	if definition.RedisService == service {
		// Get Redis Definition
		def = definition.GetRedisConfig(
			serviceID,
			version,
			util.GetVal(configs, "password", definition.RedisDefaultPassword),
		)

	} else if definition.EtcdService == service {
		// Get Etcd Definition
		def = definition.GetEtcdConfig(serviceID, version)

	} else if definition.GrafanaService == service {
		// Get Grafana Definition
		def = definition.GetGrafanaConfig(
			serviceID,
			version,
			util.GetVal(configs, "username", definition.GrafanaDefaultUsername),
			util.GetVal(configs, "password", definition.GrafanaDefaultPassword),
			util.GetVal(configs, "allowSignup", definition.GrafanaDefaultAllowSignup),
			util.GetVal(configs, "anonymousAccess", definition.GrafanaDefaultAnonymousAccess),
		)

	} else if definition.MariaDBService == service {
		// Get MariaDB Definition
		def = definition.GetMariaDBConfig(
			serviceID,
			version,
			util.GetVal(configs, "rootPassword", definition.MariaDBDefaultRootPassword),
			util.GetVal(configs, "database", definition.MariaDBDefaultDatabase),
			util.GetVal(configs, "username", definition.MariaDBDefaultUsername),
			util.GetVal(configs, "password", definition.MariaDBDefaultPassword),
		)

	} else if definition.MySQLService == service {
		// Get MySQL Definition
		def = definition.GetMySQLConfig(
			serviceID,
			version,
			util.GetVal(configs, "rootPassword", definition.MySQLDefaultRootPassword),
			util.GetVal(configs, "database", definition.MySQLDefaultDatabase),
			util.GetVal(configs, "username", definition.MySQLDefaultUsername),
			util.GetVal(configs, "password", definition.MySQLDefaultPassword),
		)

	} else if definition.PostgreSQLService == service {
		// Get PostgreSQL Definition
		def = definition.GetPostgreSQLConfig(
			serviceID,
			version,
			util.GetVal(configs, "database", definition.PostgreSQLDefaultDatabase),
			util.GetVal(configs, "username", definition.PostgreSQLDefaultUsername),
			util.GetVal(configs, "password", definition.PostgreSQLDefaultPassword),
		)

	} else if definition.MongoDBService == service {
		// Get MongoDB Definition
		def = definition.GetMongoDBConfig(
			serviceID,
			version,
			util.GetVal(configs, "database", definition.MongoDBDefaultDatabase),
			util.GetVal(configs, "username", definition.MongoDBDefaultUsername),
			util.GetVal(configs, "password", definition.MongoDBDefaultPassword),
		)

	} else if definition.ElasticSearchService == service {
		// Get ElasticSearch Definition
		def = definition.GetElasticSearchConfig(serviceID, version)

	} else if definition.GraphiteService == service {
		// Get Graphite Definition
		def = definition.GetGraphiteConfig(serviceID, version)

	} else if definition.PrometheusService == service {
		// Get Prometheus Definition
		def = definition.GetPrometheusConfig(
			serviceID,
			version,
			fmt.Sprintf("%s/%s.prometheus.yml", util.RemoveTrailingSlash(viper.GetString("app.storage.path")), serviceID),
		)

	} else if definition.ZipkinService == service {
		// Get Zipkin Definition
		def = definition.GetZipkinConfig(serviceID, version)

	} else if definition.MemcachedService == service {
		// Get Memcached Definition
		def = definition.GetMemcachedConfig(serviceID, version)

	} else if definition.MailhogService == service {
		// Get Mailhog Definition
		def = definition.GetMailhogConfig(serviceID, version)

	} else if definition.JaegerService == service {
		// Get Jaeger Definition
		def = definition.GetJaegerConfig(serviceID, version)

	} else if definition.RabbitMQService == service {
		// Get RabbitMQ Definition
		def = definition.GetRabbitMQConfig(serviceID, version)

	} else if definition.ConsulService == service {
		// Get Consul Definition
		def = definition.GetConsulConfig(serviceID, version)

	} else if definition.VaultService == service {
		// Get Vault Definition
		def = definition.GetVaultConfig(
			serviceID,
			version,
			util.GetVal(configs, "token", definition.VaultDefaultToken),
		)

	} else if definition.CassandraService == service {
		// Get Cassandra Definition
		def = definition.GetCassandraConfig(
			serviceID,
			version,
		)

	} else if definition.MinioService == service {
		// Get Minio Definition
		def = definition.GetMinioConfig(
			serviceID,
			version,
			util.GetVal(configs, "username", definition.MinioRootUser),
			util.GetVal(configs, "password", definition.MinioRootPassword),
		)

	} else if definition.MailhogService == service {
		// Get Registry Definition
		def = definition.GetRegistryConfig(serviceID, version)

	} else if definition.GhostService == service {
		// Get Ghost Definition
		def = definition.GetGhostConfig(serviceID, version)

	} else if definition.HttpbinService == service {
		// Get Httpbin Definition
		def = definition.GetHttpbinConfig(serviceID, version)

	} else if definition.EtherpadService == service {
		// Get Etherpad Definition
		def = definition.GetEtherpadConfig(serviceID, version)

	} else if definition.NagiosService == service {
		// Get Nagios Definition
		def = definition.GetNagiosConfig(serviceID, version)
	}

	err := d.destroyService(serviceID, def)

	if err != nil {
		return err
	}

	if viper.GetBool("app.containerization.autoClean") {
		return d.Prune()
	}

	return nil
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
