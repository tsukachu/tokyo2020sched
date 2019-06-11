#!/bin/sh

psql -d tokyo2020sched -h 127.0.0.1 -U user1 -f compose/db/migrate.sql
psql -d tokyo2020sched -h 127.0.0.1 -U user1 -c "\COPY competition (id, name) FROM 'compose/db/master/competitions.csv' CSV HEADER"
psql -d tokyo2020sched -h 127.0.0.1 -U user1 -c "\COPY classification (id, name, competition_id) FROM 'compose/db/master/classifications.csv' CSV HEADER"
psql -d tokyo2020sched -h 127.0.0.1 -U user1 -c "\COPY place (id, name) FROM 'compose/db/master/places.csv' CSV HEADER"
psql -d tokyo2020sched -h 127.0.0.1 -U user1 -c "\COPY olympic_schedule (id, competition_id, classification_id, title, \"begin\", \"end\", place_id, content) FROM 'compose/db/master/olympic_schedules.csv' CSV HEADER"
