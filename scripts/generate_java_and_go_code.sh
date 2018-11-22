#!/usr/bin/env bash

set -e

node parse_abi.js

export binSource=bin
export abiSource=abi.json
export destination=../../android/app/src/main/java
export package=com.wire.poc.data.web3.contracts
export destinationEtherlytics=../../etherlytics/dapp

web3j solidity generate $binSource $abiSource -o $destination  -p $package


abigen --abi $abiSource --bin $binSource --pkg contract --type UserRegistryContract --out ../build/contract/user_registry_contract.go

cp ../build/contract/user_registry_contract.go ../../server/app/contract

cp abi.json $destinationEtherlytics

rm abi.json
rm bin