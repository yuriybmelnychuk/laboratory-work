package transaction

import "fmt"

type OperationType = uint64

const (
	CREATE OperationType = iota
	CHANGE_OWNER
)

type Operation struct {
	Type            OperationType
	Address         string
	SellerSignature string
	BuyerSignature  string
	NotarySignature string
}

func (o *Operation) Print() {
	fmt.Printf("Operation Type: %d\n", o.Type)
	fmt.Printf("Operation Address: %s\n", o.Address)
	fmt.Printf("Operation SellerSignature: %s\n", o.SellerSignature)
	fmt.Printf("Operation BuyerSignature: %s\n", o.BuyerSignature)
	fmt.Printf("Operation NotarySignature: %s\n", o.NotarySignature)
}
