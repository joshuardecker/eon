package curve

import (
	"fmt"
	"testing"

	"github.com/Sucks-To-Suck/Eon/eocrypt"
)

func TestNoSavedKeys(t *testing.T) {

	// Make sure to delete eon.pem before running to fully test.
	// Both printed hashes should be the same.

	_, err := LoadPrivateKey()

	if err != nil {

		p, _ := GenerateKeys()
		fmt.Printf("1:%x\n", *eocrypt.HashBytes(p.D.Bytes()))
	}

	p, _ := LoadPrivateKey()
	fmt.Printf("2:%x\n", *eocrypt.HashBytes(p.D.Bytes()))
}

func TestSignatures(t *testing.T) {

	// Load the public private keys.
	p, err := LoadPrivateKey()

	// If they dont exist yet, generate and load them.
	if err != nil {

		GenerateKeys()

		p, _ = LoadPrivateKey()
	}

	// The sample message that will be signed.
	msg := []byte("This is my sweet message!")

	// Get the signature.
	sig, _ := Sign(p, msg)

	// Should Print 'True'. If so, Signing works.
	fmt.Println(VerifySign(&p.PublicKey, msg, sig))

	msg2 := []byte("A different message!")

	// Should Print 'False'. If so, Verifying false sigs works.
	fmt.Println(VerifySign(&p.PublicKey, msg2, sig))
}
