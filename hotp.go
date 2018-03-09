package totp

import (
	"math"
	"strings"
	"crypto/hmac"
	"encoding/binary"
	"fmt"
)

func ValidateHOTP(expectedHOTP string, counter uint64, configPath string) (bool, error) {
	opts, err := PopulateAuthOpts(configPath)
	if err != nil {
		return false, err 
	}

	generatedHOTP, err := GenerateHOTP(opts, counter)
	if err != nil {
		return false, err
	}

	return strings.Compare(expectedHOTP, generatedHOTP) == 0, nil
}

func GenerateHOTP(opts AuthOpts, counter uint64) (string, error) {
	algo, err := ReturnHash(opts.Algo)
	if err != nil {
		return "", err
	}

	buf := make([]byte, 20)
	mac := hmac.New(algo, StringToBytes(opts.Secret))
	binary.BigEndian.PutUint64(buf, counter)

	mac.Write(buf)
	sum := mac.Sum(nil)

	offset := sum[len(sum)-1] & 0xf
	value := int64(((int(sum[offset]) & 0x7f) << 24) | ((int(sum[offset+1] & 0xff)) << 16) | ((int(sum[offset+2] & 0xff)) << 8) | (int(sum[offset+3]) & 0xff))

	mod := int32(value % int64(math.Pow10(opts.Digits)))

	return fmt.Sprintf(fmt.Sprintf("%%0%dd", opts.Digits), mod), nil
}