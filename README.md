<p align="center">
    <img src="https://raw.githubusercontent.com/Clivern/Peanut/main/assets/logo.png?v=0.3.0" width="240" />
    <h3 align="center">Peanut</h3>
    <p align="center">Deploy Databases and Services Easily for Development and Testing Pipelines.</p>
    <p align="center">
        <a href="https://github.com/Clivern/Peanut/actions/workflows/build.yml">
            <img src="https://github.com/Clivern/Peanut/actions/workflows/build.yml/badge.svg">
        </a>
        <a href="https://github.com/Clivern/Peanut/actions">
            <img src="https://github.com/Clivern/Peanut/workflows/Release/badge.svg">
        </a>
        <a href="https://github.com/Clivern/Peanut/releases">
            <img src="https://img.shields.io/badge/Version-0.3.0-red.svg">
        </a>
        <a href="https://goreportcard.com/report/github.com/Clivern/Peanut">
            <img src="https://goreportcard.com/badge/github.com/Clivern/Peanut?v=0.3.0">
        </a>
        <a href="https://godoc.org/github.com/clivern/peanut">
            <img src="https://godoc.org/github.com/clivern/peanut?status.svg">
        </a>
        <a href="https://github.com/Clivern/Peanut/blob/master/LICENSE">
            <img src="https://img.shields.io/badge/LICENSE-MIT-orange.svg">
        </a>
    </p>
</p>
<br/>
<p align="center">
    <img src="https://raw.githubusercontent.com/Clivern/Peanut/main/assets/chart.png?v=0.3.0" width="80%" />
</p>
<p align="center">
    <h4 align="center">Dashboard Screenshots</h4>
    <p align="center">
        <img src="https://raw.githubusercontent.com/Clivern/Peanut/main/assets/screenshot_01.png?v=0.3.0" width="90%" />
        <img src="https://raw.githubusercontent.com/Clivern/Peanut/main/assets/screenshot_02.png?v=0.3.0" width="90%" />
        <img src="https://raw.githubusercontent.com/Clivern/Peanut/main/assets/screenshot_03.png?v=0.3.0" width="90%" />
        <img src="https://raw.githubusercontent.com/Clivern/Peanut/main/assets/screenshot_04.png?v=0.3.0" width="90%" />
    </p>
</p>


Peanut provides a REST API, Admin Dashboard and a command line tool to deploy and configure the commonly used services like databases, message brokers, graphing, tracing, caching tools ... etc. It perfectly suited for development, manual testing, automated testing pipelines where mocking is not possible and test drives.

Under the hood, it works with the containerization runtime like docker to deploy and configure the service. Destroy the service if it is a temporary one.

Technically you can achieve the same with a bunch of yaml files or using a configuration management tool or a package manager like helm but peanut is pretty small and fun to use & should speed up your workflow!

Supported Services:

- MySQL.
- MariaDB.
- PostgreSQL.
- Redis.
- Etcd.
- Grafana.
- Elasticsearch.
- MongoDB.
- Graphite.
- Prometheus.
- Zipkin.
- Memcached.
- Mailhog.
- Jaeger.
- RabbitMQ.
- Consul.
- Vault.


## Documentation

#### Run Peanut on Ubuntu

To run peanut on ubuntu, You can use the following bash script since it may take a while for a cold start. the script will install etcd, docker, docker-compose and peanut.

```zsh
$ bash < <(curl -s https://raw.githubusercontent.com/Clivern/Peanut/main/deployment/linux/install.sh)

# Get The Public IP
$ curl https://ipinfo.io/ip
x.x.x.x
```

Peanut will be running on `80` port and UI on this URL `http://x.x.x.x`. Please open this file `/etc/peanut/config.prod.yml` and adjust the following line to be your `Public IP` or `hostname`

```zsh
# App configs
app:
    ...
    # Hostname
    hostname: ${PEANUT_API_HOSTNAME:-127.0.0.1}
```

Then Restart Peanut

```zsh
$ systemctl restart peanut
```

To make sure peanut is running, Run the following from your laptop to run some redis instances for 10 minutes.

