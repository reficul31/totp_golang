package main

import (
	"fmt"
	"time"
	"math"
)

func main() {
	opts, err := PopulateAuthOpts("config.json")
	if err != nil {
		panic(err)
	}

	t := time.Now()
	counter := uint64(math.Floor(float64(t.Unix()) / float64(opts.Period)))

	generatedHOTP, err := GenerateHOTP(opts, counter)
	if err != nil {
		panic(err)
	}

	fmt.Println("Generated OTP:", generatedHOTP)
}