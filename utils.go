package ncoin_wallet

import (
	"crypto/sha256"
)

//CalculateGenericHash return the SHA256 hash of String s.
func CalculateGenericHash(s string) []byte{
	h := sha256.New()
	h.Write([]byte(s))
	return h.Sum(nil)
}
