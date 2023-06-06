package blockchain

import (
	"fmt"
	"os"

	"github.com/dgraph-io/badger"
	"github.com/vitalis-virtus/blockchain-golang/utils"
)

type BlockChain struct {
	LastHash []byte
	Database *badger.DB
}

type BlockChainIterator struct {
	CurrentHash []byte
	Database    *badger.DB
}

// InitBlockChain creates new BlockChain with initiated Genesis block
func InitBlockChain() *BlockChain {
	var lastHash []byte

	opts := badger.DefaultOptions(os.Getenv("DB_PATH"))
	opts.Logger = nil
	opts.Dir = os.Getenv("DB_PATH")
	opts.ValueDir = os.Getenv("DB_PATH")

	db, err := badger.Open(opts)
	utils.Handle(err)

	err = db.Update(func(txn *badger.Txn) error {
		if _, err := txn.Get([]byte("lh")); err == badger.ErrKeyNotFound {
			fmt.Println("No existing blockchain found")
			genesis := Genesis()
			fmt.Println("Genesis proved")
			err = txn.Set(genesis.Hash, genesis.Serialize())
			utils.Handle(err)
			err = txn.Set([]byte("lh"), genesis.Hash)

			lastHash = genesis.Hash

			return err
		} else {
			item, err := txn.Get([]byte("lh"))
			utils.Handle(err)
			err = item.Value(func(val []byte) error {
				lastHash = val
				return nil
			})
			return err
		}
	})

	utils.Handle(err)

	blockChain := &BlockChain{LastHash: lastHash, Database: db}

	return blockChain
}

// AddBlock adds new block with new data to the current blockchain
func (chain *BlockChain) AddBlock(data string) {
	// prevBlock := chain.Chain[len(chain.Chain)-1]
	// newBlock := CreateBlock(data, prevBlock.Hash)
	// chain.Chain = append(chain.Chain, newBlock)
	var lastHash []byte
	err := chain.Database.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte("lh"))
		utils.Handle(err)
		err = item.Value(func(val []byte) error {
			lastHash = val
			return nil
		})
		return err
	})
	utils.Handle(err)

	newBlock := CreateBlock(data, lastHash)

	err = chain.Database.Update(func(txn *badger.Txn) error {
		err := txn.Set(newBlock.Hash, newBlock.Serialize())
		utils.Handle(err)
		err = txn.Set([]byte("lh"), newBlock.Hash)

		chain.LastHash = newBlock.Hash

		return err
	})
	utils.Handle(err)
}

func (chain *BlockChain) Iterator() *BlockChainIterator {
	iter := BlockChainIterator{chain.LastHash, chain.Database}

	return &iter
}

func (iter *BlockChainIterator) Next() *Block {
	var oldBlock *Block
	var encodedBlock []byte

	err := iter.Database.View(func(txn *badger.Txn) error {
		item, err := txn.Get(iter.CurrentHash)
		utils.Handle(err)
		err = item.Value(func(val []byte) error {
			encodedBlock = val
			return nil
		})
		oldBlock = Deserialize(encodedBlock)

		return err
	})
	utils.Handle(err)

	iter.CurrentHash = oldBlock.PrevHash

	return oldBlock
}
