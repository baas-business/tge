#!/usr/bin/env bash

pushd interactor
version="$(go run main.go c v)"
echo $version
popd