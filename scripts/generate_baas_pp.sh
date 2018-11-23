#!/usr/bin/env bash

set -e


export binSource=BaasPP_bin
export abiSource=BaasPP_abi.json
export destinationEtherlytics=../../tge-etherlytics/dapp

abigen --abi $abiSource --bin $binSource --pkg contracts --type BaasPP --out ../deployer/contracts/baas_pp.go

cp $abiSource $destinationEtherlytics

rm $abiSource
rm $binSource


