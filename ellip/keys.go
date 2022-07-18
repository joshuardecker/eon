package ellip

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/hex"
	"io"
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
func GetKeyPair() (publicKey []byte, privateKey *ecdsa.PrivateKey) {

	// Load the private key
	privateKey, readErr := crypto.LoadECDSA("saves/key")

	// If there was an error
	if readErr != nil {

		// If the key hasnt been generated
		if os.IsNotExist(readErr) {

			generatePrivateKey()

			// Loads the newly generated private key
			privateKey, readErr = crypto.LoadECDSA("saves/key")

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
	publicKey = elliptic.Marshal(crypto.S256(), privateKey.X, privateKey.Y)

	// Define the private key
	//privateKey = make([]byte, 32)

	// Cool binary stuff that works
	//binaryBlob := key.D.Bytes()
	//copy(privateKey[32-len(binaryBlob):], binaryBlob)

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

// Gives a random message with only a size (bytes returned) input needed. Returns a byte array of the random message.
func randomMessage(arraySize uint) []byte {

	// Creates the byte buffer
	bytesBuffer := make([]byte, arraySize)

	// Read a random place for the message
	_, readErr := io.ReadFull(rand.Reader, bytesBuffer)

	// Not read correctly?
	if readErr != nil {

		panic(readErr)
	}

	return bytesBuffer
}

func SignRandMsg() (msgHash, sig []byte) {

	_, privateKey := GetKeyPair()

	msg := randomMessage(32)

	msgHash = make([]byte, 32)
	sha3.ShakeSum256(msgHash, msg)

	sig, sigErr := crypto.Sign(msgHash, privateKey)

	if sigErr != nil {

		panic(sigErr)
	}

	finalSig := make([]byte, 64)

	for index := 0; index < 64; index++ {

		finalSig[index] = sig[index]
	}

	return msgHash, finalSig
}

func ValidateSig(publicKey, msgHash, sig []byte) bool {

	return crypto.VerifySignature(publicKey, msgHash, sig)
}
