#!/bin/bash

function docker {
    echo "Installing docker ..."
    apt-get update
    apt-get install docker.io -y
    systemctl enable docker
    echo "docker installation done!"
}

function docker_compose {
    echo "Installing docker-compose ..."
    apt-get install docker-compose -y
    echo "docker-compose installation done!"
}

function etcd {
    echo "Installing etcd ..."

    ETCD_VER=v3.4.14

    GOOGLE_URL=https://storage.googleapis.com/etcd
    GITHUB_URL=https://github.com/etcd-io/etcd/releases/download
    DOWNLOAD_URL=${GOOGLE_URL}

    rm -f /tmp/etcd-${ETCD_VER}-linux-amd64.tar.gz
    rm -rf /tmp/etcd-download-test && mkdir -p /tmp/etcd-download-test

    curl -L ${DOWNLOAD_URL}/${ETCD_VER}/etcd-${ETCD_VER}-linux-amd64.tar.gz -o /tmp/etcd-${ETCD_VER}-linux-amd64.tar.gz
    tar xzvf /tmp/etcd-${ETCD_VER}-linux-amd64.tar.gz -C /tmp/etcd-download-test --strip-components=1
    rm -f /tmp/etcd-${ETCD_VER}-linux-amd64.tar.gz

    /tmp/etcd-download-test/etcd --version
    /tmp/etcd-download-test/etcdctl version

    cp /tmp/etcd-download-test/etcd /usr/local/bin/
    cp /tmp/etcd-download-test/etcdctl /usr/local/bin/

    mkdir -p /var/lib/etcd/
    mkdir /etc/etcd

    groupadd --system etcd
    useradd -s /sbin/nologin --system -g etcd etcd

    chown -R etcd:etcd /var/lib/etcd/

    echo "[Unit]
Description=Etcd KV Store
Documentation=https://github.com/etcd-io/etcd
After=network.target

[Service]
User=etcd
Type=notify
Environment=ETCD_DATA_DIR=/var/lib/etcd
Environment=ETCD_NAME=%m
ExecStart=/usr/local/bin/etcd --listen-client-urls http://0.0.0.0:2379 --advertise-client-urls http://0.0.0.0:2379
Restart=always
RestartSec=10s
LimitNOFILE=40000

[Install]
WantedBy=multi-user.target" > /etc/systemd/system/etcd.service

    systemctl daemon-reload
    systemctl start etcd.service

    echo "etcd installation done!"
}

function peanut {
    echo "Installing peanut ..."

    apt-get install jq -y

    mkdir -p /etc/peanut

    cd /etc/peanut

    PEANUT_LATEST_VERSION=$(curl --silent "https://api.github.com/repos/Clivern/Peanut/releases/latest" | jq '.tag_name' | sed -E 's/.*"([^"]+)".*/\1/' | tr -d v)

    curl -sL https://github.com/Clivern/Peanut/releases/download/v{$PEANUT_LATEST_VERSION}/peanut_{$PEANUT_LATEST_VERSION}_Linux_x86_64.tar.gz | tar xz


    echo "[Unit]
Description=Peanut
Documentation=https://github.com/clivern/peanut

[Service]
ExecStart=/etc/peanut/peanut api -c /etc/peanut/config.dist.yml
Restart=on-failure
RestartSec=2

[Install]
WantedBy=multi-user.target" > /etc/systemd/system/peanut.service

    systemctl daemon-reload
    systemctl start peanut.service

    echo "peanut installation done!"
}

docker
docker_compose
etcd
peanut
