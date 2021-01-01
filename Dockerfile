FROM golang:1.16.5

ARG PEANUT_VERSION=0.1.9

ENV GO111MODULE=on

RUN mkdir -p /app/configs
RUN mkdir -p /app/var/logs
RUN mkdir -p /app/var/storage
RUN apt-get update

WORKDIR /app

RUN curl -sL https://github.com/Clivern/Peanut/releases/download/v${PEANUT_VERSION}/peanut_${PEANUT_VERSION}_Linux_x86_64.tar.gz | tar xz
RUN rm LICENSE
RUN rm README.md

COPY ./config.dist.yml /app/configs/

EXPOSE 8000

VOLUME /app/configs
VOLUME /app/var

RUN ./peanut version

CMD ["./peanut", "api", "-c", "/app/configs/config.dist.yml"]