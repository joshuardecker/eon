package elliptic

import (
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/crypto"
)

func GenerateKeys() ecdsa.PrivateKey {

	key, keyErr := crypto.GenerateKey()

	if keyErr != nil {
		panic(keyErr)
	}

	return *key
}
