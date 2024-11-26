package demoasymmetricencryption

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
)

func AsymmetricEncryptionDemo(plainText string, key *rsa.PublicKey) ([]byte, error) {
	label := []byte("OAEP Encrypted")

	cipherText, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, key, []byte(plainText), label)
	if err != nil {
		return nil, err
	}

	return cipherText, nil
}

func AsymmetricDecryptionDemo(cipherText []byte, key *rsa.PrivateKey) ([]byte, error) {
	label := []byte("OAEP Encrypted")

	plainText, err := rsa.DecryptOAEP(sha256.New(), rand.Reader, key, cipherText, label)
	if err != nil {
		return nil, err
	}

	return plainText, nil
}
