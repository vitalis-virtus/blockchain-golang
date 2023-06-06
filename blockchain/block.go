package blockchain

type BlockChain struct {
	Chain []*Block
}

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
