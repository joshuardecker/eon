package config

import (
	"crypto/ecdsa"
	"encoding/json"
	"errors"
	"math/big"
	"net"
	"os"
)

var (
	ERR_LOADCONFIG     = errors.New("Could not load the configuration file")
	ERR_CONFIGNOTEXIST = errors.New("Config could not be loaded because it doesnt exist")
	ERR_DECODECONFIG   = errors.New("Could not decode the config file")

	ERR_ENCODE           = errors.New("Could not encode the config file")
	ERR_FAILEDCREATEFILE = errors.New("Could not create the file for the config to be dumped into")
)

const (
	CONFIG_NAME = "eon.json"     // Name of the config file.
	DUMP_NAME   = "eondump.json" // Name of the config file when dumped.
)

// The config is the configuration given to a threader. This includes the chain ID of the network, whether its PoA, PoB, or PoW,
// a trusted IP for syncing, and which signer of blocks you trust (can just be yourself).
type Config struct {

	// Id of the chain the threader is running on.
	ChainId *big.Int

	// Type of Proof used by the chain. ("PoA", "PoB", "PoW").
	ProofType string

	// The IP of the node.
	IP *net.IP

	// The trusted IP used for syncing.
	TrustedIP *net.IP

	// The public key of a trusted source.
	TrustedKey *ecdsa.PublicKey
}

// Load the default configuration file. Returns any errors if they occured.
func ConfigLoad() (*Config, error) {

	file, openErr := os.Open(CONFIG_NAME)

	// If the config does not exist.
	if os.IsNotExist(openErr) {

		return nil, ERR_CONFIGNOTEXIST
	}

	// Could not load the config.
	if openErr != nil {

		return nil, ERR_LOADCONFIG
	}

	// Create the config.
	c := new(Config)

	// Attempt to decode the json.
	decodeErr := json.NewDecoder(file).Decode(c)

	// If the json config could not be decoded.
	if decodeErr != nil {

		return nil, ERR_DECODECONFIG
	}

	// Return the good config.
	return c, nil
}

func ConfigDump(c *Config) error {

	// Creates the file "eondump.json" for the config to be dumped into.
	file, createErr := os.Create(DUMP_NAME)

	// If the file could not be created.
	if createErr != nil {

		return ERR_FAILEDCREATEFILE
	}

	// Close the file when done.
	defer file.Close()

	// Encode the config into the file.
	encodeErr := json.NewEncoder(file).Encode(c)

	// If the config could not be encoded.
	if encodeErr != nil {

		return ERR_ENCODE
	}

	return nil
}
