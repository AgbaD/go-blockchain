package main

// I decided to learn golang by building a lockchain
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

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

// a blockchain contains multiple blocks
//

type BlockChain struct {
	// array of pointers to blocks
	blocks []*Block
}

// a struct (a data structure to represent to represent our block)
type Block struct {
	// block's hash = Data + prevHash
	Hash []byte
	//any type of data
	Data []byte
	// previous block's hash
	PrevHash []byte

}

// if the block struct is an object]
// defining a function for the struct
// i suppose
func (b *Block) DeriveHash() {
	// a two dimensional slice of bytes
	// takes in byte data and prevHash
	// and the empty byte array
	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	// hash the data
	hash := sha256.Sum256((info))
	b.Hash = hash[:]
}

// outputs a pointer to a block
func CreateBlock (data string, prevHash []byte) *Block{
	// reference to a block
	block := &Block{[]byte{}, []byte(data), prevHash}
	// using the struct fuction
	// self defined
	block.DeriveHash()
	return block
}

// any blockchain struct instance can use this function
func (chain *BlockChain) AddBlock(data string) {
	prevBlock := chain.blocks[len(chain.blocks) - 1]
	block := CreateBlock(data, prevBlock.Hash)
	chain.blocks = append(chain.blocks, block)
}

// returns a pointer to a block
func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

// returns a pointer to the blockchain
func InitBlockchain() *BlockChain{
	// Inside is an array of blocks pointers with a call to the genesis function
	return &BlockChain{[]*Block{Genesis()}}
}

func main () {
	chain := InitBlockchain()
	chain.AddBlock("First block")
	chain.AddBlock("Second block")
	chain.AddBlock("Third block")

	for _, block := range chain.blocks {
		// string interpolation
		fmt.Printf("Previous hash: %x\n", block.PrevHash)
		fmt.Printf("Block data: %s\n", block.Data)
		fmt.Printf("Block hash: %x\n", block.Hash)
	}
}
