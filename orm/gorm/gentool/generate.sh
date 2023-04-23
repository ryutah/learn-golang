#!/usr/bin/env bash

set -eux

gentool \
  -db "postgres" \
  -dsn "user=postgres password=psql host=127.0.0.1 port=5432 dbname=world sslmode=disable" \
  -outPath "./gen/query" \
  -modelPkgName "gen/model" \
  -fieldNullable
