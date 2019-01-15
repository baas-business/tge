#!/usr/bin/env bash

set  -e



total=6

echo "1/$total Compiling Contracts"
truffle compile

echo "2/$total Generating Code"
pushd scripts
./generate_code.sh
popd

pushd interactor
# deploy contracts
echo "3/$total Deploying Contracts"
go run main.go d
# setup BaasToken
echo "4/$total Setting Up Baas Token"
go run main.go e t s
echo "5/$total Setting Up Baas Escrow"
#go run main.go e e s -vesting 1000

version="$(go run main.go c v)"
popd

echo "6/$total Copying configuration to web3 app and Etherlytics"


pushd scripts
./copy_config.sh $version
popd