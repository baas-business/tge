#!/usr/bin/env bash

set -e


generate_go_code(){
    name=$1
    goname=$2
    binSource=${name}_bin
    abiSource=${name}.json

    abigen --abi $abiSource --bin $binSource --pkg contracts --type ${name} --out ../interactor/contracts/${goname}

    rm $abiSource
    rm $binSource
}

# parse abi from truffle build
node parse_abi.js


generate_go_code BaasToken baas_token.go
generate_go_code BaasEscrow baas_escrow.go
generate_go_code BaasFounder baas_founder.go
generate_go_code BaasIncentives baas_incentives.go
generate_go_code BaasPP baas_pp.go
generate_go_code BaasROI baas_roi.go