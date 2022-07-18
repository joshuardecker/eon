package transactions

// The transaction that gives block reward. Only 1 of these are in each block. Has a pre-defined lock time.
type PLuX struct {
	BlockReward uint64

	LuckyMiner string

	Weight uint32
}

// A basic transaction from 1 person to another. Has a customisable lock time.
type BLuX struct {
	TxInput       []Input
	PubKey        string
	TotalCoinSent uint64
	Signature     []byte
	Msg           []byte

	Weight uint32

	TxOutput []OutPut
}

// A more advanced transaction that will have very basic scripting functionality.
type ALuX struct{}

// The input of a transaction, aka where are you getting coin from?
type Input struct {
	FromBlockHeight uint32
	TxHash          []byte
	Index           uint32
}

// The output of a transaction, aka where is the coin going? Any scripts?
type OutPut struct {
	AddressTo string
	Amount    uint64
	Index     uint32

	Script string
}
