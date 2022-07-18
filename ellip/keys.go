package ellip

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"os"

	"github.com/ethereum/go-ethereum/crypto"
)

func generatePrivateKey() {

	key, keyErr := ecdsa.GenerateKey(crypto.S256(), rand.Reader)

	if keyErr != nil {
		panic(keyErr)
	}

	saveErr := crypto.SaveECDSA("saves/key", key)

	if saveErr != nil {
		panic(saveErr)
	}
}

func GetKeyPair() (publicKey, privateKey []byte) {

	key, readErr := crypto.LoadECDSA("saves/key")

	if readErr != nil {

		if os.IsNotExist(readErr) {

			generatePrivateKey()

			key, readErr = crypto.LoadECDSA("saves/key")

		} else {

			panic(readErr)
		}
	}

	if readErr != nil {

		panic(readErr)
	}

	publicKey = elliptic.Marshal(crypto.S256(), key.X, key.Y)

	privateKey = make([]byte, 32)

	binaryBlob := key.D.Bytes()
	copy(privateKey[32-len(binaryBlob):], binaryBlob)

	return publicKey, privateKey
}
