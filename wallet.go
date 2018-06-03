// wallet.go
// Autor: NeironTeam
// Licencia: MIT License, Copyright (c) 2018 Neiron

package ncoin_wallet

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	internal "github.com/NeironTeam/ncoin/internal"
	"os"
	"runtime"
)

const WALLET_FOLDER = ".ncoin"
const SEED_SIZE = 512 //bytes

func getWalletFolder() string {
	var base string = "HOME"
	if runtime.GOOS == "windows" {
		base = "USERPROFILE"
	} else if runtime.GOOS == "plan9" {
		base = "home"
	}

	var path string = internal.Getenv("WALLET_FOLDER", WALLET_FOLDER)
	return fmt.Sprintf("%s/%s", os.Getenv(base), path)
}

// Cartera
type Wallet struct {
	address    string
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
	balance    float64
}

func (w *Wallet) SetBalance(balance float64) {
	w.balance = balance
}

func (w *Wallet) SetPublicKey(publicKey *rsa.PublicKey) {
	w.publicKey = publicKey
}

func (w *Wallet) SetPrivateKey(privateKey *rsa.PrivateKey) {
	w.privateKey = privateKey
}

func (w *Wallet) SetAddress(address string) {
	w.address = address
}

func (w *Wallet) PublicKey() *rsa.PublicKey {
	return w.publicKey
}

// Desbloquea la wallet, requiere una private_key como parametro.
func (w *Wallet) Unlock(pk uint64) {

}

// Envía la cantidad de NCoin indicada a la dirección indicada. Requiere
// candidad y dirección como parámetros .
func (w *Wallet) SendTransaction(amount float64, address uint64) {

}

// Devuelve la dirección de la cartera.
func (w *Wallet) Address() string {
	return w.address
}

// Devuelve el saldo de la cartera.
func (w *Wallet) Balance() float64 {
	return w.balance
}

// Devuelve la private_key de la cartera.
func (w *Wallet) PrivateKey() *rsa.PrivateKey {
	return w.privateKey
}

//GenerateKeys generate the rsa private and public keys, return a err if somewhat goes wrong.
func (w *Wallet) generateKeys() (e error) {
	if w.privateKey, e = rsa.GenerateKey(rand.Reader, 2048); e != nil {
		return
	}
	w.publicKey = &w.privateKey.PublicKey
	return
}

// generateAddress function return random seed hashed with sha256 to use as
// wallet address
func (w *Wallet) generateAddress() (e error) {
	var randAddress []byte = make([]byte, 512)
	if _, e = rand.Read(randAddress); e != nil {
		return
	}

	w.address = fmt.Sprintf("%x", sha256.Sum256(randAddress))
	return
}

func (w *Wallet) storeWallet() (e error) {
	// Check WALLET_FOLDER from enviroment
	var walletsPath string = getWalletFolder()

	// Check if walletsPath exits, else create it.
	if _, e = os.Stat(walletsPath); os.IsNotExist(e) {
		if e = os.Mkdir(walletsPath, os.ModePerm); e != nil {
			return
		}
	} else if e != nil {
		return
	}

	// Create walletFolder to store key pair if not exists
	var walletFolder string = fmt.Sprintf("%s/%s", walletsPath, w.address)
	if e = os.Mkdir(walletFolder, os.ModePerm); e != nil {
		return
	}

	// Generate public key pem and store it
	var pubPem []byte
	if pubPem, e = x509.MarshalPKIXPublicKey(w.publicKey); e != nil {
		return
	} else {
		if e = w.storePem(pubPem, walletFolder, true); e != nil {
			return
		}
	}

	// Generate private key pem and store it
	var privPem []byte = x509.MarshalPKCS1PrivateKey(w.privateKey)
	e = w.storePem(privPem, walletFolder, false)
	return
}

func (w *Wallet) storePem(key []byte, folder string, public bool) (e error) {
	var bType, file string
	if bType, file = "RSA PRIVATE KEY", "id_rsa"; public {
		bType = "RSA PUBLIC KEY"
		file = "id_rsa.pub"
	}

	var pemRaw []byte = pem.EncodeToMemory(
		&pem.Block{
			Type:  bType,
			Bytes: key,
		},
	)
	
	var fd *os.File
	var pemPath string = fmt.Sprintf("%s/%s", folder, file)
	if fd, e = os.Create(pemPath); e != nil {
		return
	}

	if _, e = fd.Write(pemRaw); e != nil {
		return e
	}
	fd.Close()
	return
}

func NewWallet() (w *Wallet, e error) {
	w = &Wallet{}
	if e = w.generateKeys(); e != nil {
		return
	}

	w.generateAddress()
	e = w.storeWallet()
	return
}
