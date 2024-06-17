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
