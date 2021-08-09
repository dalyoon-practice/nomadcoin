package main

import (
	"crypto/sha256"
	"fmt"
)

type block struct {
	data     string
	hash     string
	prevHash string
}

type blockchain struct {
	blocks []block
}

func generateHash(data, prevHash string) string {
	hash := sha256.Sum256([]byte(data + prevHash))
	return fmt.Sprintf("%x", hash)
}

func (b *blockchain) getLastHash() string {
	if len(b.blocks) > 0 {
		return b.blocks[len(b.blocks)-1].hash
	}

	return ""
}

func (b *blockchain) addBlocks(data string) {
	newBlock := block{data, "", b.getLastHash()}
	newBlock.hash = generateHash(newBlock.data, newBlock.prevHash)
	b.blocks = append(b.blocks, newBlock)
}

func (b *blockchain) listBlocks() {
	for _, block := range b.blocks {
		fmt.Printf("Data: %s\n", block.data)
		fmt.Printf("Hash: %s\n", block.hash)
		fmt.Printf("Prev Hash: %s\n", block.prevHash)
	}
}

func main() {
	chain := blockchain{}
	chain.addBlocks("Genesis Block")
	chain.addBlocks("Second Block")
	chain.addBlocks("Third Block")

	chain.listBlocks()
}
