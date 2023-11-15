#!/bin/bash

psql -U "${POSTGRES_USER}" -d postgres -h "${POSTGRES_DB_HOST}" -p "${POSTGRES_DB_PORT}" -c 'CREATE DATABASE "fixtures"'
