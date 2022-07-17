package transactions

// The transaction that gives block reward. Only 1 of these are in each block. Has a pre-defined lock time.
type PLuX struct {
	BlockReward uint64

	LuckyMiner []byte

	Weight Weight
}

// A basic transaction from 1 person to another. Has a customisable lock time.
type BLuX struct {
	TxInput []Input
	PubKey  []byte
}

// A more advanced transaction that will have very basic scripting functionality.
type ALuX struct{}

type Input struct {
	BlockNumber uint32
	TxHash      []byte
	Amount      uint64
}
