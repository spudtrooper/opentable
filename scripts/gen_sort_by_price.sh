#!/bin/sh
#
# Creates html/sort-by-price/index.md
#
set -e 

go mod tidy
go run createsortbyprice.go > html/sort-by-price/index.md