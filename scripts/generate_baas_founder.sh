#!/usr/bin/env bash

set -e


export binSource=BaasFounder_bin
export abiSource=BaasFounder.json

abigen --abi $abiSource --bin $binSource --pkg contracts --type BaasFounder --out ../deployer/contracts/baas_founder.go

cp $abiSource $destEtherlytics
cp $abiSource $destWebapp

rm $abiSource
rm $binSource


