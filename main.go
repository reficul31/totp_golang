package main

import "fmt"

func main() {
	opts, err := PopulateAuthOpts()
	if err != nil {
		panic(err)
	}

	generatedTOTP, err := GenerateTOTP(opts)
	if err != nil {
		panic(err)
	}

	fmt.Println("Generated OTP:", generatedTOTP)
}