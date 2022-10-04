package poa

import (
	"crypto/ecdsa"

	"github.com/Sucks-To-Suck/Eon/core/config"
)

type AuthorityEngine struct {

	// The config used by the engine.
	config *config.Config

	// The nodes private key.
	privateKey *ecdsa.PrivateKey
}

func (a *AuthorityEngine) VerifyBlock() bool { return true }
