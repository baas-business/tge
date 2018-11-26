package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/baas/tge-sol/deployer/contracts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/urfave/cli"
	"io/ioutil"
	"log"
	"math/big"
	"os"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/x-cray/logrus-prefixed-formatter"
	"golang.org/x/net/context"
)

func main() {
	var keystoreUTCPath string
	var ipcPath string
	var passwordFile string
	var httpPath string
	var privKeyString string

	customFormatter := new(prefixed.TextFormatter)
	customFormatter.TimestampFormat = "2006-01-02 15:04:05"
	customFormatter.FullTimestamp = true
	logrus.SetFormatter(customFormatter)
	log := logrus.WithField("prefix", "main")

	app := cli.NewApp()
	app.Name = "deployVRC"
	app.Usage = "this is a util to deploy validator registration contract"
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

		deployTGEContracts(txOps, client)
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func deployTGEContracts(txOps *bind.TransactOpts, client *ethclient.Client) {
	var err error

	cc := &ContractConfig{
		Deployer: &txOps.From,
	}

	cc.TokenAddress, err = deployBaasToken(txOps, client, cc)
	if err != nil {
		log.Fatal(err)
	}

	cc.FounderAddress, err = deployBaasFounder(txOps, client, cc)
	if err != nil {
		log.Fatal(err)
	}

	cc.IncentivesAddress, err = deployBaasIncentives(txOps, client, cc)
	if err != nil {
		log.Fatal(err)
	}

	cc.PPAddress, err = deployBaasPP(txOps, client, cc)
	if err != nil {
		log.Fatal(err)
	}

	cc.EscrowAddress, err = deployBaasEscrow(txOps, client, cc)
	if err != nil {
		log.Fatal(err)
	}

	cc.ROIAddress, err = deployBaasROI(txOps, client, cc)
	if err != nil {
		log.Fatal(err)
	}

	err = cc.write("contract_config_v0.json")

	if err != nil {
		log.Fatal(err)
	}
}

func deployBaasToken(txOps *bind.TransactOpts, client *ethclient.Client, cc *ContractConfig) (*common.Address, error) {
	addr, tx, _, err := contracts.DeployBaasToken(txOps, client)

	if err != nil {
		return nil, err
	}

	return deploy("Token", txOps, client, tx, &addr)
}

func deployBaasFounder(txOps *bind.TransactOpts, client *ethclient.Client, cc *ContractConfig) (*common.Address, error) {
	addr, tx, _, err := contracts.DeployBaasFounder(txOps, client, *cc.TokenAddress)

	if err != nil {
		return nil, err
	}

	return deploy("Founder", txOps, client, tx, &addr)
}

func deployBaasEscrow(txOps *bind.TransactOpts, client *ethclient.Client, cc *ContractConfig) (*common.Address, error) {
	addr, tx, _, err := contracts.DeployBaasIncentive(txOps, client, *cc.TokenAddress)

	if err != nil {
		return nil, err
	}

	return deploy("Escrow", txOps, client, tx, &addr)
}

func deployBaasIncentives(txOps *bind.TransactOpts, client *ethclient.Client, cc *ContractConfig) (*common.Address, error) {
	addr, tx, _, err := contracts.DeployBaasIncentive(txOps, client, *cc.TokenAddress)

	if err != nil {
		return nil, err
	}

	return deploy("Incentives", txOps, client, tx, &addr)
}

func deployBaasPP(txOps *bind.TransactOpts, client *ethclient.Client, cc *ContractConfig) (*common.Address, error) {
	addr, tx, _, err := contracts.DeployBaasPP(txOps, client, *cc.TokenAddress)

	if err != nil {
		return nil, err
	}

	return deploy("Private Placement", txOps, client, tx, &addr)
}

func deployBaasROI(txOps *bind.TransactOpts, client *ethclient.Client, cc *ContractConfig) (*common.Address, error) {
	addr, tx, _, err := contracts.DeployBaasROI(txOps, client, *cc.TokenAddress)

	if err != nil {
		return nil, err
	}

	return deploy("ROI", txOps, client, tx, &addr)
}

func deploy(label string, txOps *bind.TransactOpts, client *ethclient.Client, tx *types.Transaction, addr *common.Address) (*common.Address, error) {
	var err error
	log.Println("Deploying", label, addr.String(), tx.Hash().String())
	for pending := true; pending; _, pending, err = client.TransactionByHash(context.Background(), tx.Hash()) {
		if err != nil {
			return nil, err
		}
		time.Sleep(2 * time.Second)
	}

	log.Println(label, "Deployed, Address: ", addr.String())

	return addr, nil
}

type ContractConfig struct {
	Deployer          *common.Address `json:"deployer"`
	TokenAddress      *common.Address `json:"token_address"`
	EscrowAddress     *common.Address `json:"escrow_address"`
	FounderAddress    *common.Address `json:"founder_address"`
	IncentivesAddress *common.Address `json:"incentives_address"`
	PPAddress         *common.Address `json:"pp_address"`
	ROIAddress        *common.Address `json:"roi_address"`
}

func (cc *ContractConfig) write(filename string) error {
	j, err := json.Marshal(cc)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filename, j, 0644)
	fmt.Printf("%+v", cc)
	return nil
}
