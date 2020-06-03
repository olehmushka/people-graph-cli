#!/bin/bash

if [ -z $1 ]; then
  echo -e "Please provide command"
fi

BASEDIR=$(dirname "$0")

if [ "$1" == "migrate-locations" ]; then
  cd "${BASEDIR}"/migrate-locations/src/
  GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go install -ldflags="-w -s"
  go build .
  ./main.com
fi