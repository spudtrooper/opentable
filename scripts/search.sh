#!/bin/sh

scripts=$(dirname $0)

uris=$(echo "$@" | sed -e 's/ /,/g')
${scripts}/run.sh AddRestaurantsToSearchByURIs SearchEmptyRestaurants --verbose --uri "$uris"