package ellip

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/hex"
	"os"

	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/crypto/sha3"
)

// Generates and saves the private key locally. Returns nothing.
func generatePrivateKey() {

	// Create the key
	key, keyErr := ecdsa.GenerateKey(crypto.S256(), rand.Reader)

	// Error in making key?
	if keyErr != nil {

		panic(keyErr)
	}

	// Save the key
	saveErr := crypto.SaveECDSA("saves/key", key)

	// If the saves folder and the key file are not generated
	if saveErr != nil && os.IsNotExist(saveErr) {

		// Make the saves folder
		saveErr = os.Mkdir("saves", 0750)

		// Couldnt make the saves folder
		if saveErr != nil {

			panic(saveErr)
		}

		// Save the private key
		saveErr = crypto.SaveECDSA("saves/key", key)

		// Couldnt save the key
		if saveErr != nil {

			panic(saveErr)
		}

	} else if saveErr != nil {
		// Couldnt save even though the files existed

		panic(saveErr)
	}
}

// Gets the private key and generates the public key. Returns the public and private key.
func GetKeyPair() (publicKey, privateKey []byte) {

	// Load the private key
	key, readErr := crypto.LoadECDSA("saves/key")

	// If there was an error
	if readErr != nil {

		// If the key hasnt been generated
		if os.IsNotExist(readErr) {

			generatePrivateKey()

			// Loads the newly generated private key
			key, readErr = crypto.LoadECDSA("saves/key")

			// If an error occured loading the private key
			if readErr != nil {

				panic(readErr)
			}

		} else {
			// Had a error getting the key but it was already generated

			panic(readErr)
		}
	}

	// Get the public key
	publicKey = elliptic.Marshal(crypto.S256(), key.X, key.Y)

	// Define the private key
	privateKey = make([]byte, 32)

	// Cool binary stuff that works
	binaryBlob := key.D.Bytes()
	copy(privateKey[32-len(binaryBlob):], binaryBlob)

	return publicKey, privateKey
}

// Get the hash of your public key. Returns a string of the hex of the hashed public key.
func PubKeyHashStr() string {

	// Gets the public key
	publicKey, _ := GetKeyPair()

	// Define the hash
	publicKeyHash := make([]byte, 32)

	// Do the hash
	sha3.ShakeSum256(publicKeyHash, publicKey)

	return hex.EncodeToString(publicKeyHash)
}
