#!/usr/bin/env bash

set -e


export binSource=BaasIncentive_bin
export abiSource=BaasIncentive_abi.json
export destinationEtherlytics=../../tge-etherlytics/dapp

abigen --abi $abiSource --bin $binSource --pkg contracts --type BaasIncentive --out ../deployer/contracts/baas_incentive.go

cp $abiSource $destinationEtherlytics

rm $abiSource
rm $binSource


