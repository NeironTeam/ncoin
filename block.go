package ncoin_wallet

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	internal "github.com/NeironTeam/ncoin-wallet/internal"
)

// private Block struct

type block struct {
	Timestamp    string        `json:"timestamp"`
	Transactions []Transaction `json:"transactions"`
	PrevHash     string        `json:"prevHash"`
	Fee          float64       `json:"fee"`
	MerkelRoot   string        `json:"merkleRoot"`
}

func (b *block) Stringify() (s string) {
	for _, transaction := range b.Transactions {
		s = fmt.Sprintf("%s%s", s, transaction.Stringify())
	}
	s = fmt.Sprintf("%s%s%s%f%s", b.Timestamp, s, b.PrevHash, b.Fee, b.MerkelRoot)
	return
}

func (b block) ToJson() ([]byte, error) {
	return json.Marshal(b)
}

func BlockFromJson(r []byte) (b block, e error) {
	e = json.Unmarshal(r, &b)
	return
}

func (b *block) CalculateHash() string {
	return hex.EncodeToString(internal.CalculateGenericHash(b.Stringify()))
}

func (b *block) CheckHash(inputHash string) bool {
	return b.CalculateHash() == inputHash

}
