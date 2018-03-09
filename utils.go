package main

import (
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"crypto/md5"
	"hash"
	"encoding/json"
	"encoding/base32"
	"encoding/base64"
	"io/ioutil"
)

// PopulateAuthOpts populates the AuthOpts struct with the data from config.json
func PopulateAuthOpts() (AuthOpts, error) {
	opts := AuthOpts{}

	configFile, err := ioutil.ReadFile("config.json")
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

// EncodeToString encodes the string according to the encoding set in AuthOpts
func EncodeToString(sum []byte, opts AuthOpts) (string, error) {
	switch opts.Encode {
	case EncodingBase32:
		return base32.StdEncoding.EncodeToString(sum)[:opts.Digits], nil
	case EncodingBase64:
		return base64.StdEncoding.EncodeToString(sum)[:opts.Digits], nil
	}
	return "", ErrUndefinedEncoding
}

// StringToBytes converts a string to a slice of bytes
func StringToBytes(str string) []byte {
	return []byte(str)
}