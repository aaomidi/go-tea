package main

import (
	"encoding/hex"
	"fmt"
	x "github.com/aaomidi/tea-go/tea"
	"golang.org/x/crypto/tea"
)

func main() {
	test1()
	test2()
}
func test1() {

	data := [8]byte{0x01, 0x23, 0x45, 0x67, 0x89, 0xAB, 0xCD, 0xEF}

	k, _ := hex.DecodeString("A56BABCD00000000FFFFFFFFABCDEF01")

	plainText := x.Block(data)
	cipherText := x.Block([8]byte{})

	key := x.KeyFromBytes(k)

	cipher := x.Cipher{Key: key}

	cipher.Encrypt(&cipherText, &plainText)

	decryptedText := x.Block([8]byte{})
	cipher.Decrypt(&decryptedText, &cipherText)

	fmt.Println(plainText)
	fmt.Println(cipherText)
	fmt.Println(decryptedText)
}
func test2() {

	data := [8]byte{0x01, 0x23, 0x45, 0x67, 0x89, 0xAB, 0xCD, 0xEF}

	k, _ := hex.DecodeString("A56BABCD00000000FFFFFFFFABCDEF01")

	plainText := data[:]
	cipherText := make([]byte, 8)

	cipher, _ := tea.NewCipher(k)

	cipher.Encrypt(cipherText, plainText)
	decryptedText := make([]byte, 8)
	cipher.Decrypt(decryptedText, cipherText)

	fmt.Println(plainText)
	fmt.Println(cipherText)
	fmt.Println(decryptedText)

}
