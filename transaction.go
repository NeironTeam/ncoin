package ncoin_wallet

import (
    "crypto/rsa"
    "crypto"
    "crypto/rand"
    "fmt"
)

//Transaction struct contains all the attributes necessary for one transaction of the blockchain.
type Transaction struct {
    addressTo   uint64
    addressFrom uint64
    quantity    float64
    sign        []byte
    fee         float64
}

//NewTransaction its a constructor for a new Transaction struct, receive all the attribute of the Transaction except the sign
func NewTransaction(addressTo uint64, addressFrom uint64, quantity float64, fee float64) Transaction {
    return Transaction{addressFrom:addressFrom,addressTo:addressTo,quantity:quantity, fee:fee}
}

//Fee its the getter for the fee attribute
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
func (t *Transaction) Sign(privateKey *rsa.PrivateKey) error{
    hash := CalculateGenericHash(t.toString())
    sign, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hash)
    if err != nil{
        return err
    }
    t.sign = sign
    return nil
}

//Verify verify the sign of the transaction with the parameter publicKey *rsa.PublicKey. Return a err if somewhat goes wrong or nil if all it`s ok.
func (t *Transaction) Verify(publicKey *rsa.PublicKey) error{
    hash := CalculateGenericHash(t.toString())
    err := rsa.VerifyPKCS1v15(publicKey,crypto.SHA256,hash, t.GetSign() )

    return err
}

func (t *Transaction) toString() (s string){
    s = fmt.Sprintf("%d%d%f%f", t.addressTo, t.addressFrom, t.quantity, t.fee)
    return
}

//Stringify return the Transaction in string format
func (t *Transaction) Stringify() (s string){
    s = fmt.Sprintf("%d%d%f%f%s", t.addressTo, t.addressFrom, t.quantity, t.fee, t.sign)
    return
}