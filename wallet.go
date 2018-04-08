// wallet.go
// Autor: NeironTeam
// Licencia: MIT License, Copyright (c) 2018 Neiron

package ncoin_wallet

import (
    "crypto/rsa"
    "crypto/sha256"
    "golang.org/x/crypto/ripemd160"
    "github.com/akamensky/base58"
)

// Cartera
type Wallet struct {
    address     string
    privateKey  *rsa.PrivateKey
    publicKey   *rsa.PublicKey
    balance     float64
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

// Based on protocol defined in https://en.bitcoin.it/wiki/Technical_background_of_version_1_Bitcoin_addresses
//  Generates a address on base58check format assuming private and public keys
// have been already declared.
func (w *Wallet) GenerateAddress() {
    bytes := []byte(w.publicKey.E)
    data := &bytes
    var checksum [4]byte
    var base58_address string


    // step 1 & 2
    ProcessSHA256(data)
    ProcessRIPEMD160(data)

    // Add version byte
    binary_address := []byte{0}
    for i := 0; i < len(bytes); i++ {
        binary_address = append(binary_address, bytes[i])
    }
    bytes = binary_address

    //step 4 & 5
    ProcessSHA256(data)
    ProcessSHA256(data)

    // get the checksum
    for i := 0; i < 4; i++ {
        checksum[i] = bytes[i]
    }

    // Final binary address
    for i := 0; i < len(checksum); i++ {
        binary_address = append(binary_address, checksum[i])
    }

    // base58check format
    base58_address = base58.Encode(binary_address)
    w.address = base58_address
}

// Envía la cantidad de NCoin indicada a la dirección indicada. Requiere
// candidad y dirección como parámetros .
func (w *Wallet) SendTransaction(amount float64, address uint64) {

}

// Devuelve la dirección de la cartera.
func (w *Wallet) Address() uint64 {
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
