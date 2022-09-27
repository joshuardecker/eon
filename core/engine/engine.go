package engine

// Here the consensus engine is defined, tieing together the various proof of consensus engines togethe under one interface.
// To see what the engines are doing, look under the Proof Folders in the main directory.
// To see how the engine is interacted with, look in threader.
type ConsensusEngine interface {
	Verify()
}
