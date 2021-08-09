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

func main() {
	genesisBlock := block{"Genisis Block", "", ""}
	hash := sha256.Sum256([]byte(genesisBlock.data + genesisBlock.prevHash))
	hashAsHex := fmt.Sprintf("%x", hash)
	genesisBlock.hash = hashAsHex
	fmt.Println(genesisBlock)
}
