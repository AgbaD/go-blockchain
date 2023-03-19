package blockchain

import (
	"fmt"
	"log"

	badger "github.com/dgraph-io/badger"
)

// we need to create a path to the db
const dbPaths = "./tmp/blocks"

// a blockchain contains multiple blocks

type BlockChain struct {
	// array of pointers to blocks
	// By making the first letter upper case
	// it makes the field public

	Blocks []*Block
	// a poiinter to the badger database
	lastHash []byte
	Database *badger.DB
}

// to be able to get the blocks in the blockchain

// type BlockChainIterator struct {
// 	CurrentHash []byte
// 	Database    *badger.DB
// }

// returns a pointer to the blockchain
func InitBlockchain() {
	// Takes an array of blocks pointers
	// with a call to the genesis function as the first element of the array

	// return &BlockChain{[]*Block{Genesis()}}

	// // where the db will store the keys and metadata
	// // where the db will store the value

	// // open up the DB
	// // returns a tupule with a pointer to the DB
	// // and error

	// // we can access badger DB using two functions
	// // update function to read and write
	// // view fuction to read only

	// // we're passing in a closure which takes a pointer to a badger
	// // transaction and passes back an error
	// // we have access to the transaction so we can do stuff
	
	// 	// check if a blockchain has been created in the db
	// 	// if it has create an instance of the blockchain in memory
	// 	// and get the last hash in the disk db and push it into the instance
	// 	//
	// 	// if there is no exiting blockchain
	// 	// create a genesis block
	// 	// store in db
	// 	// save the last hash as genesis block hash
	// 	// create a new blockchain instance with last hash pointing to genesis block
	// 	//
	// 	// lh means last hash
	// 	// if last hash does not exist, we know we dont have a blockchain

	db, err := badger.Open(badger.DefaultOptions(dbPaths))
	if err != nil {
		log.Panic(err)
	}
	err = db.Update(func(txn *badger.Txn) error {
		// Your code here…
		if _, err := txn.Get([]byte("lh")); err == badger.ErrKeyNotFound {
			fmt.Println("No existing blockchain found")
			genesis := Genesis()
			fmt.Println("Genesis Block created ")
		}
		return nil
	  })
	  

}

// any blockchain struct instance can use this function
func (chain *BlockChain) AddBlock(data string) {
	fmt.Println(chain)
	prevBlock := chain.Blocks[len(chain.Blocks)-1]
	block := CreateBlock(data, prevBlock.Hash)
	chain.Blocks = append(chain.Blocks, block)
}

// we need to create a persistence layer which uses a key-value
// storage DB to store our blockchain

// we'll be using badgerDB, a native golang key-value storage DB
// no tables
// everything os created using key value pairs

//
//
// We need to deide how we will store our blockchain data
// for bitcoin we have two main groups
// blocks (stored with metadata which describes the blocks on the chain)
// and chain state object (state of a chain and current unspent transaction output)
// each block has its on seperate file on the disk in bitcoin
// but we wont do that here

// turn our blockchain to a blockchain iterator
// so we can iterate and get data


// we'll be iterating backwards
