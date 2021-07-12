#!/bin/bash

function peanut {
    echo "Upgrade peanut ..."

    cd /etc/peanut
    mv config.prod.yml config.back.yml

    PEANUT_LATEST_VERSION=$(curl --silent "https://api.github.com/repos/Clivern/Peanut/releases/latest" | jq '.tag_name' | sed -E 's/.*"([^"]+)".*/\1/' | tr -d v)

    curl -sL https://github.com/Clivern/Peanut/releases/download/v{$PEANUT_LATEST_VERSION}/peanut_{$PEANUT_LATEST_VERSION}_Linux_x86_64.tar.gz | tar xz

    rm config.prod.yml
    mv config.back.yml config.prod.yml

    systemctl restart peanut

    echo "peanut upgrade done!"
}

peanut
