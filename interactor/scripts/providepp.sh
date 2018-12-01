#!/usr/bin/env bash

# Example  ./providepp.sh 0x80dEF17bb76025EaBa069dE8D5eBfd8FbFc64D51 300 v0.0.2

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