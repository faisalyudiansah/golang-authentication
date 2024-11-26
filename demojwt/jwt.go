package demojwt

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func CreateAndSign(data string, secretKey string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"my-data": data,
		"iss":     "the-issuer",
		"exp":     time.Now().Add(1 * time.Hour).Unix(),
		// and other claims, alternatively, you may want to explore on how to create custome claims
	})

	signed, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return signed, nil
}

func ParseAndVerify(signed string, secretKey string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(signed, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	}, jwt.WithIssuer("the-issuer"),
		jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Name}),
		jwt.WithExpirationRequired(),
	)
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		return claims, nil
	} else {
		return nil, fmt.Errorf("unknown claims")
	}
}
