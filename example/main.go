package main

import (
	"fmt"
	totp "github.com/reficul31/totp_golang"
)

func main() {
	authOpts, err := totp.PopulateAuthOpts("config.json")
	if err != nil {
		panic(err)
	}

	otp, err := totp.GenerateTOTP(authOpts)
	if err != nil {
		panic(err)
	}

	truth, err := totp.ValidateTOTP(otp, authOpts)
	if err != nil {
		panic(err)
	}

	fmt.Println("OTP:", otp, " Validation Result:", truth)
}