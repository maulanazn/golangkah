package cryptography

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
)

const MySecret string = "abc&1*~#^2^#s0^=)^^7%b34"

func Encode(value []byte) string {
	return base64.StdEncoding.EncodeToString(value)
}

func Decode(s string) []byte {
	data, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return data
}

func EncryptV2(text, MySecret string) (string, error) {
	var iv = make([]byte, 16)

	block, err := aes.NewCipher([]byte(MySecret))
	if err != nil {
		return "", err
	}
	plainText := []byte(text)
	cfb := cipher.NewCFBEncrypter(block, iv)
	cipherText := make([]byte, len(plainText))
	cfb.XORKeyStream(cipherText, plainText)
	return Encode(cipherText), nil
}

func DecryptV2(text, MySecret string) (string, error) {
	var iv = make([]byte, 16)

	block, err := aes.NewCipher([]byte(MySecret))
	if err != nil {
		return "", err
	}
	cipherText := Decode(text)
	cfb := cipher.NewCFBDecrypter(block, iv)
	plainText := make([]byte, len(cipherText))
	cfb.XORKeyStream(plainText, cipherText)
	return string(plainText), nil
}

func RunCryptV2() {
	hash := sha256.New()
	hash.Write([]byte(MySecret))
	key := hash.Sum(nil)

	result, _ := EncryptV2("1923-10-192384-6", string(key))
	fmt.Println(result)

	resultDec, _ := DecryptV2("o7eFH9eIak7cyAODQHDNng==", string(key))
	fmt.Println(resultDec)
}
