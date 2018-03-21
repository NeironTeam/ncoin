package ncoin_wallet

import (
	"strconv"
	"crypto/sha256"
	"encoding/hex"
)

// private Block struct

type block struct {
	timestamp string
	transactions [] Transaction
	prevHash string
	fee float64
	merkelRoot string
}

func (b *block) Stingfy() string {
	var tr string = ""
	for i := 1; i < len(b.transactions); i++ {
		tr = tr + b.transactions[i].Stingfy()
	}
	return b.timestamp + tr + b.prevHash + strconv.FormatFloat(b.fee, 'f',10,64) + b.merkelRoot
}

func (b *block) CalculateHash() string {
	h := sha256.New()
	h.Write([]byte(b.Stingfy()))
	return hex.EncodeToString(h.Sum(nil))
}

func (b *block) CheckHash( inputHash string) bool {
	return b.calculateHash() == inputHash

}
