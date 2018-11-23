#!/usr/bin/env bash

set -e


export binSource=BaasROI_bin
export abiSource=BaasROI_abi.json
export destinationEtherlytics=../../tge-etherlytics/dapp

abigen --abi $abiSource --bin $binSource --pkg contracts --type BaasROI --out ../deployer/contracts/baas_roi.go

cp $abiSource $destinationEtherlytics

rm $abiSource
rm $binSource