```zsh
# To provision redis services for 10 minutes
$ curl -X POST http://$PUBLIC_IP/api/v1/service -d '{"service":"redis","configs": {},"deleteAfter":"10min"}' -H 'x-api-key: ~api~key~here~'

{
  "createdAt": "2021-07-11T09:58:11.076Z",
  "id": "aadd5741-58c5-43c7-94fd-e6c0171fe8be",
  "service": "a8138f52-3ebb-4a34-b403-1be6ad481daf",
  "status": "PENDING",
  "type": "service.deploy"
}


# To list services including the host and port to use for connection
$ curl -X GET http://$PUBLIC_IP/api/v1/service  -H 'x-api-key: ~api~key~here~'

{
  "services": [
    {
      "id": "9d655cbe-caf1-4104-b8e4-b83fd569b509",
      "service": "redis",
      "configs": {
        "address": "127.0.0.1",
        "password": "",
        "port": "49156"
      },
      "deleteAfter": "10min",
      "createdAt": "2021-07-11T09:58:13Z",
      "updatedAt": "2021-07-11T09:58:13Z"
    },
    {
      "id": "a8138f52-3ebb-4a34-b403-1be6ad481daf",
      "service": "redis",
      "configs": {
        "address": "127.0.0.2",
        "password": "",
        "port": "49155"
      },
      "deleteAfter": "",
      "createdAt": "2021-07-11T09:58:12Z",
      "updatedAt": "2021-07-11T09:58:12Z"
    }
  ]
}
```

There is also a script to upgrade peanut.

```zsh
$ bash < <(curl -s https://raw.githubusercontent.com/Clivern/Peanut/main/deployment/linux/upgrade.sh)
```


#### Linux Deployment Explained

