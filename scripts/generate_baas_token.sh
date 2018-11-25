#!/usr/bin/env bash

set -e


export binSource=BaasToken_bin
export abiSource=BaasToken.json

abigen --abi $abiSource --bin $binSource --pkg contracts --type BaasToken --out ../deployer/contracts/baas_token.go

cp $abiSource $destEtherlytics
cp $abiSource $destWebapp

rm $abiSource
rm $binSource


