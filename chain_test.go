package blockchain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	ch        *Blockchain
	blockData = map[string][]byte{
		"genesis": []byte("genesis block!"),
		"second":  []byte("This is the second block"),
		"third":   []byte("And a third one"),
	}
)

func TestAll(t *testing.T) {
	ch = NewBlockchain(Sha256, Hashcash, blockData["genesis"])
	ch.Mine(blockData["second"])
	ch.Mine(blockData["third"])

	t.Run("all", func(t *testing.T) {
		t.Run("VerifySuccessfulGenesisBlockchain", VerifySuccessfulGenesisBlockchain)
		t.Run("VerifyChain", VerifyChain)
	})
}

func VerifySuccessfulGenesisBlockchain(t *testing.T) {
	// then
	assert.Equal(t, 3, len(ch.blocks))
	assert.Equal(t, blockData["third"], ch.Tip().Data)
}

func VerifyChain(t *testing.T) {
	// 3 -> 2
	tip := ch.Tip()
	prevBlock := ch.PreviousBlock(tip)
	assert.NotNil(t, prevBlock)
	assert.Equal(t, blockData["second"], prevBlock.Data)

	// 2 -> genesis
	prevBlock = ch.PreviousBlock(prevBlock)
	assert.NotNil(t, prevBlock)
	assert.Equal(t, blockData["genesis"], prevBlock.Data)

	// genesis -> nil
	prevBlock = ch.PreviousBlock(prevBlock)
	assert.Nil(t, prevBlock)
}
