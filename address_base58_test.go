package bitcoinaddressvalidator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckBase58AddressP2PKH(t *testing.T) {
	assert := assert.New(t)

	address := "1BvBMSEYstWetqTFn5Au4m4GFg7xJaNVN2"
	addressType, err := CheckBase58Address(address)
	assert.Nil(err)

	assert.Equal(false, addressType.IsBech32)
	assert.Equal("mainnet", addressType.Network)
	assert.Equal("p2pkh", addressType.Type)
}

func TestCheckBase58AddressP2SH(t *testing.T) {
	assert := assert.New(t)

	address := "3J98t1WpEZ73CNmQviecrnyiWrnqRhWNLy"
	addressType, err := CheckBase58Address(address)
	assert.Nil(err)

	assert.Equal(false, addressType.IsBech32)
	assert.Equal("mainnet", addressType.Network)
	assert.Equal("p2sh", addressType.Type)
}

func TestCheckBase58AddressWithInvalidLength(t *testing.T) {
	assert := assert.New(t)

	address := "3J98t1WpEZ73CNmQviecrnyiWrnqRhWNLyasd"
	addressType, err := CheckBase58Address(address)
	assert.Nil(addressType)

	assert.Equal("Invalid address length", err.Error())
}

func TestCheckBase58AddressWithInvalidChecksum(t *testing.T) {
	assert := assert.New(t)

	address := "9J98t1WpEZ73CNmQviecrnyiWrnqRhWNLy"
	addressType, err := CheckBase58Address(address)
	assert.Nil(addressType)

	assert.Equal("Failed to verify checksum", err.Error())
}
