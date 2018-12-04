package utils

import (
	"fmt"
	"github.com/ellsol/go-ethertypes"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"golang.org/x/net/context"
	"time"
)

func ExecuteTransaction(label string, client *ethclient.Client, tx *types.Transaction) (error) {
	var err error
	fmt.Println(ConsoleInRed("Executing Transaction"), label, tx.Hash().String())
	PrintColoredln("GasPrice", ethertypes.NewEtherValue().SetBigInt(tx.GasPrice()).String())
	PrintColoredln("Gas", tx.Gas())

	for pending := true; pending; _, pending, err = client.TransactionByHash(context.Background(), tx.Hash()) {
		if err != nil {
			return err
		}
		fmt.Print(ConsoleInBlue("."))
		time.Sleep(2 * time.Second)
	}

	fmt.Println(ConsoleInGreen("Transaction Confirmed!!!"))
	PrintColoredln(" Address", tx.Hash().String())

	return nil
}
