#!/usr/bin/env bash

set -e

node parse_abi.js

./generate_baas_token.sh
./generate_baas_escrow.sh
./generate_baas_founder.sh
./generate_baas_incentive.sh
./generate_baas_pp.sh
./generate_baas_roi.sh