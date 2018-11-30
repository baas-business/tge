#!/usr/bin/env bash

set -e

if [ $# -eq 0 ]
  then
    echo "no version supplied"
    exit 1
fi


destWebapp=../../tge-webapp/src/data/Web3/Contracts
destWebappH=../../tge-webapp/src/data/Web3
destEtherlytics=../../tge-etherlytics/dapp
version=$1



destWebapp=../../tge-webapp/src/data/Web3/Contracts
destWebappH=../../tge-webapp/src/data/Web3
destEtherlytics=../../tge-etherlytics/dapp
version=$1


# parse abi from truffle build
node parse_abi.js

copy_contract(){
    name=$1

    binSource=${name}_bin
    abiSource=${name}.json
    source=../contracts/tge/${name}.sol

    cp $abiSource $destEtherlytics
    cp $source $destEtherlytics
    cp $abiSource $destWebapp

    rm $abiSource
    rm $binSource
}

# parse abi from truffle build
node parse_abi.js


copy_contract BaasToken
copy_contract BaasEscrow
copy_contract BaasFounder
copy_contract BaasIncentives
copy_contract BaasPP
copy_contract BaasROI

cp ../interactor/build/${version}/dapp.json ${destEtherlytics}/dapp.json
cp ../interactor/build/${version}/contract_config.json ${destWebappH}/contract_config.json