#!/usr/bin/env bash

git add .
git status -s | grep services | grep -v $1 | grep A | awk '{print $2}' | xargs -I {} rm -f {}
git status -s | grep services | grep -v $1 | grep M | awk '{print $2}' | xargs -I {} git restore --staged {}
git status -s | grep services | grep -v $1 | grep M | awk '{print $2}' | xargs -I {} git checkout -- {}
git add .
