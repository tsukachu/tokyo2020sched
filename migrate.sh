#!/bin/sh

psql -d tokyo2020 -h 127.0.0.1 -U user1 -f compose/db/migrate.sql
psql -d tokyo2020 -h 127.0.0.1 -U user1 -c "\COPY competitions (id, name) FROM 'compose/db/master/competitions.csv' CSV HEADER"
