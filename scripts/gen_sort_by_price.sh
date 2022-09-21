#!/bin/sh
#
# Creates html/sort-by-price/index.md
#
set -e 

mkdir -p output/sort-by-price
go mod tidy
go run tools/createsortbyprice.go > output/sort-by-price/index.md