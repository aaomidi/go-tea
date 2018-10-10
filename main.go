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
	testCTR()
}
func test1() {
	fmt.Println("Test 1 - My implementation")

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
	fmt.Println("Test 2 - GoLang's Implementation ")
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

func testCTR() {
	k, _ := hex.DecodeString("A56BABCD00000000FFFFFFFFABCDEF01")
	key := x.KeyFromBytes(k)

	cipher := x.Cipher{Key: key}

	sentence := "Four score and seven years ago our fathers brought forth on this continent, a new\nnation, conceived in Liberty, and dedicated to the proposition that all men are\ncreated equal."
	fmt.Println(sentence)

	data := []byte(sentence)

	dataBuffer := make([]byte, len(data))
	copy(dataBuffer, data)

	var blocks []x.Block
	chunks := x.SliceChunk(dataBuffer, 8)

	for _, v := range chunks {
		var block [8]byte
		copy(block[:], v)

		blocks = append(blocks, block)
	}

	result := make([]x.Block, len(blocks))

	// Encrypt
	cipher.CTR(result, blocks, [6]byte{0, 0, 0, 0, 0, 0})

	// Decrypt
	cipher.CTR(blocks, result, [6]byte{0, 0, 0, 0, 0, 0})

	decryptedSentence := x.JoinBlocks(blocks)

	fmt.Println(string(data))
	fmt.Println(string(decryptedSentence))
}
