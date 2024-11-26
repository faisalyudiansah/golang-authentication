package demosymmetrichash

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

func SymmetricHashDemo(s string, key string) string {
	mac := hmac.New(sha256.New, []byte(key))
	mac.Write([]byte(s))
	res := mac.Sum(nil)

	// cannot use res directly, as it will generate characters which cannot be outputted to the screen
	hmacHex := hex.EncodeToString(res)

	return hmacHex
}

func SymmetricHashCheckDemo(s string, messageMac string, key string) (bool, error) {
	macB, err := hex.DecodeString(messageMac)
	if err != nil {
		return false, err
	}

	macA := hmac.New(sha256.New, []byte(key))
	macA.Write([]byte(s))
	resMacA := macA.Sum(nil)

	return hmac.Equal(macB, resMacA), nil
}
