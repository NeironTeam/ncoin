package ncoin_wallet

import (
    "crypto/rsa"
    "crypto/sha256"
    "crypto"
    "strconv"
    "crypto/rand"
)
type Transact interface {
    NewTransaction(addressTo uint64, addressFrom uint64, quantity float64) Transaction
    GetSign() []byte
    Quantity() float64
    AddressFrom() uint64
    AddressTo() uint64
    SignTransaction(privateKey *rsa.PrivateKey) error
    Stringify() string
}

type Transaction struct {
    addressTo   uint64
    addressFrom uint64
    quantity    float64
    sign        []byte
    fee         float64
}

func (t *Transaction) NewTransaction(addressTo uint64, addressFrom uint64, quantity float64, fee float64) Transaction {
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


func (t *Transaction) SignTransaction(privateKey *rsa.PrivateKey) error{
    sha_256 := sha256.New()
    sha_256.Write([]byte(t.toString()))

    hash := sha_256.Sum(nil)
    sign, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hash)
    if err != nil{
        return err
    }
    t.sign = sign
    return nil
}

func (t *Transaction) toString() string{
    s := ""
    s += (string) (t.addressTo)
    s += (string) (t.addressFrom)
    s += strconv.FormatFloat(t.quantity, 'f', 5, 64)
    s += strconv.FormatFloat(t.fee, 'f', 5, 64)
    return s
}

func (t *Transaction) Stringify() string{
    s := ""

    return s
}