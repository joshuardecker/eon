package miner

import (
	"github.com/Sucks-To-Suck/LuncheonNetwork/core/events/block"
	"github.com/Sucks-To-Suck/LuncheonNetwork/logger"
	"github.com/Sucks-To-Suck/LuncheonNetwork/miner/target"
	"github.com/Sucks-To-Suck/LuncheonNetwork/miner/worker"
	"github.com/Sucks-To-Suck/LuncheonNetwork/util"
)

// The miner is the overseer of the workers.
// It will manage creation, management, and deletion of workers.
// All the miner needs is a block to mine, and it will get to work, no questions asked.
type Miner struct {
	block *block.Block

	logger *logger.Logger
	time   *util.Time

	isMining bool
}

// Creates and returns a new miner.
// Only one of these should exist.
func New(block *block.Block) *Miner {

	m := new(Miner)

	m.block = block
	m.logger = logger.NewLogger("Miner")
	m.time = &util.Time{}
	m.isMining = true

	return m
}

// Updates the miner to use a new block.
// Doesnt need to stop workers, as they will just mine the new block.
func (m *Miner) Update(block *block.Block) {

	m.isMining = false

	// Switches to a new block
	m.block = block
}

func (m *Miner) Run() {

	m.logger.LogBlue("Starting up the Miner...")

	target := target.Unpack(m.block.PackedTarget)

	w1 := worker.New(target)
	w2 := worker.New(target)
	w3 := worker.New(target)

	go w1.Start(m.block, 0)
	go w2.Start(m.block, 1)
	go w3.Start(m.block, 2)

	for {

		for m.isMining {

			if m.block.BlockHash != "" {

			}
		}
	}
}
