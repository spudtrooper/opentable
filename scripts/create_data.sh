#!/bin/sh
#
# Creates html/sort-by-price/data.js from data/sortedByPrice.json
#
set -e 

out=html/sort-by-price/data.js
rm -f $out
echo "window.data = [" >> $out
head -n 1000 data/sortedByPrice.json | sed 's/\}$/},/' >> $out
echo "];" >> $out

echo "wrote to $out"