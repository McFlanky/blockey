package node

import (
	"testing"

	"github.com/McFlanky/blocker/types"
	"github.com/McFlanky/blocker/util"
	"github.com/stretchr/testify/assert"
)

func TestChainHeight(t *testing.T) {
	chain := NewChain(NewMemoryBlockStore())
	for i := 0; i < 100; i++ {
		block := util.RandomBlock()
		assert.Nil(t, chain.AddBlock(block))
		assert.Equal(t, chain.Height(), i)
	}
}

func TestAddBlock(t *testing.T) {
	chain := NewChain(NewMemoryBlockStore())
	for i := 0; i < 100; i++ {
		var (
			block     = util.RandomBlock()
			blockHash = types.HashBlock(block)
		)
		assert.Nil(t, chain.AddBlock(block))

		fetchedBlock, err := chain.GetBlockByHash(blockHash)
		assert.Nil(t, err)
		assert.Equal(t, block, fetchedBlock)

		fetchedBlockByHeight, err := chain.GetBlockByHeight(i)
		assert.Nil(t, err)
		assert.Equal(t, block, fetchedBlockByHeight)
	}
}