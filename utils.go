package ncoin_wallet

import (
	"crypto/sha256"
	"golang.org/x/crypto/ripemd160"
)

//CalculateGenericHash return the SHA256 hash of String s.
func CalculateGenericHash(s string) []byte{
	h := sha256.New()
	h.Write([]byte(s))
	return h.Sum(nil)
}

// TODO (Ventura): Revisar, puede no funcionar, mirar las funciones ancestrales
func ProcessSHA256(data *[]byte){
	h := sha256.New()
	h.Write(data)
	data = h.Sum(nil)
}

// TODO (Ventura): Revisar, puede no funcionar, mirar las funciones ancestrales
func ProcessRIPEMD160(data *[]byte){
	h := ripemd160.New()
	h.Write(data)
	data = h.Sum(nil)
}
