package btcaddressvalidator

import (
	"bytes"
	"crypto/sha256"
	"errors"
	"fmt"

	"github.com/btcsuite/btcutil/base58"
	"github.com/btcsuite/btcutil/bech32"
)

// BitcoinAddressType contains information about a bitcoin address
type BitcoinAddressType struct {
	Address  string
	IsBech32 bool
	Type     string
	Network  string
}

type addressType struct {
	Type    string
	Network string
}

var (
	addressTypesBasedOnNetworkIDByte = map[byte]addressType{
		0x00: {
			Type:    "p2pkh",
			Network: "mainnet",
		},
		0x6f: {
			Type:    "p2pkh",
			Network: "testnet",
		},
		0x05: {
			Type:    "p2sh",
			Network: "mainnet",
		},
		0xc4: {
			Type:    "p2sh",
			Network: "testnet",
		},
	}

	bech32PrefixNetwork = map[string]string{
		"bc":   "mainnet",
		"tb":   "testnet",
		"bcrt": "regtest",
	}
)

// CheckBtcAddress validates a bitcoin address and returns the address information if valid.
func CheckBtcAddress(address string) (*BitcoinAddressType, error) {
	prefix := string(address[:2])

	if prefix == "bc" || prefix == "tb" {
		return CheckBech32Address(address)
	}

	return CheckBase58Address(address)
}

// CheckBase58Address validates an address against base58 format and returns the address information if valid.
func CheckBase58Address(address string) (*BitcoinAddressType, error) {
	decodedAddress := base58.Decode(address)
	if len(decodedAddress) != 25 {
		return nil, errors.New("Invalid address length")
	}

	networkIDByte := decodedAddress[0]
	body := decodedAddress[0:21]
	checksum := decodedAddress[21:25]

	bodySha256 := sha256.Sum256(body)
	expectedChecksum := sha256.Sum256(bodySha256[:])

	if !bytes.Equal(checksum, expectedChecksum[:4]) {
		return nil, errors.New("Failed to verify checksum")
	}

	if addressTypeAndNetwork, exists := addressTypesBasedOnNetworkIDByte[networkIDByte]; exists {
		return &BitcoinAddressType{
			Address:  address,
			IsBech32: false,
			Type:     addressTypeAndNetwork.Type,
			Network:  addressTypeAndNetwork.Network,
		}, nil
	}

	return nil, errors.New("Network type not found")
}

// CheckBech32Address validates an address against bech32 format and returns the address information if valid.
func CheckBech32Address(address string) (*BitcoinAddressType, error) {
	networkPrefix, decoded, err := bech32.Decode(address)
	if err != nil {
		return nil, err
	}

	// https://en.bitcoin.it/wiki/BIP_0173
	// https://github.com/bitcoin/bips/blob/master/bip-0141.mediawiki#p2wpkh
	witnessVersion := decoded[0]
	if int(witnessVersion) < 0 || int(witnessVersion) > 16 {
		return nil, errors.New("Invalid witness version")
	}

	body := decoded[1:]
	body, err = bech32.ConvertBits(body, 5, 8, false)
	if err != nil {
		return nil, err
	}
	segwitType, err := bip0141Type(body)
	if err != nil {
		return nil, err
	}

	if network, exists := bech32PrefixNetwork[networkPrefix]; exists {
		return &BitcoinAddressType{
			Address:  address,
			IsBech32: true,
			Type:     segwitType,
			Network:  network,
		}, nil
	}

	return nil, errors.New("Invalid network prefix")
}

func bip0141Type(body []byte) (string, error) {
	if len(body) == 20 {
		return "p2wpkh", nil
	}
	if len(body) == 32 {
		return "p2wsh", nil
	}
	return "", fmt.Errorf("Invalid bip0141 length: %d", len(body))
}
