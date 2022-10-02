#!/bin/sh
#
# Deploy to heroku in another directory
#
set -e

SCRIPTS=$(dirname $0)

DEPLOY_DIR=../opentable-deploy
mkdir -p $DEPLOY_DIR
mkdir -p $DEPLOY_DIR/scripts
cp $SCRIPTS/deploy.sh.txt $DEPLOY_DIR/scripts/deploy-to-keroku.sh

pushd $DEPLOY_DIR
scripts/deploy-to-keroku.sh
popd