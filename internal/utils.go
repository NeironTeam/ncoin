package internal

import (
	"crypto/sha256"
	"golang.org/x/crypto/ripemd160"
	"os"
)

//CalculateGenericHash return the SHA256 hash of String s.
func CalculateGenericHash(s string) []byte {
	h := sha256.New()
	h.Write([]byte(s))
	return h.Sum(nil)
}

func GetEnv(envVarName string, defVarValue string) (env string) {
	if env = os.Getenv(envVarName); env == "" {
		env = defVarValue
	}
	return
}

// TODO (Ventura): Revisar, puede no funcionar, mirar las funciones ancestrales
func ProcessSHA256(data []byte) []byte {
	h := sha256.New()
	h.Write(data)
	return h.Sum(nil)
}

// TODO (Ventura): Revisar, puede no funcionar, mirar las funciones ancestrales
func ProcessRIPEMD160(data []byte) []byte {
	h := ripemd160.New()
	h.Write(data)
	return h.Sum(nil)

}
