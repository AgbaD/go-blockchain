package blockchain

// import "fmt"

// a blockchain contains multiple blocks
//

type BlockChain struct {
	// array of pointers to blocks
	// By making the first letter upper case
	// it makes the field public
	Blocks []*Block
}

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

// any blockchain struct instance can use this function
func (chain *BlockChain) AddBlock(data string) {
	prevBlock := chain.Blocks[len(chain.Blocks)-1]
	block := CreateBlock(data, prevBlock.Hash)
	chain.Blocks = append(chain.Blocks, block)
}

// returns a pointer to a block
func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

// returns a pointer to the blockchain
func InitBlockchain() *BlockChain {
	// Takes an array of blocks pointers
	// with a call to the genesis function as the first element of the array
	return &BlockChain{[]*Block{Genesis()}}
}
