package main

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
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

	client, err := ethclient.Dial("https://ethereum.publicnode.com")
	if err != nil {
		log.Fatal(err)
	}

	from := derived.Address
	destination := common.HexToAddress("0x0")
	amount := big.NewInt(1000000000000000000)

	tx, err := transaction.NewTxWithClient(context.Background(), client, from, destination, amount)
	if err != nil {
		log.Fatal(err)
	}

	signedTx, err := transaction.SignTx(tx, derived.PrivateKey)
	if err != nil {
		log.Fatal(err)
	}

	if err := client.SendTransaction(context.Background(), signedTx); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Signed and broadcasted transaction")
}
