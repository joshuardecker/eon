package threader

import (
	"crypto/ecdsa"
	"math/big"
	"net"
)

// The config is the configuration given to a threader. This includes the chain ID of the network, whether its PoA, PoB, or PoW,
// a trusted IP for syncing, and which signer of blocks you trust (can just be yourself).
type Config struct {

	// Id of the chain the threader is running on.
	chainId *big.Int

	// Type of Proof used by the chain. ("PoA", "PoB", "PoW").
	proofType string

	// The trusted IP used for syncing.
	trustedIP *net.IP

	// The public key of a trusted source.
	trustedKey *ecdsa.PublicKey
}
