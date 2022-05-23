package signature

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"practice/pkg/hash"
)

type Signature struct {
	Data []byte
}

// Подписываем данные
func (s *Signature) SignData(privateKey *rsa.PrivateKey, data []byte) error {
	sum := hash.Sha1Sum(data)
	data, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA1, sum[:])
	if err != nil {
		return err
	}

	s.Data = data
	return nil
}

// Проверяем подписи
func (s *Signature) Verify(publicKey *rsa.PublicKey, data []byte) error {
	sum := hash.Sha1Sum(data)
	return rsa.VerifyPKCS1v15(publicKey, crypto.SHA1, sum[:], s.Data)
}
