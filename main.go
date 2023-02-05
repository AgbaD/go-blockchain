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
	"fmt"
	"strconv"

	"github.com/AgbaD/go-blockchain/blockchain"
)

func main() {
	chain := blockchain.InitBlockchain()
	chain.AddBlock("First block")
	chain.AddBlock("Second block")
	chain.AddBlock("Third block")

	for _, block := range chain.Blocks {
		// string interpolation
		fmt.Printf("Previous hash: %x\n", block.PrevHash)
		fmt.Printf("Block data: %s\n", block.Data)
		fmt.Printf("Block hash: %x\n", block.Hash)

		// Get a new proof for the block
		pow := blockchain.NewProof(block)
		// convert the response of the validation -a boolean- to string format
		fmt.Printf("POW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}
}
