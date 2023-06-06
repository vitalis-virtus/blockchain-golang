package main

import (
	"fmt"
	"strconv"

	"github.com/vitalis-virtus/blockchain-golang/blockchain"
)

func main() {
	chain := blockchain.InitBlockChain()
	chain.AddBlock("First block after Genesis")
	chain.AddBlock("Second block after Genesis")
	chain.AddBlock("Third block after Genesis")

	for _, block := range chain.Chain {
		fmt.Printf("Prev hash: %x\n", block.PrevHash)
		fmt.Printf("Data in block: %s\n", block.Data)
		fmt.Printf("Block hash: %x\n", block.Hash)
		fmt.Printf("Nonce: %v\n", block.Nonce)

		pow := blockchain.NewProof(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}
}
