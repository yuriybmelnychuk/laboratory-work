package account

import (
	"fmt"
	"practice/pkg/keypair"
)

type Account struct {
	Wallet []*keypair.KeyPair
}

// Создаем ключевую пару
func Generate() (*Account, error) {
	key, err := keypair.Generate()
	if err != nil {
		return nil, err
	}

	return &Account{
		Wallet: []*keypair.KeyPair{key},
	}, nil
}

func (a *Account) Print() {
	for i, keyPair := range a.Wallet {
		fmt.Printf("\nKeyPair #%d", i)
		fmt.Println("\nPrivate Key:\n", keyPair.PrivateKey)
		fmt.Println("\nPublic Key:\n", keyPair.PublicKey)
	}
}
