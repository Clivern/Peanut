## Deploy with Docker

Install docker & docker-compose

```bash
$ apt-get update
$ sudo apt install docker.io
$ sudo systemctl enable docker
$ sudo apt install docker-compose
```

Clone Peanut for docker-compose.yml file

```bash
$ git clone https://github.com/Clivern/Peanut.git peanut
$ cd peanut/deployment/docker-compose
```

Feel free to update peanut tower port, `api.key` and `api.encryptionKey`. make sure you also use these values in peanut agent config file since agents require the `tower URL`, tower `API key` and `encryptionKey` to be able to reach and communicate with peanut tower.

Run the tower and etcd. It is also recommended to run etcd anywhere where data loss is mitigated.

```bash
$ docker-compose up -d
```

Now tower should be running. User your server public IP and tower port configured before to open the dashboard and setup the admin account.

```bash
# To get the public IP
$ curl https://ipinfo.io/ip
```

In the host where backups have to take place, download peanut binary.

```bash
$ curl -sL https://github.com/Clivern/Peanut/releases/download/v1.1.0/peanut_1.1.0_Linux_x86_64.tar.gz | tar xz
```

Create agent config file. Don't forget to replace `agent.tower` configs with the `tower URL`, `apiKey` and `encryptionKey`, you can get these values from tower configs you created earlier.

```yaml
# Agent configs
agent:
    # Env mode (dev or prod)
    mode: ${PEANUT_APP_MODE:-prod}
    # HTTP port
    port: ${PEANUT_API_PORT:-8001}
    # URL
    url: ${PEANUT_API_URL:-http://127.0.0.1:8001}
    # TLS configs
    tls:
        status: ${PEANUT_API_TLS_STATUS:-off}
        pemPath: ${PEANUT_API_TLS_PEMPATH:-cert/server.pem}
        keyPath: ${PEANUT_API_TLS_KEYPATH:-cert/server.key}

    # API Configs
    api:
        key: ${PEANUT_API_KEY:-56e1a911-cc64-44af-9c5d-8c7e72ec96a1}

    # Async Workers
    workers:
        # Queue max capacity
        buffer: ${PEANUT_WORKERS_CHAN_CAPACITY:-5000}
        # Number of concurrent workers
        count: ${PEANUT_WORKERS_COUNT:-4}

    # Tower Configs
    tower:
        url: ${PEANUT_TOWER_URL:-http://127.0.0.1:8000}
        # This must match the one defined in tower config file
        apiKey: ${PEANUT_TOWER_API_KEY:-6c68b836-6f8e-465e-b59f-89c1db53afca}
        # This must match the one defined in tower config file
        encryptionKey: ${PEANUT_ENCRYPTION_KEY:-B?E(H+Mb}
        # Time interval between agent ping checks
        pingInterval: ${PEANUT_CHECK_INTERVAL:-60}

    # Backup settings
    backup:
        tmpDir: ${PEANUT_BACKUP_TMP_DIR:-/tmp}

    # Log configs
    log:
        # Log level, it can be debug, info, warn, error, panic, fatal
        level: ${PEANUT_LOG_LEVEL:-info}
        # output can be stdout or abs path to log file /var/logs/peanut.log
        output: ${PEANUT_LOG_OUTPUT:-stdout}
        # Format can be json
        format: ${PEANUT_LOG_FORMAT:-json}
```

Then run the host agent.

```
$ peanut agent -c /path/to/agent.config.yml
```

If everything is right, you should be able to see the host shown in the tower dashboard with one active agent. You can create backup crons under that host and update s3 configs in `settings` tab.

