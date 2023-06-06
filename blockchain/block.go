package blockchain

import (
	"bytes"
	"encoding/gob"

	"github.com/vitalis-virtus/blockchain-golang/utils"
)

type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
	Nonce    int
}

// CreateBlock creates and return a new block
func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash, 0}
	pow := NewProof(block)
	nonce, hash := pow.Run()
	block.Hash = hash
	block.Nonce = nonce

	return block
}

// Genesis creates first genesis block in blockchain
func Genesis() *Block {
	return CreateBlock("Genesis block", []byte{})
}

func (b *Block) Serialize() []byte {
	var res bytes.Buffer
	encoder := gob.NewEncoder(&res)

	err := encoder.Encode(b)

	utils.Handle(err)

	return res.Bytes()
}

func Deserialize(b []byte) *Block {
	var block Block
	decoder := gob.NewDecoder(bytes.NewReader(b))

	err := decoder.Decode(&block)

	utils.Handle(err)

	return &block
}
