#!/usr/bin/env bash

set -eux

podman container run -e PGPASSWORD=psql -it --rm --net host postgres:15.2 psql -h 127.0.0.1 -p 5432 -U postgres -c 'CREATE DATABASE world'
podman container run -e PGPASSWORD=psql -it -v $(pwd):/db --rm --net host postgres:15.2 sh -c "psql -h 127.0.0.1 -p 5432 -U postgres -d world < /db/world.sql"
