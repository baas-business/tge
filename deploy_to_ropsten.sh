#!/usr/bin/env bash

set  -e


if [ $# -eq 0 ]
  then
    echo "no version supplied"
    exit 1
fi

version=$1
total=7
current=0

echo "1/$total Compiling Contracts"
truffle compile

echo "2/$total Generating Code"
pushd scripts
./generate_code.sh
popd

pushd interactor
# deploy contracts
echo "3/$total Deploying Contracts"
go run main.go e d -version $version
# setup BaasToken
echo "4/$total Setting Up Baas Token"
go run main.go e s -version $version
echo "5/$total Setting Up Baas Token"
go run main.go e i s -version $version
popd

echo "6/$total Copying configuration to web3 app and Etherlytics"
pushd scripts
./copy_config.sh $version
popd

echo "7/$total Committing and Pushing"
git add -A
git commit -a -m "deployed $version to ropsten"
git tag -a $version -m 'ropsten deployment'
git push origin $version