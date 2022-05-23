package operation

import (
	"bytes"
	"crypto/rsa"
	"encoding/binary"
	"fmt"
)

type OperationType = uint64

const (
	CREATE OperationType = iota
	CHANGE_OWNER
)

type Operation struct {
	Type            OperationType
	Address         string
	SellerSignature []byte
	BuyerSignature  []byte
	NotarySignature []byte
	SellerPublicKey *rsa.PublicKey
	BuyerPublicKey  *rsa.PublicKey
	NotaryPublicKey *rsa.PublicKey
}

// Выводим все поля операций
func (o *Operation) Print() {
	fmt.Printf("Operation Type: %d\n", o.Type)
	fmt.Printf("Operation Address: %s\n", o.Address)
}
func (o *Operation) Data() []byte {
	dataBuffer := bytes.Buffer{}

	typeBytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(typeBytes, o.Type)
	dataBuffer.Write(typeBytes)

	dataBuffer.WriteString(o.Address)

	return dataBuffer.Bytes()
}
