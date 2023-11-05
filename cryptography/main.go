package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
)

func mdHashing(input string) string {
	byteInput := []byte(input)
	md5Hash := md5.Sum(byteInput)
	return hex.EncodeToString(md5Hash[:]) // by referring to it as a string
}

func EncryptV1(keyString string, stringToEncrypt string) string {
	key, _ := hex.DecodeString(keyString)
	plainvalue := []byte(stringToEncrypt)

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	ciphertext := make([]byte, aes.BlockSize+len(plainvalue))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plainvalue)

	return base64.StdEncoding.EncodeToString(ciphertext)
}

func EncryptV2(value []byte, key []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}

	gcmInstance, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	nonce := make([]byte, gcmInstance.NonceSize())
	cipheredtext := gcmInstance.Seal(nil, nonce, value, nil)

	return cipheredtext
}

func Decrypt(cipheredValue []byte, keyPhrase string) []byte {
	aesBlock, err := aes.NewCipher([]byte(mdHashing(keyPhrase)))
	if err != nil {
		panic(err)
	}

	gcmInstalce, err := cipher.NewGCM(aesBlock)
	if err != nil {
		panic(err)
	}

	nonceSize := gcmInstalce.NonceSize()
	nonce, cipheredText := cipheredValue[:nonceSize], cipheredValue[nonceSize:]
	originalValue, err := gcmInstalce.Open(nil, nonce, cipheredText, nil)
	if err != nil {
		panic(err)
	}

	return originalValue
}

func main() {
	acc_number := "1932-12-192832-2"

	key := []byte("SecretKeysSecretKeysSecretKeys12")
	keyStr := hex.EncodeToString(key)

	result := EncryptV1(keyStr, acc_number)
	fmt.Println(string(result))

	resultv2 := EncryptV2([]byte(acc_number), key)
	fmt.Println(string(resultv2))

	// resultDec := Decrypt([]byte(result), "SecretKey")
}
