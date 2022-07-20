package transactions

import (
	"encoding/hex"
	"encoding/json"
	"errors"
)

func (b *BLuX) CreateBLuX(pubKey []byte, fee uint64, msg []byte, sig []byte) error {

	if len(b.TxInput) == 0 || len(b.TxOutput) == 0 {

		return errors.New("make BLuX inputs and outputs before making the tx")
	}

	b.PubKey = hex.EncodeToString(pubKey)
	b.Fee = fee
	b.Msg = hex.EncodeToString(msg)
	b.Signature = hex.EncodeToString(sig)

	b.Weight = b.getWeight()

	// Get the hash of the tx

	return nil
}

// Creates the BLuX input needed for the transaction.
// Input is where the coin is coming from.
// Inputs are the blockHeight from where this is coming from,
// the txHash to find the transaction.
// Returns an error if occured, otherwise nil.
func (b *BLuX) CreateBLuXInput(bHeight uint32, txHash string, input uint32) error {

	if len(txHash) != 64 {

		return errors.New("hash is not correct length")
	}

	txInput := new(Input)

	txInput.FromBlockHeight = bHeight
	txInput.TxHash = txHash
	txInput.Index = input

	// Eventually have wallet scan to see if these are valid inputs

	b.TxInput = append(b.TxInput, *txInput)

	return nil
}

// Creates the BLuX output needed for the transaction.
// Output is where the coin is going.
// Inputs are where they are going, and the amount going there.
// Returns an error if occured, otherwise nil.
func (b *BLuX) CreateBLuXOutput(addressTo string, amount uint64) error {

	// In the future, have the wallet check if they have enough coin.

	txOutput := new(OutPut)

	txOutput.AddressTo = addressTo
	txOutput.Amount = amount

	b.TxOutput = append(b.TxOutput, *txOutput)

	return nil
}

func (b *BLuX) getWeight() uint32 {

	txAsBytes, jsonErr := json.Marshal(b)

	if jsonErr != nil {
		panic(jsonErr)
	}

	b.Weight = uint32(len(txAsBytes))

	return b.Weight
}
