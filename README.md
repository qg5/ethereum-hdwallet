# Ethereum HDWallet
A Go library for hierarchical deterministic wallets (HD wallets) compatible with Ethereum.

[![Go Reference](https://pkg.go.dev/badge/github.com/qg5/ethereum-hdwallet.svg)](https://pkg.go.dev/github.com/qg5/ethereum-hdwallet)

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
	"github.com/qg5/ethereum-hdwallet/ethd"
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

	fmt.Printf("Private key: 0x%x\n", crypto.FromECDSA(derived.PrivateKey)) // 0xa1abf97524bf5ed4add70cde3a7d131eec08b9ab4f7cc1e746edce7f078132c5
	fmt.Printf("Public key: 0x%x\n", crypto.CompressPubkey(&derived.PublicKey)) // 0x02717c2f423ea93de87d1589dc4aeb760c30b368bd5e8b05fc40145f5ada78b2a2
	fmt.Println("Address:", derived.Address) // 0x773d3ACc0322A90924c53536a44eF38D50CfC9D1
}
```

- You can also browse multiple examples located in the [examples folder](https://github.com/qg5/ethereum-hdwallet/tree/main/examples)

## Resources

[Ian Coleman BIP39](https://iancoleman.io/bip39/)
