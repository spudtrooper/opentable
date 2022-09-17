#!/bin/sh
#
# Persists the user's token from the cookie.
#
set -e

go mod tidy
go run tools/usecookie.go "$@"