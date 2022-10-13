package curve

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"os"

	"github.com/Sucks-To-Suck/Eon/eocrypt"
	"github.com/Sucks-To-Suck/Eon/helper/logger"
)

var (
	ERR_GENERATE      = errors.New("Could not generate a public private key pair")
	ERR_LOAD          = errors.New("Could not load a public private key pair")
	ERR_SAVE          = errors.New("Could not save a public private key pair")
	ERR_DECODE        = errors.New("Could not decode a public private key pair")
	ERR_UNCOMPRESSKEY = errors.New("The given bytes could not be uncompressed into a public key")
)

// ****
// Curve and Key Generation:

// This loads the Eon public private key pair.
func LoadPrivateKey() (*ecdsa.PrivateKey, error) {

	// Open the Eon .pem file.
	f, openErr := os.Open("eon.pem")
	defer f.Close()

	// If the file could not be opened.
	if openErr != nil {

		logger := logger.NewLogger("Curve")

		logger.LogRed(ERR_LOAD.Error() + openErr.Error())

		return nil, ERR_LOAD
	}

	// Create and load the pem into a Bytes Buffer.
	fileBuff := new(bytes.Buffer)
	_, readErr := fileBuff.ReadFrom(f)

	// If an error occured by reading the file.
	if readErr != nil {

		logger := logger.NewLogger("Curve")

		logger.LogRed(ERR_LOAD.Error() + readErr.Error())

		return nil, ERR_LOAD
	}

	// Decode the .pem into the PemBlock (if not a valid .pem file, nothing is loaded into the PemBlock).
	pemBlock, _ := pem.Decode(fileBuff.Bytes())

	// Decode the PemBlock Bytes back into a private key.
	p, decodeErr := x509.ParseECPrivateKey(pemBlock.Bytes)

	// If the private key could not be decoded.
	if decodeErr != nil {

		logger := logger.NewLogger("Curve")

		logger.LogRed(ERR_DECODE.Error() + readErr.Error())

		return nil, ERR_DECODE
	}

	return p, nil
}

// Generates and returns an ECDSA private key struct. If any errors occured, it is logged and given by the function.
// It also saves the generated keys onto the hardisk under 'eon.pem'.
func GenerateKeys() (*ecdsa.PrivateKey, error) {

	// Generates a private public key pair on the secp256r1 elliptical curve.
	private, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)

	// If an error occured, log it and return the error.
	if err != nil {

		logger := logger.NewLogger("Curve")

		logger.LogRed(ERR_GENERATE.Error() + err.Error())

		return nil, ERR_GENERATE
	}

	// Save the generated private key.
	err = SavePrivateKey(private)

	// If the keys have been generated but couldnt be saved.
	if err != nil {

		logger := logger.NewLogger("Curve")

		logger.LogRed(ERR_SAVE.Error() + err.Error())

		return private, ERR_SAVE
	}

	return private, nil
}

// Saves the given private key as 'eon.pem' in the same directory as Eon.
func SavePrivateKey(p *ecdsa.PrivateKey) error {

	// Opens / creates the .pem file.
	f, err := os.Create("eon.pem")
	defer f.Close()

	// If an error occurs creating the save.
	if err != nil {

		logger := logger.NewLogger("Curve")

		logger.LogRed(ERR_SAVE.Error() + err.Error())

		return ERR_SAVE
	}

	// Create a bytes buffer and attach it to the encoder.
	keyBytes, _ := x509.MarshalECPrivateKey(p)

	// Save the bytes buffer into a pem block and onto the hardisk.
	pem.Encode(f, &pem.Block{
		Type:  "ECDSA Public Private Key Pair",
		Bytes: keyBytes,
	})

	return nil
}

// Curve and Key Generation:
// ****

// ****
// Private Key Stuff:

// Signs the given message with the private key. Returns the signature and any errors.
func Sign(p *ecdsa.PrivateKey, msg []byte) ([]byte, error) {

	return ecdsa.SignASN1(rand.Reader, p, eocrypt.HashBytes(msg).Bytes())
}

// Private Key Stuff:
// ****

// ****
// Public Key Stuff:

// Checks whether a given message was signed by the given public key. Returns true if signed by them, false otherwise.
func VerifySign(p *ecdsa.PublicKey, msg []byte, sig []byte) bool {

	return ecdsa.VerifyASN1(p, eocrypt.HashBytes(msg).Bytes(), sig)
}

// Returns the compressed form of a public key.
func CompressPub(p *ecdsa.PublicKey) []byte {

	return elliptic.MarshalCompressed(elliptic.P256(), p.X, p.Y)
}

// Uncompresses the given bytes into a public key. Returns an error if occured.
func UncompressPub(k []byte) (*ecdsa.PublicKey, error) {

	x, y := elliptic.UnmarshalCompressed(elliptic.P256(), k)

	// X == nil if it failed to Unmarshal the given bytes 'k'.
	if x == nil {

		return nil, ERR_UNCOMPRESSKEY
	}

	// Return the uncompressed public key.
	return &ecdsa.PublicKey{
		Curve: elliptic.P256(),
		X:     x,
		Y:     y,
	}, nil
}

// Public Key Stuff:
// ****
