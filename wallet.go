// wallet.go
// Autor: NeironTeam
// Licencia: MIT License, Copyright (c) 2018 Neiron

package ncoin_wallet

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/binary"
	"encoding/pem"
	"fmt"
	internal "github.com/NeironTeam/ncoin-wallet/internal"
	"github.com/akamensky/base58"
	"io"
	"os"
)

const WALLET_FOLDER = ".ncoin"

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

// TODO: Import/Export wallet with public/private key pair

// Based on protocol defined in https://en.bitcoin.it/wiki/Technical_background_of_version_1_Bitcoin_addresses
//  Generates a address on base58check format assuming private and public keys
// have been already declared.
func (w *Wallet) generateAddress() {
	var data []byte
	var wr io.Writer = io.Writer(bytes.NewBuffer(data))
	binary.Write(wr, binary.LittleEndian, w.publicKey.E)

	var checksum []byte
	var base58_address string

	// step 1 & 2
	data = internal.ProcessSHA256(data)
	data = internal.ProcessRIPEMD160(data)

	// Add version byte
	var binaryAddress []byte
	for i := 0; i < len(data); i++ {
		binaryAddress = append(binaryAddress, data[i])
	}
	data = binaryAddress

	//step 4 & 5
	internal.ProcessSHA256(data)
	internal.ProcessSHA256(data)

	// get the checksum
	for i := 0; i < 4; i++ {
		checksum[i] = data[i]
	}

	// Final binary address
	for i := 0; i < len(checksum); i++ {
		binaryAddress = append(binaryAddress, checksum[i])
	}

	// base58check format
	base58_address = base58.Encode(binaryAddress)
	w.address = base58_address
}

func (w *Wallet) storeWallet() (e error) {
	// Check WALLET_FOLDER from enviroment
	var walletsPath = internal.GetEnv("WALLET_FOLDER", WALLET_FOLDER)

	// Check if walletsPath exits, else create it.
	if _, e = os.Stat(walletsPath); os.IsNotExist(e) {
		if e = os.Mkdir(walletsPath, os.ModePerm); e != nil {
			return nil
		}
	} else {
		return
	}

	// Create walletFolder to store key pair
	var walletFolder string = fmt.Sprintf("%s/%s", walletsPath, w.address)
	if e = os.Mkdir(walletFolder, os.ModePerm); e != nil {
		return
	}

	// Generate public key pem and store it
	var pubPem []byte
	if pubPem, e = x509.MarshalPKIXPublicKey(w.publicKey); e != nil {
		return
	} else if e = w.storePem(pubPem, walletFolder, true); e != nil {
		return
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
	if e = w.generateKeys(); e != nil {
		return
	}

	w.generateAddress()
	e = w.storeWallet()
	return
}
