package ncoin_wallet

import (
	"fmt"
	"github.com/NeironTeam/ncoin/internal"
	"github.com/go-redis/redis"
	"sort"
)

const REDIS_PATH = "localhost"
const REDIS_PORT = "6379"

// Blockchain struct contains all blockchain blocks and reference to last block
type Blockchain struct {
	blocks    blockList
	lastBlock block
}

type blockList []block

func (bs blockList) Len() int           { return len(bs) }
func (bs blockList) Swap(i, j int)      { bs[i], bs[j] = bs[j], bs[i] }
func (bs blockList) Less(i, j int) bool { return bs[i].Timestamp < bs[j].Timestamp }

// AddBlock function appends new block to the blockchain and reference that
// block as lastBlock.
func (b *Blockchain) AddBlock(n block) {
	b.blocks = append(b.blocks, n)
	b.lastBlock = n
}

func (b *Blockchain) Store() {
	var (
		redisPath string = internal.Getenv("REDIS_PATH", REDIS_PATH)
		redisPort string = internal.Getenv("REDIS_PORT", REDIS_PORT)
	)

	client := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", redisPath, redisPort),
	})

	sort.Sort(b.blocks)
	for _, bl := range b.blocks {
		var hash string = bl.CalculateHash()
		if exists, err := client.HExists("blocks", hash).Result(); err != nil {
			panic(err)
		} else if exists {
			break
		}

		if jbl, err := bl.ToJson(); err != nil {
			panic(err)
		} else if err := client.HSet("blocks", hash, jbl); err != nil {
			fmt.Println(err)
		}
	}
}

func (b *Blockchain) Load() {
	var (
		redisPath string = internal.Getenv("REDIS_PATH", REDIS_PATH)
		redisPort string = internal.Getenv("REDIS_PORT", REDIS_PORT)
	)

	client := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", redisPath, redisPort),
	})

	if hashes, err := client.HKeys("blocks").Result(); err != nil {
		panic(err)
	} else {
		for _, hash := range hashes {
			var val *redis.StringCmd = client.HGet("blocks", hash)
			if res, err := val.Result(); err != nil {
				panic(err)
			} else if bl, err := BlockFromJson([]byte(res)); err != nil {
				panic(b)
			} else {
				b.blocks = append(b.blocks, bl)
			}
		}
	}
}

// GetMerkleTreeRoot function iterates over whole blockchain blocks and
// generates its hash. Then combines that hash and returns the root of blcks
// hash Merkle Tree.
func (b *Blockchain) GetMerkleTreeRoot() (hash string) { return }
