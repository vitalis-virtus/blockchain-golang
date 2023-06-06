package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

type BlockChain struct {
	chain []*Block
}

type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
}

func (b *Block) DeriveHash() {
	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	hash := sha256.Sum256(info)
	b.Hash = hash[:]
}

func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash}
	block.DeriveHash()
	return block
}

func (chain *BlockChain) AddBlock(data string) {
	prevBlock := chain.chain[len(chain.chain)-1]
	newBlock := CreateBlock(data, prevBlock.Hash)
	chain.chain = append(chain.chain, newBlock)
}

func Genesis() *Block {
	return CreateBlock("Genesis block", []byte{})
}

func InitBlockChain() *BlockChain {
	newChain := BlockChain{[]*Block{Genesis()}}
	return &newChain
}

func main() {
	chain := InitBlockChain()
	chain.AddBlock("First block after Genesis")
	chain.AddBlock("Second block after Genesis")
	chain.AddBlock("Third block after Genesis")

	for _, block := range chain.chain {
		fmt.Printf("Prev hash: %x\n", block.PrevHash)
		fmt.Printf("Data in block: %s\n", block.Data)
		fmt.Printf("Block hash: %x\n", block.Hash)

	}
}
