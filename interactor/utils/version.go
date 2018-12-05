package utils

import "io/ioutil"

const versionFilename = "current_version"

func ReadVersion() (string, error) {
	i, err := ioutil.ReadFile(versionFilename)

	if err != nil {
		return "", err
	}

	return string(i), nil
}

func WriteVersion(version string) error {

	return ioutil.WriteFile(versionFilename, []byte(version), 0644)
}
