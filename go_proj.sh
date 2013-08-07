#!/bin/sh
mkdir $1
mkdir $1/bin
mkdir $1/src
mkdir $1/pkg
git add .
git commit -m "$1 init"
git push
