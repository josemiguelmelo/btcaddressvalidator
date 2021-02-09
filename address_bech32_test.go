package btcaddressvalidator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckBech32AddressP2WSH(t *testing.T) {
	assert := assert.New(t)

	address := "bc1qar0srrr7xfkvy5l643lydnw9re59gtzzwf5mdq"
	addressType, err := CheckBech32Address(address)
	assert.Nil(err)

	assert.Equal(true, addressType.IsBech32)
	assert.Equal("mainnet", addressType.Network)
	assert.Equal("p2wpkh", addressType.Type)
}

func TestCheckBech32AddressP2WPKH(t *testing.T) {
	assert := assert.New(t)

	address := "bc1qc7slrfxkknqcq2jevvvkdgvrt8080852dfjewde450xdlk4ugp7szw5tk9"
	addressType, err := CheckBech32Address(address)
	assert.Nil(err)

	assert.Equal(true, addressType.IsBech32)
	assert.Equal("mainnet", addressType.Network)
	assert.Equal("p2wsh", addressType.Type)
}

func TestCheckBech32AddressP2WPKHTestnet(t *testing.T) {
	assert := assert.New(t)

	address := "tb1qc7slrfxkknqcq2jevvvkdgvrt8080852dfjewde450xdlk4ugp7s4xzyv2"
	addressType, err := CheckBech32Address(address)
	assert.Nil(err)

	assert.Equal(true, addressType.IsBech32)
	assert.Equal("testnet", addressType.Network)
	assert.Equal("p2wsh", addressType.Type)
}

func TestCheckBech32AddressWithInvalidNetwork(t *testing.T) {
	assert := assert.New(t)

	address := "tc1qw508d6qejxtdg4y5r3zarvary0c5xw7kg3g4ty"
	addressType, err := CheckBech32Address(address)
	// assert equality
	assert.Equal("Invalid network prefix", err.Error())
	assert.Nil(addressType)
}

func TestCheckBech32AddressWithInvalidChecksum(t *testing.T) {
	assert := assert.New(t)

	address := "bc1qw508d6qejxtdg4y5r3zarvary0c5xw7kv8f3t5"
	addressType, err := CheckBech32Address(address)
	// assert equality
	assert.Contains(err.Error(), "checksum failed")
	assert.Nil(addressType)
}

func TestCheckBech32AddressWithInvalidVersion(t *testing.T) {
	assert := assert.New(t)

	address := "BC13W508D6QEJXTDG4Y5R3ZARVARY0C5XW7KN40WF2"
	addressType, err := CheckBech32Address(address)
	// assert equality
	assert.Equal("Invalid witness version", err.Error())
	assert.Nil(addressType)
}

func TestBIP0141Type(t *testing.T) {
	assert := assert.New(t)
	size32 := [32]byte{}
	size20 := [20]byte{}
	invalidSize := [10]byte{}
	networkType, err := bip0141Type(size32[:])
	assert.Nil(err)
	assert.Equal("p2wsh", networkType)

	networkType, err = bip0141Type(size20[:])
	assert.Nil(err)
	assert.Equal("p2wpkh", networkType)

	networkType, err = bip0141Type(invalidSize[:])
	assert.Equal("Invalid bip0141 length: 10", err.Error())
	assert.Equal("", networkType)
}
