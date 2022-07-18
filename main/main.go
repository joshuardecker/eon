package main

import (
	"fmt"

	"github.com/Sucks-To-Suck/LuncheonNetwork/blockchain"
	"github.com/Sucks-To-Suck/LuncheonNetwork/ellip"
	"github.com/Sucks-To-Suck/LuncheonNetwork/utilities"
)

func main() {

	fmt.Println("*****")
	//****
	// Test new features area

	msg, sig := ellip.SignRandMsg()
	pub, _ := ellip.GetKeyPair()

	if ellip.ValidateSig(pub, msg, sig) {

		fmt.Println("Worked!")

		fmt.Println(len(pub))
	} else {
		fmt.Println("Didnt work")
	}

	t := new(utilities.Time)
	t.CurrentTime()

	// Test new features area
	//****
	fmt.Println("*****")

	//****
	// Starts Mining

	miner := new(blockchain.Miner)

	block := new(blockchain.Block)
	//block.CreateBlock()

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
