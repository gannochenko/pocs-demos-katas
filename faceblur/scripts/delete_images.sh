#!/bin/bash

PGPASSWORD=${POSTGRES_PASSWORD} psql -U ${POSTGRES_USER} -d ${POSTGRES_DB} -h localhost -p ${POSTGRES_DB_PORT} -c 'DELETE from "images"';
