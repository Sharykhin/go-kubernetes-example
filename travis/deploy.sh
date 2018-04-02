#!/bin/bash

if ([ $TRAVIS_BRANCH == "master" ] && [ $TRAVIS_PULL_REQUEST == "false" ])
then
  VERSION=$(git log -n 1 master --pretty=format:"%H")
  echo 'Website deployed on ${VERSION}'
else
  echo "Build successful, but not publishing!"
fi
