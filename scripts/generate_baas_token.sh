#!/usr/bin/env bash

set -e


export binSource=BaasToken_bin
export abiSource=BaasToken_abi.json
export destinationEtherlytics=../../tge-etherlytics/dapp

abigen --abi $abiSource --bin $binSource --pkg contracts --type BaasToken --out ../deployer/contracts/baas_token.go

cp $abiSource $destinationEtherlytics

rm $abiSource
rm $binSource


