#!/usr/bin/env bash

set -e


export binSource=BaasPP_bin
export abiSource=BaasPP.json

abigen --abi $abiSource --bin $binSource --pkg contracts --type BaasPP --out ../deployer/contracts/baas_pp.go

cp $abiSource $destEtherlytics
cp $abiSource $destWebapp

rm $abiSource
rm $binSource


