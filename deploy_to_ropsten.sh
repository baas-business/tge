#!/usr/bin/env bash

set  -e


if [ $# -eq 0 ]
  then
    echo "no version supplied"
    exit 1
fi

version=$1

echo "1/6 Compiling Contracts"
truffle compile

echo "2/6 Generating Code"
pushd scripts
./generate_code.sh
popd

pushd interactor
# deploy contracts
echo "3/6 Deploying Contracts"
go run main.go e d -version $version
# setup BaasToken
echo "4/6 Setting Up Baas Token"
go run main.go e s -version $version
popd

echo "5/6 Copying configuration to web3 app and Etherlytics"
pushd scripts
./copy_config.sh $version
popd

echo "6/6 Committing and Pushing"
git add -A
git commit -a -m "deployed $version to ropsten"
git tag -a $version -m 'ropsten deployment'
git push origin $version