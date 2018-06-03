package ncoin_wallet

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"encoding/json"
	"fmt"
	internal "github.com/NeironTeam/ncoin/internal"
	"encoding/hex"
)

//Transaction struct contains all the attributes necessary for one transaction of the blockchain.
type Transaction struct {
	addressTo   uint64  `json:"addressTo"`
	addressFrom uint64  `json:"addressFrom"`
	quantity    float64 `json:"quantity"`
	sign        []byte  `json:"quantity"`
	fee         float64 `json:"Fee"`
}

//NewTransaction its a constructor for a new Transaction struct, receive all the attribute of the Transaction except the sign
func NewTransaction(addressTo uint64, addressFrom uint64, quantity float64, fee float64) Transaction {
	return Transaction{addressFrom: addressFrom, addressTo: addressTo, quantity: quantity, fee: fee}
}

//Fee its the getter for the Fee attribute
func (t *Transaction) Fee() float64 {
	return t.fee
}

//GetSign its the getter for the sign attribute
func (t *Transaction) GetSign() []byte {
	return t.sign
}

//Quantity its the getter for the quantity attribute
func (t *Transaction) Quantity() float64 {
	return t.quantity
}

//AddressFrom its the getter for the addressFrom attribute
func (t *Transaction) AddressFrom() uint64 {
	return t.addressFrom
}

//AddressTo its the getter for the addressTo attribute
func (t *Transaction) AddressTo() uint64 {
	return t.addressTo
}

//Sign sign the transaction with the privateKey *rsa.PrivateKey and set the attribute sign. Return a err if somewhat goes wrong or nil if all it`s ok.

func (t *Transaction) Sign(privateKey *rsa.PrivateKey) (e error) {
	var hash []byte = internal.CalculateGenericHash(t.toString())
	t.sign, e = rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hash)
	return
}

//Verify verify the sign of the transaction with the parameter publicKey *rsa.PublicKey. Return a err if somewhat goes wrong or nil if all it`s ok.
func (t *Transaction) Verify(publicKey *rsa.PublicKey) error {
	hash := internal.CalculateGenericHash(t.toString())
	err := rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, hash, t.GetSign())

	return err
}

func (t *Transaction) toString() (s string) {
	s = fmt.Sprintf("%d%d%f%f", t.addressTo, t.addressFrom, t.quantity, t.fee)
	return
}

//Stringify return the Transaction in string format
func (t *Transaction) Stringify() (s string) {
	s = fmt.Sprintf("%d%d%f%f%s", t.addressTo, t.addressFrom, t.quantity, t.fee, t.sign)
	return
}

func (t *Transaction) CalculateHash() string {
	return hex.EncodeToString(internal.CalculateGenericHash(t.Stringify()))
}

func (t *Transaction) CheckHash(inputHash string) bool {
	return t.CalculateHash() == inputHash

}

func TransactionFromJson(r []byte) (t *Transaction, e error) {
	e = json.Unmarshal(r, &t)
	return
}

func (t *Transaction) ToJson() ([]byte, error) {
	return json.Marshal(t)
}
