package ncoin_wallet

import (
	"crypto/sha256"
	"crypto/rsa"
    "crypto/rand"
)

<<<<<<< HEAD
=======
//CalculateGenericHash return the SHA256 hash of String s.
>>>>>>> e06f0dac2aa4dc7ec96a2e82e79c1b171ff3946c
func CalculateGenericHash(s string) []byte{
	h := sha256.New()
	h.Write([]byte(s))
	return h.Sum(nil)
}

<<<<<<< HEAD

func GenerateKeys() (*rsa.PrivateKey, *rsa.PublicKey) { return }
=======
//GenerateKeys generate the rsa private and public keys, return a err if somewhat goes wrong.
func GenerateKeys() (*rsa.PrivateKey, *rsa.PublicKey, error) {
	privateKey, err := rsa.GenerateKey( rand.Reader, 2048)
    if err != nil {
        return nil, nil, err
    }
    var publicKey = privateKey.PublicKey
	return privateKey, &publicKey, err
}


>>>>>>> e06f0dac2aa4dc7ec96a2e82e79c1b171ff3946c
