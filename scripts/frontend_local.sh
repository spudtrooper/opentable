#!/bin/sh
#
# Runs the frontend server locally.
#
set -e

scripts="$(dirname "$0")"

go mod tidy
go run frontend_main.go --port_for_testing 8080 --host localhost "$@"