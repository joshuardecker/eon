package transactions

import (
	"errors"
)

// Sets the weight of the PLuX
func (p *PLuX) GetWeight() error {

	weightErr := p.weight.WeightPLuX(p)

	if weightErr != nil {

		return weightErr
	}

	return nil
}

// Sets the block reward input of the PLuX
func (p *PLuX) SetBlockReward(blockReward uint64) {

	p.blockReward = blockReward
}

// Sets the address hash that will receive the block reward
func (p *PLuX) SetLuckyMiner(luckyMiner []byte) error {

	// If input nothing / nil
	if luckyMiner == nil {

		return errors.New("Input a non nil id for the lucky miner.")
	}

	p.luckyMiner = luckyMiner

	return nil
}

// Adds transaction fees to the block reward
func (p *PLuX) CalculateRewards(txFees uint64) {

	p.blockReward += txFees
}
