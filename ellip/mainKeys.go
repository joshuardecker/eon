package ellip

import (
	"crypto/ecdsa"
)

// May be used later
type MainKey struct {
	privKey ecdsa.PrivateKey
	pubKey  []byte
	loaded  bool
}

// Gets the private key and generates the public key. Returns nothing.
func (m *MainKey) GetMainKeyPair() {

	m.pubKey, m.privKey = GetKeyPair("key")

	m.loaded = true
}

func (m *MainKey) SignMsgWithMain() (msgHash, sig []byte) {

	if !m.loaded {

		m.GetMainKeyPair()
	}

	return SignRandMsg(&m.privKey)
}

func (m *MainKey) MainKeyHash() string {

	if !m.loaded {

		m.GetMainKeyPair()
	}

	return PubKeyHashStr(m.pubKey)
}
