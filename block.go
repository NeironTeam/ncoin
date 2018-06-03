package ncoin_wallet

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	internal "github.com/NeironTeam/ncoin/internal"
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

func (b *block) CalculateMerkleTree() (hashList []string, merkleRoot string){
	for _, transaction := range b.Transactions {
		hashList = append(hashList, transaction.CalculateHash())
	}
	merkleRoot = CalculateMerkleRoot(hashList)
	return
}

func CalculateMerkleRoot(hashList []string) (root string) {
	if len(hashList) == 1 {
		return hashList[0]
	}
	var newLevel []string = make([]string, 0)
	var pos int = 0;
	for pos < len(hashList) {
		if pos+1 == len(hashList){
			newLevel = append(newLevel, hex.EncodeToString(internal.CalculateGenericHash(hashList[pos]+hashList[pos+1])))
		} else {
			newLevel = append(newLevel, hex.EncodeToString(internal.CalculateGenericHash(hashList[pos]+hashList[pos])))
		}
	}
	CalculateMerkleRoot(newLevel)
	return
}