#!/bin/sh
#
# Exports DB as JSON.
#
set -e

mkdir -p data
mkdir -p data/scripts

mongoexport --db=opentable --collection=restaurants --type=json > data/restaurants.json

cat playgrounds/sort_by_price.mongodb | sed 's#use(#//use(#g' > data/scripts/sort_by_price.js
mongo mongodb://localhost:27017/seats data/scripts/sort_by_price.js
mongoexport --db=opentable --collection=sortedByPrice --type=json --fields=restaurantName,restaurantWebsite,title,price > data/sortedByPrice.json