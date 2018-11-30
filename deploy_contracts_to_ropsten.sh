#!/usr/bin/env bash

set  -e


if [ $# -eq 0 ]
  then
    echo "No destination supplied"
    exit 1
fi

truffle compile


pushd interactor
go run main.go deploy -version $1

