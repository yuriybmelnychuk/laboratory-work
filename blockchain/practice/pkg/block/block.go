package block

import (
	"bytes"
	"practice/pkg/hash"
	"practice/pkg/transaction"
)

type Block struct {
	Hash         []byte
	Transactions []*transaction.Transaction
}

// Создаем новый блок
func New(transactions []*transaction.Transaction) *Block {
	return &Block{
		Transactions: transactions,
	}
}

// Добавляем транзакции
func (bl *Block) AddTransaction(transaction *transaction.Transaction) {
	bl.Transactions = append(bl.Transactions, transaction)
}

func (bl *Block) CalculateHash() {
	buffer := bytes.NewBuffer(nil)
	for _, tx := range bl.Transactions {
		tx.CalculateId()
		buffer.Write(tx.Id[:])
	}

	digest := hash.Sha1Sum(buffer.Bytes())
	bl.Hash = digest[:]
}
