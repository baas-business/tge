#!/usr/bin/env bash

set -e

node parse_abi.js


export destWebapp=../../tge-webapp/src/data/Web3/Contracts
export destEtherlytics=../../tge-etherlytics/dapp

./generate_baas_token.sh
./generate_baas_escrow.sh
./generate_baas_founder.sh
./generate_baas_incentives.sh
./generate_baas_pp.sh
./generate_baas_roi.sh