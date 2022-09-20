#!/bin/sh
#
# Serves with local HTTP server.
#
set -e

echo "open http://localhost:8000"
python -m http.server --directory html