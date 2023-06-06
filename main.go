package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/vitalis-virtus/blockchain-golang/blockchain"
	"github.com/vitalis-virtus/blockchain-golang/utils"
)

type CommandLine struct {
	blockChain *blockchain.BlockChain
}

// PrintUsage prints help commands to the terminal
func (cli *CommandLine) printUsage() {
	fmt.Println("Usage:")
	fmt.Println(" add -block BLOCK_DATA - add a block to the chain")
	fmt.Println(" print - Prints the blocks in the chain")
}

// validateArguments allow us to validate arguments which we passed throw the command line
func (cli *CommandLine) validateArguments() {
	if len(os.Args) < 2 {
		cli.printUsage()
		runtime.Goexit() // exits the application throw the shutting down the goroutine
	}
}

func (cli *CommandLine) addBlock(data string) {
	cli.blockChain.AddBlock(data)
	fmt.Println("Block added!")
}

func (cli *CommandLine) printChain() {
	iter := cli.blockChain.Iterator()

	for {
		block := iter.Next()
		fmt.Printf("Prev hash: %x\n", block.PrevHash)
		fmt.Printf("Data in block: %s\n", block.Data)
		fmt.Printf("Block hash: %x\n", block.Hash)
		fmt.Printf("Nonce: %v\n", block.Nonce)
		pow := blockchain.NewProof(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()

		if len(block.PrevHash) == 0 {
			break
		}
	}
}

func (cli *CommandLine) Run() {
	cli.validateArguments()

	addBlockCmd := flag.NewFlagSet("add", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("print", flag.ExitOnError)
	addBlockData := addBlockCmd.String("block", "", "Block data")

	switch os.Args[1] {
	case "add":
		err := addBlockCmd.Parse(os.Args[2:])
		utils.Handle(err)
	case "print":
		err := printChainCmd.Parse(os.Args[2:])
		utils.Handle(err)
	default:
		cli.printUsage()
		runtime.Goexit()
	}

	if addBlockCmd.Parsed() {
		if *addBlockData == "" {
			addBlockCmd.Usage()
			runtime.Goexit()
		}
		cli.addBlock(*addBlockData)
	}

	if printChainCmd.Parsed() {
		cli.printChain()
	}
}

func main() {
	err := godotenv.Load()
	if err != nil {
		utils.Handle(fmt.Errorf("Error loading .env file: %s", err))
	}

	defer os.Exit(0)
	chain := blockchain.InitBlockChain()
	defer chain.Database.Close()

	cli := CommandLine{chain}
	cli.Run()
}
