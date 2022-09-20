#!/bin/sh
#
# Exports DB as JSON.
#
set -e

mkdir -p data
mkdir -p data/scripts

mongoexport --db=opentable --collection=restaurants --type=json > data/restaurants.json

cat playgrounds/sort-by-price.mongodb | sed 's#use(#//use(#g' > data/scripts/sort-by-price.js
mongo mongodb://localhost:27017/seats data/scripts/sort-by-price.js
mongoexport --db=opentable --collection=sortedByPrice --type=json --fields=restaurantName,restaurantWebsite,title,price > data/sortedByPrice.json