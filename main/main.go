package main

import (
	"fmt"

	"github.com/Sucks-To-Suck/LuncheonNetwork/blockchain"
	"github.com/Sucks-To-Suck/LuncheonNetwork/ellip"
)

func main() {

	fmt.Println("*****")
	//****
	// Test new features area

	msg, sig := ellip.SignRandMsg()
	pub, _ := ellip.GetKeyPair()

	if ellip.ValidateSig(pub, msg, sig) == true {

		fmt.Println("Worked!")
	} else {
		fmt.Println("Didnt work")
	}

	// Test new features area
	//****
	fmt.Println("*****")

	//****
	// Starts Mining

	miner := new(blockchain.Miner)

	block := new(blockchain.Block)
	block.ConstructBlock()

	targetErr := miner.InputTarget(0x1dffffff) //1d00ffff

	if targetErr != nil {

		panic(targetErr)
	}

	_, minerErr := miner.Start(*block)

	if minerErr != nil {

		panic(minerErr)
	}

	// Starts the mining
	//****
}
