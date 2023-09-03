#!/usr/bin/env bash

set -eu

gcloud builds submit --pack image="gcr.io/$(gcloud config get project 2>/dev/null)/sample-go"
