[![Build Status](https://travis-ci.com/josemiguelmelo/btcaddressvalidator.svg?branch=master)](https://travis-ci.com/josemiguelmelo/btcaddressvalidator)
[![Codacy Badge](https://app.codacy.com/project/badge/Grade/1d316ba621044aa9873837475d3cfe10)](https://www.codacy.com/gh/josemiguelmelo/btcaddressvalidator/dashboard?utm_source=github.com&utm_medium=referral&utm_content=josemiguelmelo/btcaddressvalidator&utm_campaign=Badge_Grade)
[![Codacy Badge](https://app.codacy.com/project/badge/Coverage/1d316ba621044aa9873837475d3cfe10)](https://www.codacy.com/gh/josemiguelmelo/btcaddressvalidator/dashboard?utm_source=github.com&utm_medium=referral&utm_content=josemiguelmelo/btcaddressvalidator&utm_campaign=Badge_Coverage)

# bitcoin-address-validator

Golang implementation of a Bitcoin address validator.

## Install

To install as a library:

    go get github.com/josemiguelmelo/btcaddressvalidator

## Usage

The btc address validator return as result a `BitcoinAddressType` if valid and an error in case of invalid address. 

The type `BitcoindAddressType` contains the following information about the bitcoin address address:

-   Address - Bitcoin address string
-   IsBech32 - Address has Bech32 format
-   Type - Address formats
    -   For base58 addresses:
        -   p2pkh - Begins with the number 1. Example: 1BvBMSEYstWetqTFn5Au4m4GFg7xJaNVN2
        -   p2sh - Begins with the number 3. Example: 3J98t1WpEZ73CNmQviecrnyiWrnqRhWNLy
    -   For Bech32 addresses - begins with `bc1`:
        -   p2wpkh
        -   p2wsh
-   Network - The bitcoin network where the address belongs to 
    -   mainnet
    -   testnet
    -   regtest

```go
type BitcoinAddressType struct {
	Address  string
	IsBech32 bool
	Type     string
	Network  string
}
```

### Base58 Address

To validate a Base58 btc address:

```go
address := "1BvBMSEYstWetqTFn5Au4m4GFg7xJaNVN2"
addressType, err := btcaddressvalidator.CheckBase58Address(address)
```

### Bech32 Address

To validate a Bech32 btc address:

```go
address := "bc1qar0srrr7xfkvy5l643lydnw9re59gtzzwf5mdq"
addressType, err := CheckBech32Address(address)
```

### Btc Address (independent of the type)

If you don't want to check for a specific type of btc address, you can use the `CheckBtcAddress`. This method checks for both bech32 and base58 addresses.

```go
base58Address := "1BvBMSEYstWetqTFn5Au4m4GFg7xJaNVN2"
base58AddressType, err := CheckBtcAddress(base58Address)

bech32Address := "bc1qar0srrr7xfkvy5l643lydnw9re59gtzzwf5mdq"
bech32AddressType, err := CheckBtcAddress(bech32Address)
```

## License

The MIT License (MIT). Please see [LICENSE file](https://github.com/josemiguelmelo/btcaddressvalidator/blob/master/LICENSE) for more information.
