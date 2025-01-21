#!/bin/bash

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

rm -rf ${DIR}/../.data/storage/faceblur-images/*
PGPASSWORD=${POSTGRES_PASSWORD} psql -U ${POSTGRES_USER} -d ${POSTGRES_DB} -h localhost -p ${POSTGRES_DB_PORT} -c 'DELETE from "images"';
