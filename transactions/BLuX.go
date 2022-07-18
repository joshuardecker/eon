package transactions

import "encoding/json"

func (b *BLuX) AddInput(blockHeight uint32, TxHash string, Index uint32) error {

	// Make an input
	i := new(Input)

	i.FromBlockHeight = blockHeight
	i.TxHash = TxHash
	i.Index = Index

	return nil
}

func (b *BLuX) GetWeight() uint32 {

	txAsBytes, jsonErr := json.Marshal(b)

	if jsonErr != nil {
		panic(jsonErr)
	}

	b.Weight = uint32(len(txAsBytes))

	return b.Weight
}