Download [the latest peanut binary](https://github.com/Clivern/Peanut/releases). Make it executable from everywhere.

```zsh
$ export PEANUT_LATEST_VERSION=$(curl --silent "https://api.github.com/repos/Clivern/Peanut/releases/latest" | jq '.tag_name' | sed -E 's/.*"([^"]+)".*/\1/' | tr -d v)

$ curl -sL https://github.com/Clivern/Peanut/releases/download/v{$PEANUT_LATEST_VERSION}/peanut_{$PEANUT_LATEST_VERSION}_Linux_x86_64.tar.gz | tar xz
```

Then install `etcd` cluster or a single node! please refer to etcd docs or bin directory inside this repository.

Install the virtualization runtime! By default peanut uses `docker` & `docker-compose`

```zsh
$ apt-get update
$ apt-get install docker.io -y
$ systemctl enable docker

$ apt-get install docker-compose -y
```

Create the configs file `config.yml` from `config.dist.yml`. Something like the following:

```yaml
# App configs
app:
    # Env mode (dev or prod)
    mode: ${PEANUT_APP_MODE:-dev}
    # HTTP port
    port: ${PEANUT_API_PORT:-8000}
    # Hostname
    hostname: ${PEANUT_API_HOSTNAME:-127.0.0.1}
    # TLS configs
    tls:
        status: ${PEANUT_API_TLS_STATUS:-off}
        pemPath: ${PEANUT_API_TLS_PEMPATH:-cert/server.pem}
        keyPath: ${PEANUT_API_TLS_KEYPATH:-cert/server.key}

    # Containerization runtime (supported docker)
    containerization: ${PEANUT_CONTAINERIZATION_RUNTIME:-docker}

    # App Storage
    storage:
        # Type (only local supported)
        type: ${PEANUT_STORAGE_TYPE:-local}
        # Local Path
        path: ${PEANUT_STORAGE_PATH:-/tmp}

    # API Configs
    api:
        key: ${PEANUT_API_KEY:-6c68b836-6f8e-465e-b59f-89c1db53afca}

    # Async Workers
    workers:
        # Queue max capacity
        buffer: ${PEANUT_WORKERS_CHAN_CAPACITY:-5000}
        # Number of concurrent workers
        count: ${PEANUT_WORKERS_COUNT:-4}

    # Runtime, Requests/Response and Peanut Metrics
    metrics:
        prometheus:
            # Route for the metrics endpoint
            endpoint: ${PEANUT_METRICS_PROM_ENDPOINT:-/metrics}

    # Application Database
    database:
        # Database driver
        driver: ${PEANUT_DB_DRIVER:-etcd}

        # Etcd Configs
        etcd:
            # Etcd database name or prefix
            databaseName: ${PEANUT_DB_ETCD_DB:-peanut}
            # Etcd username
            username: ${PEANUT_DB_ETCD_USERNAME:- }
            # Etcd password
            password: ${PEANUT_DB_ETCD_PASSWORD:- }
            # Etcd endpoints
            endpoints: ${PEANUT_DB_ETCD_ENDPOINTS:-http://127.0.0.1:2379}
            # Timeout in seconds
            timeout: 30

    # Log configs
    log:
        # Log level, it can be debug, info, warn, error, panic, fatal
        level: ${PEANUT_LOG_LEVEL:-info}
        # Output can be stdout or abs path to log file /var/logs/peanut.log
        output: ${PEANUT_LOG_OUTPUT:-stdout}
        # Format can be json
        format: ${PEANUT_LOG_FORMAT:-json}
```

The run the `peanut` with `systemd`

```zsh
$ peanut api -c /path/to/config.yml
```

Deploy your first redis server!

```zsh
$ curl -X POST http://127.0.0.1:8000/api/v1/service -d '{"service":"redis"}' -H 'x-api-key: ~api~key~here~'
```


#### To run the Admin Dashboard (Development Only):

Clone the project or your own fork:

```zsh
$ git clone https://github.com/Clivern/Peanut.git
```

Create the dashboard config file `web/.env` from `web/.env.dist`. Something like the following:

```
VUE_APP_API_URL=http://localhost:8080
```

Then you can either build or run the dashboard

```zsh
# Install npm packages
$ cd web
$ npm install
$ npm install -g npx

# Add api server url to frontend
$ echo "VUE_APP_API_URL=http://127.0.0.1:8000" > .env

$ cd ..

# Validate js code format
$ make check_ui_format

# Format UI
$ make format_ui

# Run Vuejs app
$ make serve_ui

# Build Vuejs app
$ make build_ui

# Any changes to the dashboard, must be reflected to cmd/pkged.go
# You can use these commands to do so
$ go get github.com/markbates/pkger/cmd/pkger
$ make package
```

#### The command line tool

In order to interact with peanut API server, you can either do basic API calls or use the provided command line tool. It is still not finished yet but it will be ready soon.


#### Supported Services

Here is a list of all supported services so far and the API call to deploy them.

- MySQL.

```zsh
$ curl -X POST http://127.0.0.1:8000/api/v1/service \
    -d '{"service":"mysql","configs": {"rootPassword": "root", "database": "peanut", "username": "peanut", "password": "secret"}}' \
    -H 'x-api-key: ~api~key~here~'
```

- MariaDB.

```zsh
$ curl -X POST http://127.0.0.1:8000/api/v1/service \
    -d '{"service":"mariadb","configs": {"rootPassword": "root", "database": "peanut", "username": "peanut", "password": "secret"}}' \
    -H 'x-api-key: ~api~key~here~'
```

- PostgreSQL.

```zsh
$ curl -X POST http://127.0.0.1:8000/api/v1/service \
    -d '{"service":"postgresql","configs": {"database": "peanut", "username": "peanut", "password": "secret"}}' \
    -H 'x-api-key: ~api~key~here~'
```

- Redis.

```zsh
$ curl -X POST http://127.0.0.1:8000/api/v1/service \
    -d '{"service":"redis","configs": {"password": "secret"}}' \
    -H 'x-api-key: ~api~key~here~'
```

- Etcd.

```zsh
$ curl -X POST http://127.0.0.1:8000/api/v1/service \
    -d '{"service":"etcd"}' \
    -H 'x-api-key: ~api~key~here~'
```

- Grafana.

```zsh
$ curl -X POST http://127.0.0.1:8000/api/v1/service \
    -d '{"service":"grafana","configs": {"username": "admin", "password": "admin", "allowSignup": "false", "anonymousAccess": "true"}}' \
    -H 'x-api-key: ~api~key~here~'
```

- Elasticsearch.

```zsh
$ curl -X POST http://127.0.0.1:8000/api/v1/service \
    -d '{"service":"elasticsearch"}' \
    -H 'x-api-key: ~api~key~here~'
```

- MongoDB.

```zsh
$ curl -X POST http://127.0.0.1:8000/api/v1/service \
    -d '{"service":"mongodb","configs": {"database": "peanut", "username": "peanut", "password": "secret"}}' \
    -H 'x-api-key: ~api~key~here~'
```

- Graphite.

```zsh
$ curl -X POST http://127.0.0.1:8000/api/v1/service \
    -d '{"service":"graphite"}' \
    -H 'x-api-key: ~api~key~here~'
```

- Prometheus.

```zsh
$ curl -X POST http://127.0.0.1:8000/api/v1/service \
    -d '{"service":"prometheus"}' \
    -H 'x-api-key: ~api~key~here~'

# Configs can be provided as base64 encoded string (use https://www.base64encode.org/)
$ curl -X POST http://127.0.0.1:8000/api/v1/service \
    -d '{"service":"prometheus", "configs": {"configsBase64Encoded": "Z2xvYmFsOgogIGV2YWx1YXRpb25faW50ZXJ2YWw6IDE1cwogIHNjcmFwZV9pbnRlcnZhbDogMTVzCnJ1bGVfZmlsZXM6IH4Kc2NyYXBlX2NvbmZpZ3M6CiAgLQogICAgam9iX25hbWU6IHByb21ldGhlCiAgICBzY3JhcGVfaW50ZXJ2YWw6IDVzCiAgICBzdGF0aWNfY29uZmlnczoKICAgICAgLQogICAgICAgIHRhcmdldHM6CiAgICAgICAgICAtICJsb2NhbGhvc3Q6OTA5MCI="}}' \
    -H 'x-api-key: ~api~key~here~'
```

- Zipkin.

```zsh
$ curl -X POST http://127.0.0.1:8000/api/v1/service \
    -d '{"service":"zipkin"}' \
    -H 'x-api-key: ~api~key~here~'
```

- Memcached.

```zsh
$ curl -X POST http://127.0.0.1:8000/api/v1/service \
    -d '{"service":"memcached"}' \
    -H 'x-api-key: ~api~key~here~'
```

- Mailhog.

```zsh
$ curl -X POST http://127.0.0.1:8000/api/v1/service \
    -d '{"service":"mailhog"}' \
    -H 'x-api-key: ~api~key~here~'
```

- Jaeger.

```zsh
$ curl -X POST http://127.0.0.1:8000/api/v1/service \
    -d '{"service":"jaeger"}' \
    -H 'x-api-key: ~api~key~here~'
```

- RabbitMQ.

```zsh
$ curl -X POST http://127.0.0.1:8000/api/v1/service \
    -d '{"service":"rabbitmq"}' \
    -H 'x-api-key: ~api~key~here~'
```

- Consul.

```zsh
$ curl -X POST http://127.0.0.1:8000/api/v1/service \
    -d '{"service":"consul"}' \
    -H 'x-api-key: ~api~key~here~'
```

- Vault.

```zsh
$ curl -X POST http://127.0.0.1:8000/api/v1/service \
    -d '{"service":"vault","configs": {"token": "peanut"}}' \
    -H 'x-api-key: ~api~key~here~'
```


To create a temporary service, you will need to add extra parameter while creating it.

```zsh
# It will be deleted after 30 seconds
$ curl -X POST http://127.0.0.1:8000/api/v1/service \
    -d '{"service":"etcd", "deleteAfter": "30sec"}' \
    -H 'x-api-key: ~api~key~here~'

# It will be deleted after 20 minutes
$ curl -X POST http://127.0.0.1:8000/api/v1/service \
    -d '{"service":"etcd", "deleteAfter": "20min"}' \
    -H 'x-api-key: ~api~key~here~'

# It will be deleted after 1 hour
$ curl -X POST http://127.0.0.1:8000/api/v1/service \
    -d '{"service":"etcd", "deleteAfter": "1hours"}' \
    -H 'x-api-key: ~api~key~here~'

# It will be deleted after 3 days
$ curl -X POST http://127.0.0.1:8000/api/v1/service \
    -d '{"service":"etcd", "deleteAfter": "3days"}' \
    -H 'x-api-key: ~api~key~here~'
```

To list all services running on a single or multi nodes.

```zsh
$ curl -X GET http://127.0.0.1:8000/api/v1/service \
    -H 'x-api-key: ~api~key~here~'
```

To delete a service.

```zsh
$ curl -X DELETE http://127.0.0.1:8000/api/v1/service/:serviceId \
    -H 'x-api-key: ~api~key~here~'
```

To get async job status like a deployment status.

```zsh
$ curl -X DELETE http://127.0.0.1:8000/api/v1/job/:serviceId/:jobId \
    -H 'x-api-key: ~api~key~here~'
```

To get service versions.

```zsh
$ curl -X GET http://127.0.0.1:8000/api/v1/tag/$serviceType/$fromCacheStatus \
    -H 'x-api-key: ~api~key~here~'

$ curl -X GET http://127.0.0.1:8000/api/v1/tag/mysql/true \
    -H 'x-api-key: ~api~key~here~'
```


## Versioning

For transparency into our release cycle and in striving to maintain backward compatibility, Peanut is maintained under the [Semantic Versioning guidelines](https://semver.org/) and release process is predictable and business-friendly.

See the [Releases section of our GitHub project](https://github.com/clivern/peanut/releases) for changelogs for each release version of Peanut. It contains summaries of the most noteworthy changes made in each release.


## Bug tracker

If you have any suggestions, bug reports, or annoyances please report them to our issue tracker at https://github.com/clivern/peanut/issues


## Security Issues

If you discover a security vulnerability within Peanut, please send an email to [hello@clivern.com](mailto:hello@clivern.com)


## Contributing

We are an open source, community-driven project so please feel free to join us. see the [contributing guidelines](CONTRIBUTING.md) for more details.


## License

© 2021, Clivern. Released under [MIT License](https://opensource.org/licenses/mit-license.php).

**Peanut** is authored and maintained by [@clivern](http://github.com/clivern).
