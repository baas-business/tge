#!/usr/bin/env bash

set -e


export binSource=BaasIncentives_bin
export abiSource=BaasIncentives.json

abigen --abi $abiSource --bin $binSource --pkg contracts --type BaasIncentive --out ../deployer/contracts/baas_incentive.go

cp $abiSource $destEtherlytics
cp $abiSource $destWebapp

rm $abiSource
rm $binSource


