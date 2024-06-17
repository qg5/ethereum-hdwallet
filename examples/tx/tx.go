package main

import (
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/qg5/ethereum-hdwallet/ethd"
	"github.com/qg5/ethereum-hdwallet/transaction"
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

	nonce := uint64(0)
	destination := common.HexToAddress("0x0")
	amount := big.NewInt(1000000000000000000)
	gasLimit := uint64(21000)
	gasPrice := big.NewInt(21000000000)

	tx := transaction.NewTx(nonce, &destination, amount, gasLimit, gasPrice, []byte{})
	signedTx, err := transaction.SignTx(tx, derived.PrivateKey)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Signed tx hex:", signedTx)
}
