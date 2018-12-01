#!/usr/bin/env bash

set  -e
truffle compile
pushd scripts
./generate_code.sh
