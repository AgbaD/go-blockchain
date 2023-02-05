package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"log"
	"math"
	"math/big"
)

// consensus algorithms
// proof of; algorithms

// forcing the network to do computational work to be able
// to sign/add a block to the blockchain
// so they have to be true to the proof the send
// as it can be validated easily by other participants

// Take the data from the block

// create a counter (nonce) which starts at 0

// create a hash of the data + the nonce

// check the hash to see if it meets requirements

// if not; repeat

// Requirements
// The First few bytes must contain 0s
// For bitcoin, it was originally 20 and it changes from time to time

// lets get to it

// static
// in a real blockchain, it has to increase overtime
// due to increase in number of miners growing
// and also computation power
const Difficulty = 18

type ProofOfWork struct {
	Block  *Block
	Target *big.Int
}

// produce a pointer to a POW
func NewProof(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	// it takes 256, which is the number of bytes in one hash
	// the subtracts the difficulty from it
	// then use the target to shift the number of bytes over by ...
	// lsh = left shift
	target.Lsh(target, uint(265-Difficulty))

	pow := &ProofOfWork{b, target}
	return pow
}

// create a method on a struct
func (pow *ProofOfWork) InitData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.Block.PrevHash,
			pow.Block.Data,
			ToHex(int64(nonce)),
			ToHex(int64(Difficulty)),
		},
		[]byte{},
	)
	return data
}

func (pow *ProofOfWork) Run() (int, []byte) {
	var intHash big.Int
	var hash [32]byte

	nonce := 0

	for nonce < math.MaxInt64 {
		// prepare data
		data := pow.InitData(nonce)
		// then hash to sha256
		hash := sha256.Sum256(data)
		fmt.Printf("Hash: \r%x", hash)
		// then convert the hash to  a big int
		intHash.SetBytes(hash[:])
		// then compare the int to the target
		if intHash.Cmp(pow.Target) == -1 {
			break
		} else {
			nonce++
		}
	}
	fmt.Println()
	fmt.Println(nonce)
	fmt.Println(hash)
	fmt.Println()
	return nonce, hash[:]
}


// to validate a block
// relatively simpler than getting the nonce
func(pow *ProofOfWork) Validate() bool {
	var intHash big.Int

	data := pow.InitData(pow.Block.Nonce)
	hash := sha256.Sum256(data)
	fmt.Printf("\r%x", hash)
	intHash.SetBytes(hash[:])

	return intHash.Cmp(pow.Target) == -1
}

// we need to add the nonce and Difficulty to the InitData function; the
// point is the find a value (nonce) for which when added to the
// block data and hashed, it results in a hash that meets the
// requirement placed. Then the miner sends the block and the
// nonce once they are done so any other person
// can take the previous block hash, the proposed block data
// and the nonce and vefiry if it meets the requirements
//
//

// a utility function
func ToHex(num int64) []byte {
	// creates a new bytes buffer
	buff := new(bytes.Buffer)
	// Takes the number and decode into bytes
	// BigEndian defines how we want the bytes to be organised
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}
	return buff.Bytes()
}
