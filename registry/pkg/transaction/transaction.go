package transaction

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
)

type Transaction struct {
	Id         [32]byte
	Nonce      uint64
	Operations []Operation
}

func CreateTransaction(nonce uint64, operations []Operation) Transaction {
	return Transaction{
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
		transactionBuffer.WriteString(o.SellerSignature)
		transactionBuffer.WriteString(o.BuyerSignature)
		transactionBuffer.WriteString(o.NotarySignature)
	}

	transactionBytes := transactionBuffer.Bytes()
	transactionHash := sha256.Sum256(transactionBytes)

	tx.Id = transactionHash
}

func (tx *Transaction) Print() {
	fmt.Printf("Transaction Id: %x\n", tx.Id)
	fmt.Printf("Transaction Nonce: %d\n", tx.Nonce)
	fmt.Printf("Transaction Operations:\n\n")

	for _, operation := range tx.Operations {
		operation.Print()
		fmt.Println()
	}
}
