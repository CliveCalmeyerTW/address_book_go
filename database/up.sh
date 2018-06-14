#!/usr/bin/env bash

source vars.sh

MIGPATH=filesystem:./migrations

docker run \
    --name ${DBNAME} \
    --env POSTGRES_USER=${DBUSER} \
    --env POSTGRES_PASSWORD=${DBPASS} \
    --env POSTGRES_DB=${DBNAME} \
    --detach \
    --publish ${DBPORT}:${DBPORT} \
    postgres

sleep 5

flyway \
    -baselineOnMigrate=true \
    -user=${DBUSER} \
    -password=${DBPASS} \
    -locations=${MIGPATH} \
    -url=jdbc:postgresql://${DBHOST}/${DBNAME} \
    migrate