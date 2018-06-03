package ncoin_wallet

import (
	"encoding/hex"
	"fmt"
	internal "github.com/NeironTeam/ncoin-wallet/internal"
)

// private Block struct

type block struct {
	timestamp string
	transactions []Transaction
	prevHash string
	fee float64
	merkelRoot string
}

func (b *block) Stringify()(s string){
	for _, transaction := range b.transactions {
		s = fmt.Sprintf("%s%s", s, transaction.Stringify())
	}
	s = fmt.Sprintf("%s%s%s%f%s", b.timestamp, s, b.prevHash, b.fee, b.merkelRoot)
	return
}

func (b *block) CalculateHash() string {
	return hex.EncodeToString(internal.CalculateGenericHash(b.Stringify()))
}

func (b *block) CheckHash( inputHash string) bool {
	return b.CalculateHash() == inputHash

}
