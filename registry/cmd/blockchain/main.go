package main

import (
	"registry/pkg/transaction"
)

func main() {
	operations := []transaction.Operation{
		{
			Type:            transaction.CHANGE_OWNER,
			Address:         "Main st. 1, apt. 1", // Данные первого объекта недвижимости
			SellerSignature: "signature 1",        // Подпись продавца
			BuyerSignature:  "signature 2",        // Подпись покупателя
			NotarySignature: "signature 3",        // Подпись нотариуса
		},
		{
			Type:            transaction.CHANGE_OWNER,
			Address:         "Main st. 1, apt. 2", // Данные второго объекта недвижимости
			SellerSignature: "signature 1",        // Подпись продавца
			BuyerSignature:  "signature 2",        // Подпись покупателя
			NotarySignature: "signature 3",        // Подпись нотариуса
		},
	}

	tx := transaction.CreateTransaction(123, operations)
	tx.CalculateId()
	tx.Print()
}
