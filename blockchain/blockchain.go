package blockchain

import (
	"crypto/sha256"
	"fmt"
	"sync"
)

type Block struct {
	Data     string
	Hash     string
	PrevHash string
}

type blockchain struct {
	blocks []*Block
}

var b *blockchain
var once sync.Once

func getLastHash() string {
	blocks := GetBlockchain().blocks
	if len(blocks) == 0 {
		return ""
	}

	return blocks[len(blocks)-1].Hash
}

func createBlock(data string) *Block {
	newBlock := Block{data, "", getLastHash()}
	newBlock.generateHash()
	return &newBlock
}

func (b *Block) generateHash() {
	hash := sha256.Sum256([]byte(b.Data + b.PrevHash))
	b.Hash = fmt.Sprintf("%x", hash)
}

func (b *blockchain) AddBlock(data string) {
	b.blocks = append(b.blocks, createBlock(data))
}

func GetBlockchain() *blockchain {
	if b == nil {
		once.Do(func() {
			b = &blockchain{}
			b.AddBlock("Genesis")
		})
	}
	return b
}

func (b *blockchain) AllBlocks() []*Block {
	return b.blocks
}
