#!/usr/bin/env bash

set -e


export binSource=BaasEscrow_bin
export abiSource=BaasEscrow.json


abigen --abi $abiSource --bin $binSource --pkg contracts --type BaasEscrow --out ../deployer/contracts/baas_escrow.go

cp $abiSource $destEtherlytics
cp $abiSource $destWebapp

rm $abiSource
rm $binSource


