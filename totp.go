package totp

import (
	"time"
	"strings"
	"math"
)

// ValidateTOTP validates the expectedTOTP with the generatedTOTP
func ValidateTOTP(expectedTOTP string, opts AuthOpts) (bool, error) {
	generatedTOTP, err := GenerateTOTP(opts)
	if err != nil {
		return false, err
	}
	return strings.Compare(expectedTOTP, generatedTOTP) == 0, nil
}

// GenerateTOTP generates a TOTP based upon the AuthOpts provided
func GenerateTOTP(opts AuthOpts) (string, error) {
	t := time.Now()
	counter := uint64(math.Floor(float64(t.Unix()) / float64(opts.Period)))
	return GenerateHOTP(opts, counter)
}