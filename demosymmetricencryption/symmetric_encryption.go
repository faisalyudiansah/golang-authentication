package demosymmetricencryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
)

func SymmetricEncryptionDemo(plainText string, key string) ([]byte, error) {
	aes, err := aes.NewCipher([]byte(key))
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(aes)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	_, err = rand.Read(nonce)
	if err != nil {
		return nil, err
	}

	cipherText := gcm.Seal(nonce, nonce, []byte(plainText), nil)

	return cipherText, nil
}

func SymmetricDecryptionDemo(cipherText []byte, key string) (string, error) {
	aes, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(aes)
	if err != nil {
		return "", err
	}

	sz := gcm.NonceSize()
	nonce, cipherText := cipherText[:sz], cipherText[sz:]

	plainText, err := gcm.Open(nil, []byte(nonce), []byte(cipherText), nil)
	if err != nil {
		return "", err
	}

	return string(plainText), nil
}
