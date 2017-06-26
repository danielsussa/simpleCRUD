#!/usr/bin/env bash

git checkout master || exit
git merge "$TRAVIS_COMMIT" || exit
git push ... # here need some authorization and url