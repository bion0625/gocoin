package blockchain

import (
	"sync"
)

type blockChain struct {
	NewestHash	string	`json:"newestHash"`
	Height		int		`json:"height"`
}

var b *blockChain
var once sync.Once

func (b *blockChain) AddBlock(data string) {
	block := createBlcok(data, b.NewestHash, b.Height)
	b.NewestHash = block.Hash
	b.Height = block.Height
}

func BlockChain() *blockChain {
	if b == nil {
		once.Do(func() {
			b = &blockChain{"", 0}
			b.AddBlock("Genesis Block")
		})
	}
	return b
}
