package utils

import (
	"fmt"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"golang.org/x/net/context"
	"log"
	"time"
)

func ExecuteTransaction(label string, client *ethclient.Client, tx *types.Transaction) (error) {
	var err error
	log.Println("Executing Transaction ", label, tx.Hash().String())
	for pending := true; pending; _, pending, err = client.TransactionByHash(context.Background(), tx.Hash()) {
		if err != nil {
			return err
		}
		fmt.Print(".")
		time.Sleep(2 * time.Second)
	}

	log.Println(label, "Transaction Confirmed, Address: ", tx.Hash().String())

	return nil
}
