package main

import (
	"time"
	"strings"
	"crypto/hmac"
)

// ValidateTOTP validates the expectedTOTP with the generatedTOTP
func ValidateTOTP(expectedTOTP string, configPath string) (bool, error) {
	opts, err := PopulateAuthOpts(configPath)
	if err != nil {
		return false, err 
	}

	generatedTOTP, err := GenerateTOTP(opts)
	if err != nil {
		return false, err
	}

	return strings.Compare(expectedTOTP, generatedTOTP) == 0, nil
}

// GenerateTOTP generates a TOTP based upon the AuthOpts provided
func GenerateTOTP(opts AuthOpts) (string, error) {
	counter := time.Now().Second() / opts.Period
	
	algo, err := ReturnHash(opts.Algo)
	if err != nil {
		return "", err
	}

	mac := hmac.New(algo, StringToBytes(opts.Secret))
	mac.Write(StringToBytes(string(counter)))

	return EncodeToString(mac.Sum(nil), opts)
}