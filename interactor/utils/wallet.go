package utils

import (
	"bufio"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"io/ioutil"
	"os"
)

func PrivateKeyFromWalletAndPasswordFile(keystoreUTCPath string, passwordFile string) (*keystore.Key, error) {
	password, err := password(passwordFile)

	if err != nil {
		return nil, err
	}

	keyJSON, err := ioutil.ReadFile(keystoreUTCPath)

	if err != nil {
		return nil, err
	}

	return keystore.DecryptKey(keyJSON, password)
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
