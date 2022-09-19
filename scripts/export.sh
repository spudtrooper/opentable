#!/bin/sh
#
# Exports DB as JSON.
#
set -e

mkdir -p data
# mongoexport --db=opentable --collection=restaurants --type=json > data/restaurants.json
mongoexport --db=opentable --collection=sortedByPrice --type=json --fields=restaurantName,restaurantWebsite,title,price > data/sortedByPrice.json