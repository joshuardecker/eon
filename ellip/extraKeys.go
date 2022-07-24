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

// Function gets and if needed generates a public and private key pair.
// Input the save file name to get / make the key pair.
// Returns the public key, and then the private key.
func GetKeyPair(saveName string) ([]byte, ecdsa.PrivateKey) {

	// Load the private key
	privateKey, readErr := crypto.LoadECDSA("saves/" + saveName)

	// If there was an error
	if readErr != nil {

		// If the key hasnt been generated
		if os.IsNotExist(readErr) {

			// Creates the main private key
			generateRandPrivKey(saveName)

			// Loads the newly generated private key
			privateKey, readErr = crypto.LoadECDSA("saves/" + saveName)

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
	publicKey := elliptic.Marshal(crypto.S256(), privateKey.X, privateKey.Y)

	return publicKey, *privateKey
}

// Generates and saves a new private key into the saveName provided.
// Input needed is the name of the new key save.
// Function only intended to be used by GetKeyPair().
// Returns nothing.
func generateRandPrivKey(saveName string) {

	// Create the key
	key, keyErr := ecdsa.GenerateKey(crypto.S256(), rand.Reader)

	// Error in making key?
	if keyErr != nil {

		panic(keyErr)
	}

	// Save the key
	saveErr := crypto.SaveECDSA("saves/"+saveName, key)

	// If the saves folder and the key file are not generated
	if saveErr != nil && os.IsNotExist(saveErr) {

		// Make the saves folder
		saveErr = os.Mkdir("saves", 0750)

		// Couldnt make the saves folder
		if saveErr != nil {

			panic(saveErr)
		}

		// Save the private key
		saveErr = crypto.SaveECDSA("saves/"+saveName, key)

		// Couldnt save the key
		if saveErr != nil {

			panic(saveErr)
		}

	} else if saveErr != nil {
		// Couldnt save even though the files existed

		panic(saveErr)
	}
}

// This function signs a hash of a randomly generated message.
// The hashing and randomness is used for security.
// Input is the private key that will be used to sign the message.
// Output is the hash of the message, and then the signature.
func SignRandMsg(privKey *ecdsa.PrivateKey) (msgHash, sig []byte) {

	// Generates the random message
	msg := randomMessage(32)

	// Makes and does the hash
	msgHash = make([]byte, 32)
	sha3.ShakeSum256(msgHash, msg)

	// Signs the message
	sig, sigErr := crypto.Sign(msgHash, privKey)

	// If an error occured
	if sigErr != nil {

		panic(sigErr)
	}

	// Init the final return sig
	finalSig := make([]byte, 64)

	// Removes the wierd v value at the end that isnt used for verifying signatures
	for index := 0; index < 64; index++ {

		finalSig[index] = sig[index]
	}

	return msgHash, finalSig
}

// This function signs a hash of a inputted message.
// The hashing and randomness is used for security.
// Input is the private key that will be used to sign the message, and the actual message.
// Output is the hash of the message, and then the signature.
func SignMsg(privKey *ecdsa.PrivateKey, msg []byte) (msgHash, sig []byte) {

	// Makes and does the hash
	msgHash = make([]byte, 32)
	sha3.ShakeSum256(msgHash, msg)

	// Signs the message
	sig, sigErr := crypto.Sign(msgHash, privKey)

	// If an error occured
	if sigErr != nil {

		panic(sigErr)
	}

	// Init the final return sig
	finalSig := make([]byte, 64)

	// Removes the wierd v value at the end that isnt used for verifying signatures
	for index := 0; index < 64; index++ {

		finalSig[index] = sig[index]
	}

	return msgHash, finalSig
}

// Function makes a random message.
// Inputs the size of the byte array returned.
// Returns the random message as a byte array.
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

// Hashes the public key inputted for extra security.
// Input a public key to be hashed.
// Outputs the hex string of the public key.
func PubKeyHashStr(pubKey []byte) string {

	// Empty public key?
	if len(pubKey) == 0 {

		panic("input non empty pubKey into PubKeyHashStr")
	}

	// Define the hash
	publicKeyHash := make([]byte, 32)

	// Do the hash
	sha3.ShakeSum256(publicKeyHash, pubKey)

	return hex.EncodeToString(publicKeyHash)
}

// Validates a signature.
// Returns true if valid, false if not valid.
func ValidateSig(publicKey, msgHash, sig []byte) bool {

	return crypto.VerifySignature(publicKey, msgHash, sig)
}
