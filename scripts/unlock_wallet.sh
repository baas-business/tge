#!/usr/bin/env bash
curl --data '{"method":"personal_unlockAccount","params":["0x80def17bb76025eaba069de8d5ebfd8fbfc64d51","hunter2",null],"id":1,"jsonrpc":"2.0"}' -H "Content-Type: application/json" -X POST http://35.196.119.92:8545

curl --data '{"method":"personal_unlockAccount","params":["0x80def17bb76025eaba069de8d5ebfd8fbfc64d51","hunter2",null],"id":1,"jsonrpc":"2.0"}' -H "Content-Type: application/json" -X POST http://localhost:8545