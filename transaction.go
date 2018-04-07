package ncoin_wallet

import (
    "crypto/rsa"
    "crypto"
    "crypto/rand"
    "fmt"
)


type Transaction struct {
    addressTo   uint64
    addressFrom uint64
    quantity    float64
    sign        []byte
    fee         float64
}

func NewTransaction(addressTo uint64, addressFrom uint64, quantity float64, fee float64) Transaction {
    return Transaction{addressFrom:addressFrom,addressTo:addressTo,quantity:quantity, fee:fee}
}

func (t *Transaction) Fee() float64 {
    return t.fee
}


func (t *Transaction) GetSign() []byte {
    return t.sign
}

func (t *Transaction) Quantity() float64 {
    return t.quantity
}


func (t *Transaction) AddressFrom() uint64 {
    return t.addressFrom
}


func (t *Transaction) AddressTo() uint64 {
    return t.addressTo
}


func (t *Transaction) Sign(privateKey *rsa.PrivateKey) error{
    hash := CalculateGenericHash(t.toString())
    sign, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hash)
    if err != nil{
        return err
    }
    t.sign = sign
    return nil
}

func (t *Transaction) Verify(publicKey *rsa.PublicKey) error{
    hash := CalculateGenericHash(t.toString())
    err := rsa.VerifyPKCS1v15(publicKey,crypto.SHA256,hash, t.GetSign() );

    return err
}

func (t *Transaction) toString() (s string){
    s = fmt.Sprintf("%d%d%f%f", t.addressTo, t.addressFrom, t.quantity, t.fee)
    return
}

func (t *Transaction) Stringify() (s string){
    s = fmt.Sprintf("%d%d%f%f%s", t.addressTo, t.addressFrom, t.quantity, t.fee, t.sign)
    return
}