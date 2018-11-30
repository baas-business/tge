package utils

import (
	"bufio"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"io/ioutil"
	"log"
	"os"
)

func PrivateKeyFromWalletAndPasswordFile(keystoreUTCPath string, passwordFile string) (*keystore.Key, error) {
	password, err := password(passwordFile)

	if err != nil {
		return nil, fmt.Errorf("couldn't read password file %v", err.Error())
	}

	log.Println("Reading wallet ", keystoreUTCPath)
	keyJSON, err := ioutil.ReadFile(keystoreUTCPath)

	if err != nil {
		return nil, fmt.Errorf("couldn't read wallet file %v %v", keystoreUTCPath, err.Error())
	}

	key, err := keystore.DecryptKey(keyJSON, password)

	log.Println("Loaded key: ", key.Address)

	return key, err
}

func password(passwordFile string) (string, error) {
	file, err := os.Open(passwordFile)
	if err != nil {
		return "", err
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	return scanner.Text(), nil
}
