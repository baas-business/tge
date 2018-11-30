#!/usr/bin/env bash

set  -e


if [ $# -eq 0 ]
  then
    echo "No destination supplied"
    exit 1
fi

version=$1

truffle compile


pushd interactor
go run main.go deploy -version $version
popd

git add -A
git commit -a -m "deployed $version to ropsten"
git tag -a $version -m 'ropsten deployment'
git push origin $version