package transactions

import (
	"errors"
)

// Sets the weight of the PLuX
func (p *PLuX) GetWeight() error {

	weightErr := p.Weight.WeightPLuX(p)

	if weightErr != nil {

		return weightErr
	}

	return nil
}

// Sets the block reward input of the PLuX
func (p *PLuX) SetBlockReward(blockReward uint64) {

	p.BlockReward = blockReward
}

// Sets the address hash that will receive the block reward
func (p *PLuX) SetLuckyMiner(luckyMiner []byte) error {

	// If input nothing / nil
	if luckyMiner == nil {

		return errors.New("input a non nil id for the lucky miner")
	}

	p.LuckyMiner = luckyMiner

	return nil
}

// Adds transaction fees to the block reward
func (p *PLuX) CalculateRewards(txFees uint64) {

	p.BlockReward += txFees
}
