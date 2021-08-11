package blockchain

import (
	"crypto/sha256"
	"fmt"
	"sync"
)

type block struct {
	data     string
	hash     string
	prevHash string
}

type blockchain struct {
	blocks []*block
}

var b *blockchain
var once sync.Once

func getLastHash() string {
	blocks := GetBlockchain().blocks
	if len(blocks) == 0 {
		return ""
	}

	return blocks[len(blocks)-1].hash
}

func (b *block) generateHash() {
	hash := sha256.Sum256([]byte(b.data + b.prevHash))
	b.hash = fmt.Sprintf("%x", hash)
}

func createBlock(data string) *block {
	newBlock := block{data, "", getLastHash()}
	newBlock.generateHash()
	return &newBlock
}

func GetBlockchain() *blockchain {
	if b == nil {
		once.Do(func() {
			b = &blockchain{}
			b.blocks = append(b.blocks, createBlock("Genesis Block"))
		})
	}
	return b
}
