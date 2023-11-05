package main

import (
	"crypto/aes"
	"fmt"
)

func main() {
	text := "100 RED BALLOONS"

	key := []byte("WWWWXXXXYYYYZZZZ")
	aesBlock, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	cipherKey := make([]byte, 16)
	aesBlock.Encrypt(cipherKey, []byte(text))
	fmt.Printf("%x\n", cipherKey)

	decrypted := []byte("WWWWXXXXYYYYZZZZ")
	aesBlock.Decrypt(decrypted, cipherKey)
	fmt.Printf("%s", decrypted)
}
