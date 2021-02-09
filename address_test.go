package btcaddressvalidator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckBtcAddressWithBase58(t *testing.T) {
	assert := assert.New(t)

	address := "1BvBMSEYstWetqTFn5Au4m4GFg7xJaNVN2"
	addressType, err := CheckBtcAddress(address)
	assert.Nil(err)

	assert.Equal(false, addressType.IsBech32)
	assert.Equal("mainnet", addressType.Network)
	assert.Equal("p2pkh", addressType.Type)
}

func TestCheckBtcAddressWithBech32(t *testing.T) {
	assert := assert.New(t)

	address := "bc1qar0srrr7xfkvy5l643lydnw9re59gtzzwf5mdq"
	addressType, err := CheckBtcAddress(address)
	assert.Nil(err)

	assert.Equal(true, addressType.IsBech32)
	assert.Equal("mainnet", addressType.Network)
	assert.Equal("p2wpkh", addressType.Type)
}
