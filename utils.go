package ncoin_wallet

import (
	"crypto/sha256"
	"crypto/rsa"
    "crypto/rand"
)

func CalculateGenericHash(s string) []byte{
	h := sha256.New()
	h.Write([]byte(s))
	return h.Sum(nil)
}


func GenerateKeys() (*rsa.PrivateKey, *rsa.PublicKey, error) {
	privateKey, err := rsa.GenerateKey( rand.Reader, 2048)
    if err != nil {
        return nil, nil, err
    }
    var publicKey = privateKey.PublicKey
	return privateKey, &publicKey, err
}


