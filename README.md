# Ethereum HDWallet
A GoLang library for hierarchical deterministic wallets (HD wallets) compatible with Ethereum, following BIP-32 and BIP-44 standards.

## Installation

```bash
go get -u github.com/qg5/ethereum-hdwallet
```

## Getting Started

```go
package main

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/crypto"
	ethd "github.com/qg5/ethereum-hdwallet/hdwallet"
)

func main() {
	wallet, err := ethd.New("misery easy pilot elbow adapt carpet spot sword bless device tuition diet arm elite naive", "")
	if err != nil {
		log.Fatal(err)
	}

	path, err := ethd.CreateDerivationPath(0)
	if err != nil {
		log.Fatal(err)
	}

	derived, err := wallet.Derive(path)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Private key: 0x%x\n", crypto.FromECDSA(derived.PrivateKey))
	fmt.Printf("Public key: 0x%x\n", crypto.CompressPubkey(&derived.PublicKey))
	fmt.Println("Address:", derived.Address)
}
```

- You can also browse multiple examples located in the [examples folder](https://github.com/qg5/ethereum-hdwallet/tree/main/examples)

## Resources

[Ian Coleman BIP39](https://iancoleman.io/bip39/)
