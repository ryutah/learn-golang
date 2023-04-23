#!/usr/bin/env bash

set -eux

podman container run -it --rm -p 5432:5432 -e POSTGRES_PASSWORD=psql postgres:15.2
