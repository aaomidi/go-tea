package tea

import (
	"encoding/binary"
)

const (
	delta uint32 = 0x9e3779b9
)

var (
	e = binary.BigEndian
)

type Key [16]byte
type Block [8]byte
type Feistel = uint32
type IV = [6]byte

type Cipher struct {
	Key *Key
}

func (b *Block) Left() Feistel {
	return e.Uint32(b[:4])
}

func (b *Block) Right() Feistel {
	return e.Uint32(b[4:])
}
func (b *Block) Source() []byte {
	return b[0:8]
}

func (b *Block) fromInt(v0, v1 *uint32) {
	e.PutUint32(b[:4], *v0)
	e.PutUint32(b[4:], *v1)
}

func BlockFromBytes(b []byte) *Block {
	var info [8]byte

	copy(info[0:8], b)

	r := Block(info)

	return &r
}

func KeyFromBytes(b []byte) *Key {
	var info [16]byte

	copy(info[0:16], b)

	r := Key(info)

	return &r
}

func (k *Key) Array() [4]uint32 {
	var result [4]uint32

	result[0] = e.Uint32(k[0:])
	result[1] = e.Uint32(k[4:])
	result[2] = e.Uint32(k[8:])
	result[3] = e.Uint32(k[12:])

	return result
}

func (c *Cipher) Increment(iv []byte) {
	carry := true
	for i := len(iv) - 1; i >= 0; i-- {
		val := iv[i]
		// A->F
		if val > 64 {
			val -= 64 - 9
		} else {
			val -= 48
		}
		if carry {
			val += 1
			carry = false
		}
		if val == 16 {
			val = 0
			carry = true
		}

		if val >= 10 {
			iv[i] = val + 64 - 9
		} else {
			iv[i] = val + 48
		}
	}
}
