package ncoin_wallet

import (
	"crypto/sha256"
	"crypto/rsa"
)

func CalculateGenericHash(string s) []byte{
	h := sha256.New()
	h.Write([]byte(s))
	return h.Sum(nil)
}


func GenerateKeys() *rsa.PrivateKey, *rsa.PublicKey {

}


