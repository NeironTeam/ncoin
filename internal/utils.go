package internal

import (
	"fmt"
	"crypto/sha256"
	"golang.org/x/crypto/ripemd160"
	"os"
)

const baseUri = "http://%s:%s"
const DEFAULT_WALLET_PORT = "11811"
const DEFAULT_WALLET_HOST = "localhost"

func GetHost() string {
	var walletHost string = Getenv("WALLET_HOST", DEFAULT_WALLET_HOST)
	var walletPort string = Getenv("WALLET_PORT", DEFAULT_WALLET_PORT)

	return fmt.Sprintf("%s:%s", walletHost, walletPort)
}

func GetHostUri() string {
	var walletHost string = Getenv("WALLET_HOST", DEFAULT_WALLET_HOST)
	var walletPort string = Getenv("WALLET_PORT", DEFAULT_WALLET_PORT)

	return fmt.Sprintf(baseUri, walletHost, walletPort)
}

func ComposeHostUri(path string) string {
	return fmt.Sprintf("%s%s", GetHostUri(), path)
}

//CalculateGenericHash return the SHA256 hash of String s.
func CalculateGenericHash(s string) []byte {
	h := sha256.New()
	h.Write([]byte(s))
	return h.Sum(nil)
}

func Getenv(key string, def string) (value string) {
	if value = os.Getenv(key); value == "" {
		value = def
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
