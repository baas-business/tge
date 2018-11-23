#!/usr/bin/env bash

set -e


export binSource=BaasEscrow_bin
export abiSource=BaasEscrow_abi.json
export destinationEtherlytics=../../tge-etherlytics/dapp

abigen --abi $abiSource --bin $binSource --pkg contracts --type BaasEscrow --out ../deployer/contracts/baas_escrow.go

cp $abiSource $destinationEtherlytics

rm $abiSource
rm $binSource


