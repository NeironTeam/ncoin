package wallet

// Blockchain struct contains all blockchain blocks and reference to last block
type Blockchain struct {
	blocks []block
	lastBlock block
}

// AddBlock function appends new block to the blockchain and reference that 
// block as lastBlock.
func (b *Blockchain) AddBlock(n block) {}

// GetMerkleTreeRoot function iterates over whole blockchain blocks and 
// generates its hash. Then combines that hash and returns the root of blcks 
// hash Merkle Tree.
func (b *blockchain) GetMerkleTreeRoot() (hash string) { return }
