#!/usr/bin/env bash

set -e


export binSource=BaasFounder_bin
export abiSource=BaasFounder_abi.json
export destinationEtherlytics=../../tge-etherlytics/dapp

abigen --abi $abiSource --bin $binSource --pkg contracts --type BaasFounder --out ../deployer/contracts/baas_founder.go

cp $abiSource $destinationEtherlytics

rm $abiSource
rm $binSource


