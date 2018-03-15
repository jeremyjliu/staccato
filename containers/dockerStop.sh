#!/bin/bash

DIR="$(dirname "$0")"
docker-compose -f "$DIR/docker-postgres.yml" stop