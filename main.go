package main

import (
	"encoding/hex"
	"fmt"
	"github.com/aaomidi/blockcipher/cbc"
	"github.com/aaomidi/blockcipher/ecb"

	"github.com/aaomidi/blockcipher/ctr"

	"github.com/aaomidi/blockcipher"
	x "github.com/aaomidi/go-tea/tea"
	"golang.org/x/crypto/tea"
)

func main() {
	test1()
	test2()
	test3()
	testCTR()
	testECB()
	testCBC()
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
	fmt.Println("Test 2 - GoLang's Implementation")
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
func test3() {
	fmt.Println("Test 3 - My Interface Implementation")
	data := [8]byte{0x01, 0x23, 0x45, 0x67, 0x89, 0xAB, 0xCD, 0xEF}

	k, _ := hex.DecodeString("A56BABCD00000000FFFFFFFFABCDEF01")

	plainText := data[:]
	cipherText := make([]byte, 8)

	key := x.KeyFromBytes(k)

	cipher := x.Cipher{Key: key}

	cipher.EncryptBlock(cipherText, plainText)
	decryptedText := make([]byte, 8)
	cipher.DecryptBlock(decryptedText, cipherText)

	fmt.Println(plainText)
	fmt.Println(cipherText)
	fmt.Println(decryptedText)
}

func testCTR() {
	fmt.Println("Test 3 - CTR")

	k, _ := hex.DecodeString("A56BABCD00000000FFFFFFFFABCDEF01")
	key := x.KeyFromBytes(k)

	cipher := x.Cipher{Key: key}

	method := ctr.CTR{
		IVIncrement:  ctr.IVIncrement(&cipher),
		CryptoMethod: blockcipher.CryptoMethod(&cipher),
	}

	sentence := "Four score and seven years ago our fathers brought forth on this continent, a new\nnation, conceived in Liberty, and dedicated to the proposition that all men are\ncreated equal."

	data := []byte(sentence)

	dataBuffer := make([]byte, len(data))
	copy(dataBuffer, data)

	src := x.SliceChunk(dataBuffer, 8)
	dst := make([][]byte, len(src))

	fmt.Println(src)

	method.Apply(dst, src, []byte{0, 5, 0, 0, 0, 0, 0, 2})
	fmt.Println(dst)

	method.Apply(src, dst, []byte{0, 5, 0, 0, 0, 0, 0, 2})
	fmt.Println(src)
	fmt.Println(string(x.JoinBytes(src)))
}

func testECB() {
	k, _ := hex.DecodeString("A56BABCD00000000FFFFFFFFABCDEF01")
	key := x.KeyFromBytes(k)

	cipher := x.Cipher{Key: key}

	method := ecb.ECB{
		CryptoMethod: blockcipher.CryptoMethod(&cipher),
	}

	sentence := "Four score and seven years ago our fathers brought forth on this continent, a new\nnation, conceived in Liberty, and dedicated to the proposition that all men are\ncreated equal."

	data := []byte(sentence)

	dataBuffer := make([]byte, len(data))
	copy(dataBuffer, data)

	src := x.SliceChunk(dataBuffer, 8)
	dst := make([][]byte, len(src))

	fmt.Println(src)

	method.Encrypt(dst, src)
	fmt.Println(dst)

	method.Decrypt(src, dst)
	fmt.Println(src)
	fmt.Println(string(x.JoinBytes(src)))
}

func testCBC() {
	k, _ := hex.DecodeString("A56BABCD00000000FFFFFFFFABCDEF01")
	key := x.KeyFromBytes(k)

	cipher := x.Cipher{Key: key}

	method := cbc.CBC{
		CryptoMethod: blockcipher.CryptoMethod(&cipher),
	}

	sentence := "Four score and seven years ago our fathers brought forth on this continent, a new\nnation, conceived in Liberty, and dedicated to the proposition that all men are\ncreated equal."

	data := []byte(sentence)

	dataBuffer := make([]byte, len(data))
	copy(dataBuffer, data)

	src := x.SliceChunk(dataBuffer, 8)
	dst := make([][]byte, len(src))

	fmt.Println(src)

	method.Encrypt(dst, src, []byte{2, 2, 2, 2, 2, 2, 2, 2})
	fmt.Println(dst)

	method.Decrypt(src, dst, []byte{2, 2, 2, 2, 2, 2, 2, 2})
	fmt.Println(src)
	fmt.Println(string(x.JoinBytes(src)))
}
