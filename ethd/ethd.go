package ethd

import (
	"crypto/ecdsa"
	"errors"
	"fmt"

	hd "github.com/btcsuite/btcd/btcutil/hdkeychain"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/tyler-smith/go-bip39"
)

type Wallet struct {
	masterKey *hd.ExtendedKey
}

type Derived struct {
	ExtendedKey *hd.ExtendedKey
	Address     common.Address
	PrivateKey  *ecdsa.PrivateKey
	PublicKey   ecdsa.PublicKey
}

// DefaultDerivationPath is the default derivation path used
var DefaultDerivationPath = "m/44'/60'/0'/0/%d"

// New creates a new Wallet instance from a mnemonic and optional password
func New(mnemonic, password string) (*Wallet, error) {
	seed, err := bip39.NewSeedWithErrorChecking(mnemonic, password)
	if err != nil {
		return nil, err
	}

	masterKey, err := hd.NewMaster(seed, &chaincfg.MainNetParams)
	if err != nil {
		return nil, err
	}

	return &Wallet{masterKey: masterKey}, nil
}

// NewFromSeed creates a new Wallet instance from a seed
func NewFromSeed(seed []byte) (*Wallet, error) {
	masterKey, err := hd.NewMaster(seed, &chaincfg.MainNetParams)
	if err != nil {
		return nil, err
	}

	return &Wallet{masterKey: masterKey}, nil
}

// ParseDerivationPath parses the provided derivation path
func ParseDerivationPath(path string) (accounts.DerivationPath, error) {
	parsed, err := accounts.ParseDerivationPath(path)
	if err != nil {
		return nil, err
	}

	return parsed, nil
}

// CreateDerivationPath creates a derivation path from an index using DefaultDerivationPath
func CreateDerivationPath(index int) (accounts.DerivationPath, error) {
	path := fmt.Sprintf(DefaultDerivationPath, index)
	return ParseDerivationPath(path)
}

// Derive derives an extended key from the masterKey using the provided derivation path.
// It also derives the corresponding private key, public key, and address at that path
func (w *Wallet) Derive(path accounts.DerivationPath) (*Derived, error) {
	key, err := deriveKey(w.masterKey, path)
	if err != nil {
		return nil, err
	}

	privateKey, err := getPrivateKey(key)
	if err != nil {
		return nil, err
	}

	publicKey := privateKey.PublicKey

	return &Derived{
		ExtendedKey: key,
		Address:     crypto.PubkeyToAddress(publicKey),
		PrivateKey:  privateKey,
		PublicKey:   publicKey,
	}, nil
}

// Account returns an ethereum account
func (d *Derived) Account() accounts.Account {
	return accounts.Account{Address: d.Address}
}

// deriveKey derives an extended key from the provided key using the provided derivation
// path
func deriveKey(key *hd.ExtendedKey, path accounts.DerivationPath) (*hd.ExtendedKey, error) {
	var err error

	for _, n := range path {
		key, err = key.Derive(n)
		if err == hd.ErrInvalidChild {
			return nil, errors.New("the extended key at this index is invalid, increase the index and try again")
		}

		if err != nil {
			return nil, err
		}
	}

	return key, nil
}

// getPrivateKey gets the ecdsa private key from the extended key
func getPrivateKey(key *hd.ExtendedKey) (*ecdsa.PrivateKey, error) {
	privKey, err := key.ECPrivKey()
	if err != nil {
		return nil, err
	}

	return privKey.ToECDSA(), nil
}
