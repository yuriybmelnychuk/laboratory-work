package transaction

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"practice/pkg/hash"
	"practice/pkg/operation"
)

type Transaction struct {
	Id         [20]byte
	Nonce      uint64
	Operations []*operation.Operation
}

// Создаем транзакцию
func CreateTransaction(nonce uint64, operations []*operation.Operation) *Transaction {
	return &Transaction{
		Nonce:      nonce,
		Operations: operations,
	}
}

func (tx *Transaction) CalculateId() {
	transactionBuffer := bytes.Buffer{}

	nonceBytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(nonceBytes, tx.Nonce)
	transactionBuffer.Write(nonceBytes)

	for _, o := range tx.Operations {
		typeBytes := make([]byte, 8)
		binary.LittleEndian.PutUint64(typeBytes, o.Type)

		transactionBuffer.Write(typeBytes)
		transactionBuffer.WriteString(o.Address)
		transactionBuffer.Write(o.SellerSignature)
		transactionBuffer.Write(o.BuyerSignature)
		transactionBuffer.Write(o.NotarySignature)
	}

	transactionBytes := transactionBuffer.Bytes()
	transactionHash := hash.Sha1Sum(transactionBytes)

	tx.Id = transactionHash
}

func (tx *Transaction) Print() {
	fmt.Printf("\nTransaction Id: %x\n", tx.Id)
	fmt.Printf("Transaction Nonce: %d\n\n", tx.Nonce)
	fmt.Printf("Transaction Operations:\n")

	for i, operation := range tx.Operations {
		fmt.Printf("#%d\n", i)
		operation.Print()
		fmt.Println()
	}
}
