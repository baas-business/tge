package main

import (
	"bufio"
	"fmt"
	"github.com/baas/tge-sol/deployer/contracts"
	"github.com/baas/tge-sol/deployer/deployer"
	"github.com/baas/tge-sol/deployer/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/urfave/cli"
	"io/ioutil"
	"log"
	"math/big"
	"os"
)

func main() {
	var keystoreUTCPath string
	var ipcPath string
	var passwordFile string
	var httpPath string
	var privKeyString string

	app := cli.NewApp()
	app.Name = "deployVRC"
	app.Usage = "this is a util to deploy tge contracts"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "keystoreUTCPath",
			Value:       "./wallets/wallet_android.json",
			Usage:       "Location of keystore",
			Destination: &keystoreUTCPath,
		},
		cli.StringFlag{
			Name:        "ipcPath",
			Usage:       "Filename for IPC socket/pipe within the datadir",
			Destination: &ipcPath,
		},
		cli.StringFlag{
			Name:        "httpPath",
			Value:       "https://ropsten.infura.io/qyXg12v7VcFCETVWfUaJ/",
			Usage:       "HTTP-RPC server listening interface",
			Destination: &httpPath,
		},
		cli.StringFlag{
			Name:        "passwordFile",
			Value:       "./wallets/password.txt",
			Usage:       "Password file for unlock account",
			Destination: &passwordFile,
		},
		cli.StringFlag{
			Name:        "privKey",
			Usage:       "Private key to unlock account",
			Destination: &privKeyString,
		},
	}

	app.Action = func(c *cli.Context) {
		// Set up RPC client
		var rpcClient *rpc.Client
		var err error
		var txOps *bind.TransactOpts

		// Uses HTTP-RPC if IPC is not set
		if ipcPath == "" {
			rpcClient, err = rpc.Dial(httpPath)
		} else {
			rpcClient, err = rpc.Dial(ipcPath)
		}
		if err != nil {
			log.Fatal(err)
		}

		client := ethclient.NewClient(rpcClient)

		// User inputs private key, sign tx with private key
		if privKeyString != "" {
			privKey, err := crypto.HexToECDSA(privKeyString)
			if err != nil {
				log.Fatal(err)
			}
			txOps = bind.NewKeyedTransactor(privKey)
			txOps.Value = big.NewInt(0)

			// User inputs keystore json file, sign tx with keystore json
		} else {
			// #nosec - Inclusion of file via variable is OK for this tool.
			file, err := os.Open(passwordFile)
			if err != nil {
				log.Fatal(err)
			}

			scanner := bufio.NewScanner(file)
			scanner.Split(bufio.ScanWords)
			scanner.Scan()
			password := scanner.Text()

			// #nosec - Inclusion of file via variable is OK for this tool.
			keyJSON, err := ioutil.ReadFile(keystoreUTCPath)
			if err != nil {
				log.Fatal(err)
			}
			privKey, err := keystore.DecryptKey(keyJSON, password)
			if err != nil {
				log.Fatal(err)
			}

			txOps = bind.NewKeyedTransactor(privKey.PrivateKey)
			txOps.Value = big.NewInt(0)
		}

		//info.ShowInfo(client, "contract_config_v2.json")
		SetupContracts(txOps, client)

		//DeployContracts(txOps, client)
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func DeployContracts(txOps *bind.TransactOpts, client *ethclient.Client) {
	cc, err := deployer.DeployTGEContracts(txOps, client)

	if err != nil {
		log.Fatal(err)
	}

	err = cc.Write("contract_config_v2.json")

	if err != nil {
		log.Fatal(err)
	}
}
func SetupContracts(txOps *bind.TransactOpts, client *ethclient.Client) {

	cc, err := deployer.LoadContractConfig("contract_config_v2.json")

	contract, err := contracts.NewBaasToken(*cc.TokenAddress, client)

	if err != nil {
		log.Fatal(err)
	}

	ii, err := contract.IsInitialized(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("initialized: ", ii)

	//Setup(opts *bind.TransactOpts, escrowAddress common.Address, ppAddress common.Address, founderAddress common.Address, incentivesAddress common.Address) (*types.Transaction, error) {
	tx, err := contract.Setup(txOps, *cc.EscrowAddress, *cc.PPAddress, *cc.FounderAddress, *cc.IncentivesAddress)
	err = utils.ExecuteTransaction("Setup Token", client, tx)

	if err != nil {
		log.Fatal(err)
	}
}
