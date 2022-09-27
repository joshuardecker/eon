package gas

import "math/big"

// Gas (measured in Virtual Gas Units) tracks data transaction size.
// This with gas price is used to determine the fees on the layer 1 threader is applicable.
// 1 Vgu = 1 byte taken up by a data transaction.
type Gas uint64

// Determines the amount of gas you get per native token.
// This also means that the highest gas price is 1 token per gas.
type GasPrice *big.Int
