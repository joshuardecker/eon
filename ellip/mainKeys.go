package ellip

import (
	"crypto/ecdsa"
	"encoding/hex"
)

// The main key used by nodes.
// The struct simply makes it easier to interact with.
type MainKey struct {
	privKey ecdsa.PrivateKey
	pubKey  []byte
	loaded  bool
}

// Gets the private key and generates the public key (into the struct).
// Returns nothing.
func (m *MainKey) GetMainKeyPair() {

	m.pubKey, m.privKey = GetKeyPair("key")

	m.loaded = true
}

// Signs a message with the main keys.
// Also does not need keys to be loaded / generated before hand.
// Returns the hash of the message and the signature of that hash.
func (m *MainKey) SignMsgWithMain() (msgHash, sig []byte) {

	// If the keys have not been loaded.
	if !m.loaded {

		m.GetMainKeyPair()
	}

	return SignRandMsg(&m.privKey)
}

// Gets the hash of the main public key.
// Also does not need keys to be loaded / generated before hand.
// Returns the hex string of that hash.
func (m *MainKey) MainKeyHash() string {

	if !m.loaded {

		m.GetMainKeyPair()
	}

	return PubKeyHashStr(m.pubKey)
}

// This function gets and returns the hex string version of the public key.
func (m *MainKey) GetPubKeyStr() string {

	if !m.loaded {

		m.GetMainKeyPair()
	}

	return hex.EncodeToString(m.pubKey)
}
