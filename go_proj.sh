#!/bin/sh
mkdir $1
cat "$1 init" >> $1/README.md
mkdir $1/bin
touch $1/bin/init
mkdir $1/src
touch $1/src/$1.go
touch $1/src/$1_test.go
mkdir $1/pkg
touch $1/pkg/init
git add .
git commit -m "$1 init"
git push
