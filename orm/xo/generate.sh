#!/usr/bin/env bash

set -eux

xo schema \
  --go-field-tag='json:"{{ .SQLName }}" db:"{{ .SQLName }}"' \
  postgres://postgres:psql@127.0.0.1:5432/world?sslmode=disable
