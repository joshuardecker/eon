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

type AuthorityEngine struct {

	// The config used by the engine.
	config *config.Config

	// The nodes private key.
	privateKey *ecdsa.PrivateKey
}

// Takes the given block and signs it with the nodes private key.
func (a *AuthorityEngine) ValidateBlock(b *block.Block) error {

	// Calculate the block hash.
	blockHash := b.CalcHash()

	// Sign the block hash.
	sig, err := curve.Sign(a.privateKey, blockHash.GetBytes())

	// If an error occured signing the hash.
	if err != nil {

		return ERR_SIGNATURE
	}

	// Set the signature.
	b.Head.SetSig(sig)

	return nil
}

// Returns true if the block is valid, false if it isnt (according to the config given for PoA Engine).
func (a *AuthorityEngine) VerifyBlock(b *block.Block) bool {

	// Uncompress the coinbase.
	coinbase, keyErr := curve.UncompressPub(b.Head.GetCoinbase())

	// If an error occured, the block is assumed invalid.
	if keyErr != nil {

		return false
	}

	// If the coinbase is not equal to the trusted key.
	if coinbase != a.config.TrustedKey {

		return false
	}

	blockHash := b.GetHash()

	// Returns true if the sig is valid by the coinbase, false if not.
	return curve.VerifySign(coinbase, blockHash.GetBytes(), b.Head.GetSig())
}
