package main

import (
	"fmt"

	"golang-authentication/demojwt"
)

const secretkey = "12345678901234567890123456789012"

func main() {
	// demo encoding
	// -------------
	// demoencoding.Base64Demo("hello world abc")

	// demo symmetric hash
	// -------------------
	// hm := demosymmetrichash.SymmetricHashDemo("hello world", secretkey)
	// fmt.Println(hm)
	// res, err := demosymmetrichash.SymmetricHashCheckDemo("hello world", hm, secretkey)
	// if err != nil {
	// 	fmt.Println("error: ", err)
	// 	return
	// }
	// fmt.Println("message is the same: ", res)

	// demo asymmetric encryption
	// --------------------------
	// privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	// if err != nil {
	// 	fmt.Println("error: ", err)
	// 	return
	// }

	// cipherText, err := demoasymmetricencryption.AsymmetricEncryptionDemo("this is the secret text", &privateKey.PublicKey)
	// if err != nil {
	// 	fmt.Println("error: ", err)
	// 	return
	// }

	// plainText, err := demoasymmetricencryption.AsymmetricDecryptionDemo(cipherText, privateKey)
	// if err != nil {
	// 	fmt.Println("error: ", err)
	// 	return
	// }

	// fmt.Println("cipher text: ", string(cipherText))
	// fmt.Println("encoded cipher text: ", base64.StdEncoding.EncodeToString(cipherText))
	// fmt.Println("plain text: ", string(plainText))

	// demo symmetric encryption
	// -------------------------
	// cipherText, err := demosymmetricencryption.SymmetricEncryptionDemo("hello world", secretkey)
	// if err != nil {
	// 	fmt.Println("error: ", err)
	// 	return
	// }
	// fmt.Println("cipher text:", string(cipherText))

	// plainText, err := demosymmetricencryption.SymmetricDecryptionDemo(cipherText, secretkey)
	// if err != nil {
	// 	fmt.Println("error: ", err)
	// 	return
	// }

	// fmt.Println(plainText)

	// Demo hashing password
	// ---------------------
	// pwd := "demo password"
	// hash, err := demohashingpassword.HashPassword(pwd, 12)
	// if err != nil {
	// 	fmt.Println("error: ", err)
	// 	return
	// }

	// isCorrectPassword, err := demohashingpassword.CheckPassword(pwd, hash)
	// if err != nil {
	// 	fmt.Println("error: ", err)
	// 	return
	// }

	// fmt.Println("hashed password: ", string(hash))
	// fmt.Println("is correct password: ", isCorrectPassword)

	// Demo Creating-Signing and Verifying-Parse JWT
	tokenString, err := demojwt.CreateAndSign("this is the payload", secretkey)
	if err != nil {
		fmt.Println("error: ", err)
		return
	}

	fmt.Println("token string: ", tokenString)

	claims, err := demojwt.ParseAndVerify(tokenString, secretkey+"1")
	if err != nil {
		fmt.Println("error: ", err)
		return
	}

	fmt.Println("claims: ", claims)

	// demojwt2
	// jwtProvider := demojwt2.NewJWTProviderHS256("dino", secretkey)

	// tokenString, err := jwtProvider.CreateToken(2)
	// if err != nil {
	// 	fmt.Println("error: ", err)
	// 	return
	// }

	// fmt.Println("token string: ", tokenString)

	// claims, err := jwtProvider.VerifyToken(tokenString)

	// if err != nil {
	// 	fmt.Println("error: ", err)
	// 	return
	// }

	// fmt.Println("claims: ", claims)

}
