package main

// I decided to learn golang by building a blockchain
// So I'm not really familiar with a lot of concepts yet but
// I plan to learn on the Go!! (see what I did there)
//
// Cheers to learning Golang!!!
// Cheers to building and learning on the blockchain!!!!
//
//////////////////////////////////////////////////////////////////

// Building a blockchain
// A public decentralised, trustless database based on a peer to peer network

// the imports
// a set of tools or libraries to make the task faster

// the imports to github checks online first
// then it checks the local commits

import (
	"flag"
	"fmt"
	"os"
	"runtime"
	"strconv"

	"github.com/AgbaD/go-blockchain/blockchain"
)

type CommandLine struct {
	Blockchain *blockchain.BlockChain
}

func (cli *CommandLine) printUsage() {
	fmt.Println("Usage:")
	fmt.Println("add -block BLOCK_DATA - add a block to the chain")
	fmt.Println("print - Prints the blocks on the chain ")

}

func (cli *CommandLine) ValidateArgs() {
	if len(os.Args) < 2 {
		cli.printUsage()
		// exits the application by shutting down the goroutine
		runtime.Goexit()
	}
}

func (cli *CommandLine) addBlock(data string) {
	cli.Blockchain.AddBlock(data)
	fmt.Println("Added Block")
}

func (cli *CommandLine) printChain() {
	iter := cli.Blockchain.Iterator()
	for {
		block := iter.Next()

		// string interpolation
		fmt.Printf("Previous hash: %x\n", block.PrevHash)
		fmt.Printf("Block data: %s\n", block.Data)
		fmt.Printf("Block hash: %x\n", block.Hash)
		// fmt.Printf("Block Nonce: %x\n", block.Nonce)
		// Get a new proof for the block
		pow := blockchain.NewProof(block)
		// convert the response of the validation -a boolean- to string format
		fmt.Printf("POW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()

		if len(block.PrevHash) == 0 {
			break
		}
	}
}

func (cli *CommandLine) Run() {
	cli.ValidateArgs()

	addBlockCmd := flag.NewFlagSet("add", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("print", flag.ExitOnError)
	addBlockData := addBlockCmd.String("block", "", "Block data")

	switch os.Args[1] {
	case "add":
		err := addBlockCmd.Parse(os.Args[2:])
		blockchain.Handle(err)

	case "print":
		err := printChainCmd.Parse(os.Args[2:])
		blockchain.Handle(err)

	default:
		cli.printUsage()
		runtime.Goexit()
	}

	if addBlockCmd.Parsed() {
		if *addBlockData == "" {
			addBlockCmd.Usage()
			runtime.Goexit()
		} else {
			cli.addBlock(*addBlockData)
		}
	}

	if printChainCmd.Parsed() {
		cli.printChain()
	}

}

func main() {
	// chain := blockchain.InitBlockchain()
	// chain.AddBlock("First block")
	// fmt.Println()
	// chain.AddBlock("Second block")
	// fmt.Println()
	// chain.AddBlock("Third block")
	// fmt.Println()

	// for _, block := range chain.Blocks {
	// 	// string interpolation
	// 	fmt.Printf("Previous hash: %x\n", block.PrevHash)
	// 	fmt.Printf("Block data: %s\n", block.Data)
	// 	fmt.Printf("Block hash: %x\n", block.Hash)
	// 	// fmt.Printf("Block Nonce: %x\n", block.Nonce)
	// 	// Get a new proof for the block
	// 	pow := blockchain.NewProof(block)
	// 	// convert the response of the validation -a boolean- to string format
	// 	fmt.Printf("POW: %s\n", strconv.FormatBool(pow.Validate()))
	// 	fmt.Println()
	// }

	// fail safes to help close the database and give it time
	// garbage collect the keys and values
	defer os.Exit(0)
	chain := blockchain.InitBlockchain()
	// only executes if the go channel is able to exit properly
	defer chain.Database.Close()

	cli := CommandLine{chain}
	cli.Run()

}
