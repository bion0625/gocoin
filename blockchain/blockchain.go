package blockchain

import (
	"crypto/sha256"
	"fmt"
	"sync"
)

type Block struct {
	Data     string	`json:"data"`
	Hash     string	`json:"hash"`
	PrevHash string	`json:"prehash,omitempty"`
	Height	 int	`json:"height"`
}

type blockChain struct {
	blocks []*Block
}

var b *blockChain
var once sync.Once

func (b *Block) calculateHash() {
	hash := sha256.Sum256([]byte(b.Data + b.PrevHash))
	b.Hash = fmt.Sprintf("%x", hash)
}

func getLastHash() string {
	totalBlocks := len(GetBlockChain().blocks)
	if totalBlocks == 0 {
		return ""
	}
	return GetBlockChain().blocks[totalBlocks-1].Hash
}

func createBlock(data string) *Block {
	newBlock := Block{data, "", getLastHash(), len(GetBlockChain().blocks)+1}
	newBlock.calculateHash()
	return &newBlock
}

func (b *blockChain) AddBlock(data string){
	b.blocks = append(b.blocks, createBlock(data))
}

func GetBlockChain() *blockChain {
	if b == nil {
		once.Do(func() {
			b = &blockChain{}
			b.AddBlock("Genesis Block")
		})
	}
	return b
}

func (b *blockChain) AllBlocks() []*Block {
	return b.blocks
}

func (b *blockChain) GetBlock(height int) *Block {
	return b.blocks[height-1]
}