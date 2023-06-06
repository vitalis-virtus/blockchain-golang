package blockchain

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math"
	"math/big"
	"os"
	"strconv"

	"github.com/vitalis-virtus/blockchain-golang/utils"
)

var Difficulty, _ = strconv.Atoi(os.Getenv("DIFFICULTY"))

// Requirements: the first bytes must contain 0s which si derived from difficulty

type ProofOfWork struct {
	Block *Block
	// Target is the number that represents the requirement difficulty.It contain 0s at start
	Target *big.Int
}

// NewProof create new ProofOfWork
func NewProof(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	// left shift performs a bitwise shift to the left
	target.Lsh(target, uint(256-Difficulty))

	pow := &ProofOfWork{b, target}

	return pow
}

// InitData take the data from the block
func (pow *ProofOfWork) InitData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.Block.PrevHash,
			pow.Block.Data,
			utils.ConvertToHex(int64(nonce)),
			utils.ConvertToHex(int64(Difficulty)),
		},
		[]byte{})

	return data
}

// Run minig new block - it searching correct nonce and represents block hash
func (pow *ProofOfWork) Run() (int, []byte) {
	var intHash big.Int
	var hash [32]byte

	// create a counter (nonce) which starts at 0
	nonce := 0

	for nonce < math.MaxInt64 {
		data := pow.InitData(nonce)
		hash = sha256.Sum256(data)

		fmt.Printf("\r%x", hash)
		intHash.SetBytes(hash[:])

		// we compare two big.Int numbers and check for corectness of the nonce
		if intHash.Cmp(pow.Target) == -1 {
			break
		} else {
			nonce++
		}
	}
	fmt.Println()

	return nonce, hash[:]
}

// Validate check the correctness of nonce
func (pow *ProofOfWork) Validate() bool {
	var intHash big.Int

	data := pow.InitData(pow.Block.Nonce)

	hash := sha256.Sum256(data)
	intHash.SetBytes(hash[:])

	// we compare two big.Int numbers and check for corectness of the nonce
	return intHash.Cmp(pow.Target) == -1
}
