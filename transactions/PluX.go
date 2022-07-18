package transactions

import (
	"encoding/json"
	"errors"
)

// Sets the block reward input of the PLuX
func (p *PLuX) SetBlockReward(blockReward uint64) {

	p.BlockReward = blockReward
}

// Adds transaction fees to the block reward
func (p *PLuX) CalculateRewards(txFees uint64) {

	p.BlockReward += txFees
}

// Sets the address hash that will receive the block reward
func (p *PLuX) SetLuckyMiner(luckyMiner string) error {

	// If input nothing / nil
	if len(luckyMiner) == 0 {

		return errors.New("input an id for the lucky miner")
	}

	p.LuckyMiner = luckyMiner

	return nil
}

// Gets the weight of the PLuX
func (p *PLuX) GetWeight() uint32 {

	txAsBytes, jsonErr := json.Marshal(p)

	if jsonErr != nil {
		panic(jsonErr)
	}

	p.Weight = uint32(len(txAsBytes))

	return p.Weight
}
