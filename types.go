package totp

// Algorithm defines the Algorithm used for TOTP generation
type Algorithm int

// Encoding defines the encoder used to convert to string
type Encoding int

const (
	// AlgorithmSHA1 is utility to refer to "crypto/sha1"
	AlgorithmSHA1   Algorithm = iota
	// AlgorithmSHA256 is utility to refer to "crypto/sha256"
	AlgorithmSHA256
	// AlgorithmSHA512 is utility to refer to "crypto/sha512"
	AlgorithmSHA512
	// AlgorithmMD5 is utility to refer to "crypto/md5"
	AlgorithmMD5
)

const (
	// EncodingBase32 is utility to refer to "encoding/base32"
	EncodingBase32 Encoding = iota
	// EncodingBase64 is utility to refer to "encoding/base64"
	EncodingBase64
)

// AuthOpts is used to define the arguments to be passed to the function generating the TOTP
type AuthOpts struct {
	Algo   Algorithm `json:"algorithm"`
	Encode Encoding  `json:"encoding"`
	Digits int       `json:"digits"`
	Period int       `json:"period"`
	Secret string    `json:"secret"`
}