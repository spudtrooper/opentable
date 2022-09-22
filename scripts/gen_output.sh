#!/bin/sh
#
# Creates html/sort-by-price/index.md
#
set -e 

go mod tidy

mkdir -p output/sort-by-price
go run tools/createsortbyprice.go > output/sort-by-price/index.md

mkdir -p output/menu-item-histogram
go run tools/menuitemhistogram.go > output/menu-item-histogram/index.md