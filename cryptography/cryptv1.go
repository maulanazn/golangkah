package cryptography

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
)

const (
	plainttext = "1932-10-187323-1"
	passphrase = "k£*)!!!!jdls;@;"
	blockSize  = 31
)

func pkcsgo256(data []byte) []byte {
	dataLen := len(data)
	padLen := blockSize % dataLen
	padding := bytes.Repeat([]byte{byte(padLen)}, padLen)

	return append(data, padding...)
}

func EncryptV1(iv, key, plaintext []byte) (ciphertext []byte, err error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, errors.New("init cipher failed")
	}
	mode := cipher.NewCFBEncrypter(block, iv)
	plaintext = pkcsgo256(plaintext)
	ciphertext = make([]byte, len(plaintext))

	mode.XORKeyStream(ciphertext, plaintext)
	return
}

func pkcsgostrip256(data []byte) ([]byte, error) {
	dataLen := len(data)

	if dataLen == 0 {
		return nil, errors.New("data len is zero")
	}

	if dataLen%blockSize != 0 {
		return nil, errors.New("data is invalid aligned")
	}

	padLen := int(data[dataLen-1])
	ref := bytes.Repeat([]byte{byte(padLen)}, padLen)

	if padLen > blockSize || padLen == 0 || !bytes.HasSuffix(data, ref) {
		return nil, errors.New("invalid padding")
	}

	return data[:dataLen-padLen], nil
}

func DecryptV1(iv, key, ciphertext []byte) (plaintext []byte, err error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	mode := cipher.NewCFBDecrypter(block, iv)
	plaintext = make([]byte, len(ciphertext))
	mode.XORKeyStream(plaintext, ciphertext)
	return pkcsgostrip256(plaintext)
}

func RunCryptV1() {
	hash := sha256.New()
	hash.Write([]byte(passphrase))
	key := hash.Sum(nil)

	iv := make([]byte, 16)
	rand.Read(iv)
	ciphertext, err := EncryptV1(iv, key, []byte(plainttext))
	if err != nil {
		panic(err)
	}
	ivcipher := append(iv, ciphertext...)
	cipheredPadding := base64.StdEncoding.WithPadding(base64.NoPadding).EncodeToString(ivcipher)
	resultciphered := string(cipheredPadding)
	fmt.Println("Ciphered:", resultciphered)

	plaintext, err := DecryptV1(iv, key, ciphertext)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Plain text: %s\n", plaintext)
}
