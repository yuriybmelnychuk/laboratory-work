package main

import (
	"fmt"
	"practice/pkg/account"
	"practice/pkg/block"
	"practice/pkg/blockchain"
	"practice/pkg/hash"
	"practice/pkg/operation"
	"practice/pkg/signature"
	"practice/pkg/transaction"
)

func main() {
	operations := []*operation.Operation{
		{ // Данные первого объекта недвижимости
			Type:    operation.CHANGE_OWNER,
			Address: "Main st. 1, apt. 1",
		},
		{ // Данные второго объекта недвижимости
			Type:    operation.CHANGE_OWNER,
			Address: "Main st. 1, apt. 2",
		},
	}
	// Получаем хеш
	hash := hash.Sha1Sum([]byte(""))
	fmt.Printf("\nHash: %x\n", hash)
	// Создаем аккаунт продавца
	sellerAccount, err := account.Generate()
	if err != nil {
		fmt.Println("Cannot generate seller account")
	}
	// Создаем аккаунт покупателя
	buyerAccount, err := account.Generate()
	if err != nil {
		fmt.Println("Cannot generate buyer account")
	}
	// Создаем аккаунт нотариуса
	notaryAccount, err := account.Generate()
	if err != nil {
		fmt.Println("Cannot generate notary account")
	}

	sellerAccount.Print()
	buyerAccount.Print()
	notaryAccount.Print()

	// Подписываем операции 3-мя участниками
	for _, operation := range operations {
		data := operation.Data()

		sellerSig := signature.Signature{}
		err := sellerSig.SignData(sellerAccount.Wallet[0].PrivateKey, data)
		if err != nil {
			fmt.Println("Cannot sign data by seller")
			continue
		}

		operation.SellerSignature = sellerSig.Data

		buyerSig := signature.Signature{}
		err = buyerSig.SignData(buyerAccount.Wallet[0].PrivateKey, data)
		if err != nil {
			fmt.Println("Cannot sign data by buyer")
			continue
		}

		operation.BuyerSignature = buyerSig.Data

		notarySig := signature.Signature{}
		err = notarySig.SignData(notaryAccount.Wallet[0].PrivateKey, data)
		if err != nil {
			fmt.Println("Cannot sign data by notary")
			continue
		}

		operation.NotarySignature = notarySig.Data

		fmt.Printf("\n\nSellerSignature: \n%v\n\nBuyerSignature: \n%v\n\nNotarySignature: \n%v\n", operation.SellerSignature, operation.BuyerSignature, operation.NotarySignature)

		operation.SellerPublicKey = &sellerAccount.Wallet[0].PrivateKey.PublicKey
		operation.BuyerPublicKey = &buyerAccount.Wallet[0].PrivateKey.PublicKey
		operation.NotaryPublicKey = &notaryAccount.Wallet[0].PrivateKey.PublicKey

		// fmt.Printf("\n\nSellerPublicKey!!!: \n%v\n\nBuyerPublicKey!!!: \n%v\n\nNotaryPublicKey!!!: \n%v\n", operation.SellerPublicKey, operation.BuyerPublicKey, operation.NotaryPublicKey)

	}

	// Проверяем валидность подписей в операциях
	for _, operation := range operations {
		data := operation.Data()

		sellerSig := &signature.Signature{
			Data: operation.SellerSignature,
		}

		err := sellerSig.Verify(&sellerAccount.Wallet[0].PrivateKey.PublicKey, data)
		if err != nil {
			fmt.Println("Cannot verify signature by seller")
			continue
		}

		buyerSig := &signature.Signature{
			Data: operation.BuyerSignature,
		}

		err = buyerSig.Verify(&buyerAccount.Wallet[0].PrivateKey.PublicKey, data)
		if err != nil {
			fmt.Println("Cannot verify signature by buyer")
			continue
		}

		notarySig := &signature.Signature{
			Data: operation.NotarySignature,
		}

		err = notarySig.Verify(&notaryAccount.Wallet[0].PrivateKey.PublicKey, data)
		if err != nil {
			fmt.Println("Cannot verify signature by notary")
			continue
		}
	}

	tx := transaction.CreateTransaction(123, operations)
	tx.CalculateId()
	tx.Print()

	block := &block.Block{
		Transactions: []*transaction.Transaction{
			tx,
		},
	}

	block.CalculateHash()

	bc := blockchain.InitBlockchain()
	err = bc.ValidateBlock(block)
	if err != nil {
		fmt.Println("Failed to validate block:", err)
	}

	bc.PrintBlockchain()
}
