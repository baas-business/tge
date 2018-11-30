#!/usr/bin/env bash


name=$1
goname=$2
export binSource=${name}_bin
export abiSource=${name}.json
export source=../contracts/tge/${name}.sol

abigen --abi $abiSource --bin $binSource --pkg contracts --type ${name} --out ../deployer/contracts/${goname}

cp $abiSource $destEtherlytics
cp $source $destEtherlytics
cp $abiSource $destWebapp


rm $abiSource
rm $binSource
