package transactions

import (
	"encoding/json"

	"github.com/Sucks-To-Suck/LuncheonNetwork/ellip"
)

// Creates a PLuX, where your publickey hash is the winner of the block.
// Will be put in your version of the block, and only the global version if you mine it first.
// Input is the current block reward.
// Returns errors if they occur.
func (p *PLuX) CreatePLuX(blockReward uint64) error {

	// Init the main key
	mainKey := new(ellip.MainKey)

	// Set the variables in the tx
	p.BlockReward = blockReward
	p.LuckyMiner = mainKey.MainKeyHash()

	// Get the tx's weight
	p.Weight = p.GetWeight()

	return nil
}

// Adds transaction fees to the block reward.
// Input is the transaction fees.
// Returns nothing.
func (p *PLuX) CalculateRewards(txFees uint64) {

	p.BlockReward += txFees
}

// Calculates the weight of the PLuX.
// Returns the weight of the PLuX.
func (p *PLuX) GetWeight() uint32 {

	txAsBytes, jsonErr := json.Marshal(p)

	if jsonErr != nil {
		panic(jsonErr)
	}

	p.Weight = uint32(len(txAsBytes))

	return p.Weight
}
