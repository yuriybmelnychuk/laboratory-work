package blockchain

import (
	"errors"
	"fmt"
	"practice/pkg/block"
	"practice/pkg/signature"
	"practice/pkg/transaction"
)

type Blockchain struct {
	BlockHistory []*block.Block
	TxDatabase   []*transaction.Transaction
}

// Проверяем валидность
func (bl *Blockchain) ValidateBlock(block *block.Block) error {
	for _, tx := range block.Transactions {
		for _, op := range tx.Operations {
			data := op.Data()

			sellerSig := &signature.Signature{
				Data: op.SellerSignature,
			}

			err := sellerSig.Verify(op.SellerPublicKey, data)
			if err != nil {
				return errors.New("cannot verify signature by seller")
			}

			buyerSig := &signature.Signature{
				Data: op.BuyerSignature,
			}

			err = buyerSig.Verify(op.BuyerPublicKey, data)
			if err != nil {
				return errors.New("cannot verify signature by buyer")
			}

			notarySig := &signature.Signature{
				Data: op.NotarySignature,
			}

			err = notarySig.Verify(op.NotaryPublicKey, data)
			if err != nil {
				return errors.New("cannot verify signature by notary")
			}
		}
	}

	bl.BlockHistory = append(bl.BlockHistory, block)
	bl.TxDatabase = append(bl.TxDatabase, block.Transactions...)

	return nil
}

func InitBlockchain() *Blockchain {
	return &Blockchain{}
}
func (bl *Blockchain) PrintBlockchain() {
	fmt.Println("Blockchain blocks:")
	for i, block := range bl.BlockHistory {
		fmt.Printf("#%d\n", i)
		fmt.Printf("Block Hash: %x\n", block.Hash)
		fmt.Printf("Transactions: %v\n", len(block.Transactions))
	}

	fmt.Println()
	fmt.Println("Blockchain transactions:")

	for i, tx := range bl.TxDatabase {
		fmt.Printf("#%d\n", i)
		tx.Print()
	}
}
