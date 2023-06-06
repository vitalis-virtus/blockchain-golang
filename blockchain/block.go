package blockchain

import (
	"bytes"
	"crypto/sha256"
)

type BlockChain struct {
	Chain []*Block
}

type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
}

// DeriveHash derives hash from block data and prev block hash and then write derived hash in current block.Hash
func (b *Block) DeriveHash() {
	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	hash := sha256.Sum256(info)
	b.Hash = hash[:]
}

// CreateBlock creates and return a new block
func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash}
	block.DeriveHash()
	return block
}

// AddBlock adds new block with new data to the current blockchain
func (chain *BlockChain) AddBlock(data string) {
	prevBlock := chain.Chain[len(chain.Chain)-1]
	newBlock := CreateBlock(data, prevBlock.Hash)
	chain.Chain = append(chain.Chain, newBlock)
}

// Genesis creates first genesis block in blockchain
func Genesis() *Block {
	return CreateBlock("Genesis block", []byte{})
}

// InitBlockChain creates new BlockChain with initiated Genesis block
func InitBlockChain() *BlockChain {
	newChain := BlockChain{[]*Block{Genesis()}}
	return &newChain
}
