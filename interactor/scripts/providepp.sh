#!/usr/bin/env bash

# Example  .scripts/providepp.sh 0x80dEF17bb76025EaBa069dE8D5eBfd8FbFc64D51 300 v0.0.2

set -e

if [ -z "$1" ]
  then
    echo "no target supplied"
    exit 1
fi

if [ -z "$2" ]
  then
    echo "no amount supplied"
    exit 1
fi

if [ -z "$3" ]
  then
    echo "no version supplied"
    exit 1
fi

go run main.go e pp p -target $1 -amount $2 -version $3

#go run main.go e pp p -target 0xc4a5096a2785761dc7c24861d95d1a5cd33b7234 -amount 100000 -version v0.0.10 -discounted