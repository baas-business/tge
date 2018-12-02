#!/usr/bin/env bash

set  -e


if [ $# -eq 0 ]
  then
    echo "No destination supplied"
    exit 1
fi

version=$1

echo "Compiling Contracts"
truffle compile

echo "Generating Code"
pushd scripts
./generate_code.sh
popd

pushd interactor
# deploy contracts
echo "Deploying Contracts"
go run main.go e d -version $version
# setup BaasToken
echo "Setting Up Baas Token"
go run main.go e s -version $version
popd

echo "Copying configuration to web3 app and Etherlytics"
pushd scripts
./copy_config.sh $version
popd

echo "Committing and Pushing"
git add -A
git commit -a -m "deployed $version to ropsten"
git tag -a $version -m 'ropsten deployment'
git push origin $version