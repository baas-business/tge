#!/usr/bin/env bash

set -e


export binSource=BaasROI_bin
export abiSource=BaasROI.json

abigen --abi $abiSource --bin $binSource --pkg contracts --type BaasROI --out ../deployer/contracts/baas_roi.go

cp $abiSource $destEtherlytics
cp $abiSource $destWebapp

rm $abiSource
rm $binSource


