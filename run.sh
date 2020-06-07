#!/bin/bash

function migrate_locations {
  GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go install -ldflags="-w -s"
  go build .
  ./main.com
}

function migrate_country {
  COUNTRY_NAME="$1"
  COUNTRY_CODE="$2"
  COUNTRY_NAME="$COUNTRY_NAME" COUNTRY_CODE="$COUNTRY_CODE" GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go install -ldflags="-w -s"
  COUNTRY_NAME="$COUNTRY_NAME" COUNTRY_CODE="$COUNTRY_CODE" go build .
  COUNTRY_NAME="$COUNTRY_NAME" COUNTRY_CODE="$COUNTRY_CODE" ./main.com
}

if [ -z $1 ]; then
  echo -e "Please provide command"
fi

BASEDIR=$(dirname "$0")

if [ "$1" == "migrate-locations" ]; then
  cd "${BASEDIR}"/go-cli/src/migrate-locations/
  migrate_locations
  cd ../migrate-country/
  migrate_country "United States" "US"
  migrate_country "United Kingdom" "UK"
fi
