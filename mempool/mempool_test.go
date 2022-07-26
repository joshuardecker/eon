package mempool

import (
	"fmt"
	"testing"

	"github.com/Sucks-To-Suck/LuncheonNetwork/blockchain"
	"github.com/Sucks-To-Suck/LuncheonNetwork/ellip"
	"github.com/Sucks-To-Suck/LuncheonNetwork/wallet"
)

func TestCreateTx(t *testing.T) {

	bc := blockchain.InitBlockchain()
	miner := new(blockchain.Miner)
	key := new(ellip.MainKey)
	wal := wallet.Init(&bc)
	mem := Init(&wal)

	miner.Start(&bc.Blocks[0])

	fmt.Println("Balance:", wal.ScanChainForBalance(key.GetPubKeyStr()))

	tx := wal.CreateTx("kaimorton123", 2000)

	test := mem.AddTx(&tx)
	fmt.Println("Added tx:", test)
}
