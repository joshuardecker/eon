package poa

import (
	"crypto/ecdsa"
	"errors"

	"github.com/Sucks-To-Suck/Eon/core/config"
	"github.com/Sucks-To-Suck/Eon/eocrypt/curve"
	"github.com/Sucks-To-Suck/Eon/types/block"
)

var (
	ERR_SIGNATURE = errors.New("Couldnt create the signature of the given block.")
)

// The Authority Engine (AE) is the proof of consensus engine for proof of authority, aka listen to a trusted party.
// The trusted party is a public key provided in the eon config. Any blocks created by the trusted party are assumed to be valid.
type AuthorityEngine struct {

	// The config used by the engine.
	config *config.Config

	// The nodes private key.
	privateKey *ecdsa.PrivateKey
}

// Creates and returns an authority engine with the given config and private key.
func NewAuthorityEngine(c config.Config, p ecdsa.PrivateKey) *AuthorityEngine {

	// Create the engine.
	ae := new(AuthorityEngine)

	// Apply the parameters.
	ae.config = &c
	ae.privateKey = &p

	return ae
}

// Takes the given block and signs it with the nodes private key.
func (a *AuthorityEngine) ValidateBlock(b *block.Block) error {

	err := b.Sign(a.privateKey)

	return err
}

// Returns true if the block is valid, false if it isnt (according to the config given for PoA Engine).
func (a *AuthorityEngine) VerifyBlock(b *block.Block) bool {

	// Get the blockhash (the message that was signed).
	bHash := b.Hash()

	// Returns true if the block was signed by the trusted key, false if not.
	return curve.VerifySign(a.config.TrustedKey, bHash.Bytes(), b.Header().Signature())
}
