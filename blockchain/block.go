package blockchain

import (
	"bytes"
	"encoding/gob"
	"log"
)

// a struct (a data structure to represent to represent our block)
type Block struct {
	// block's hash = Data + prevHash
	Hash []byte
	//any type of data
	Data []byte
	// previous block's hash
	PrevHash []byte
	// the nonce value that meets the target for the block
	Nonce int
}

// if the block struct is an object]
// defining a function for the struct
// i suppose
// func (b *Block) DeriveHash() {
// 	// a two dimensional slice of bytes
// 	// takes in byte data and prevHash
// 	// dummy way of hashing for now
// 	//
// 	// and the empty byte array
// 	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
// 	// hash the data
// 	hash := sha256.Sum256(info)
// 	b.Hash = hash[:]
// }

// outputs a pointer to a block
func CreateBlock(data string, prevHash []byte) *Block {
	// reference to a block
	block := &Block{[]byte{}, []byte(data), prevHash, 0}
	// using the struct fuction
	// self defined
	// block.DeriveHash()

	// run the pow algorithm
	pow := NewProof(block)
	nonce, hash := pow.Run()
	// fmt.Println(nonce)
	// fmt.Println(hash)
	// fmt.Println()

	block.Hash = hash[:]
	block.Nonce = nonce
	return block
}

// returns a pointer to a block
func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}


// Badger DB only accepts arrays or slices of bytes
// we need a way to convert our block from and to the 
// right formats

func(b *Block) Serialize() []byte{
	// a resolve
	var res bytes.Buffer
	// we create an encoder on the resolve
	encoder := gob.NewEncoder(&res)
	err := encoder.Encode(b) 
	// function returns an error so error handling
	Handle(err)
	// Byte representation of the block
	return res.Bytes()
}

func Deserialize(data []byte ) *Block{
	var block Block
	decoder := gob.NewDecoder(bytes.NewReader(data))
	err := decoder.Decode(&block)
	Handle(err)
	// Byte representation of the block
	return &block
}

func Handle(err error) {
	if err != nil {
		log.Panic(err)
	}
}

