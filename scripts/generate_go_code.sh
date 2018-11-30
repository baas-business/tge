#!/usr/bin/env bash

set -e

node parse_abi.js


export destWebapp=../../tge-webapp/src/data/Web3/Contracts
export destWebappH=../../tge-webapp/src/data/Web3
export destEtherlytics=../../tge-etherlytics/dapp

export version="v4"

cp ../deployer/dapp_${version}.json ${destEtherlytics}/dapp.json
cp ../deployer/contract_config_${version}.json ${destWebappH}/contract_config.json

./generate.sh  BaasToken baas_token.go
./generate.sh BaasEscrow baas_escrow.go
./generate.sh BaasFounder baas_founder.go
./generate.sh BaasIncentives baas_incentives.go
./generate.sh BaasPP baas_pp.go
./generate.sh BaasROI baas_roi.go