package utils

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

func RPCClient(httpPath string) (*ethclient.Client, error){
	rpcClient, err := rpc.Dial(httpPath)

	if err != nil {
		return nil, err
	}

	return ethclient.NewClient(rpcClient), nil
}
