package transaction

import (
	"context"
	"crypto/ecdsa"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

// NewTx creates a new legacy transaction
func NewTx(nonce uint64, to *common.Address, amount *big.Int, gasLimit uint64, gasPrice *big.Int, data []byte) *types.Transaction {
	return types.NewTx(&types.LegacyTx{Nonce: nonce, GasPrice: gasPrice, Gas: gasLimit, To: to, Value: amount, Data: data})
}

// NewTxWithClient makes use of the ethereum client in order to build a new transaction.
// This method uses types.DynamicFeeTx and not types.LegacyTx
func NewTxWithClient(ctx context.Context, client *ethclient.Client, from common.Address, to common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	nonce, err := client.PendingNonceAt(ctx, from)
	if err != nil {
		return nil, err
	}

	gasLimit, err := client.EstimateGas(ctx, ethereum.CallMsg{To: &to, Value: amount, Data: data})
	if err != nil {
		return nil, err
	}

	gasPrice, err := client.SuggestGasPrice(ctx)
	if err != nil {
		return nil, err
	}

	gasTipCap, err := client.SuggestGasTipCap(ctx)
	if err != nil {
		return nil, err
	}

	chainID, err := client.NetworkID(ctx)
	if err != nil {
		return nil, err
	}

	tx := types.NewTx(&types.DynamicFeeTx{
		ChainID:    chainID,
		Nonce:      nonce,
		To:         &to,
		Value:      amount,
		Gas:        gasLimit,
		GasFeeCap:  gasPrice,
		GasTipCap:  gasTipCap,
		Data:       data,
		AccessList: nil,
	})

	return tx, nil
}

// SignTx signs a raw transaction
func SignTx(tx *types.Transaction, privateKey *ecdsa.PrivateKey) (*types.Transaction, error) {
	signer := types.LatestSignerForChainID(tx.ChainId())
	return types.SignTx(tx, signer, privateKey)
}
