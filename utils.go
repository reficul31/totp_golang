package totp

import (
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"crypto/md5"
	"hash"
	"encoding/json"
	"io/ioutil"
)

// PopulateAuthOpts populates the AuthOpts struct with the data from config.json
func PopulateAuthOpts(configPath string) (AuthOpts, error) {
	opts := AuthOpts{}

	configFile, err := ioutil.ReadFile(configPath)
	if err != nil {
		return opts, err
	}

	if err = json.Unmarshal(configFile, &opts); err != nil {
		return opts, err
	}

	return opts, nil
}

// ReturnHash returns a hashing function based upon the Algorithm specified
func ReturnHash(algo Algorithm) (func() hash.Hash, error) {
	switch algo {
	case AlgorithmSHA1:
		return sha1.New, nil
	case AlgorithmSHA256:
		return sha256.New, nil
	case AlgorithmSHA512:
		return sha512.New, nil
	case AlgorithmMD5:
		return md5.New, nil
	}
	return nil, ErrUndefinedAlgorithm
}

// StringToBytes converts a string to a slice of bytes
func StringToBytes(str string) []byte {
	return []byte(str)
}