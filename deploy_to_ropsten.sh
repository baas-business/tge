#!/usr/bin/env bash

set  -e


if [ $# -eq 0 ]
  then
    echo "No destination supplied"
    exit 1
fi

version=$1

truffle compile


pushd scripts
./copy_config.sh
popd

pushd interactor
# deploy contracts
go run main.go e d -version $version
# setup BaasToken
go run main.go e s -version $version
popd

pushd scripts
./copy_config.sh
popd


git add -A
git commit -a -m "deployed $version to ropsten"
git tag -a $version -m 'ropsten deployment'
git push origin $version