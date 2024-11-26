package demoencoding

import (
	"encoding/base64"
	"fmt"
)

func Base64Demo(s string) {
	encStr := base64.URLEncoding.EncodeToString([]byte(s))
	fmt.Println("encoded string: ", encStr)

	decByte, err := base64.URLEncoding.DecodeString(encStr)
	if err != nil {
		fmt.Println("error decoding: ", err)
	}
	fmt.Println("decoded string: ", string(decByte))
}
