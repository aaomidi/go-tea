package tea

func (c *Cipher) EncryptBlock(dst, src []byte) {
	var source Block
	copy(source[:], src)
	var destination Block
	c.Encrypt(&destination, &source)
	copy(dst, destination[:])
}

func (c *Cipher) DecryptBlock(dst, src []byte) {
	var source Block
	copy(source[:], src)
	var destination Block
	c.Decrypt(&destination, &source)
	copy(dst, destination[:])
}
func (c *Cipher) Encrypt(dst, src *Block) {
	v0, v1 := src.Left(), src.Right()

	k := c.Key.Array()

	k0, k1, k2, k3 := k[0], k[1], k[2], k[3]

	sum := uint32(0)

	for i := 0; i < 32; i++ {
		sum += delta

		v0 += ((v1 << 4) + k0) ^ (v1 + sum) ^ ((v1 >> 5) + k1)
		v1 += ((v0 << 4) + k2) ^ (v0 + sum) ^ ((v0 >> 5) + k3)
	}
	dst.fromInt(&v0, &v1)
}

func (c *Cipher) Decrypt(dst, src *Block) {
	v0, v1 := src.Left(), src.Right()

	k := c.Key.Array()

	k0, k1, k2, k3 := k[0], k[1], k[2], k[3]

	delta := uint32(delta)
	sum := delta * uint32(32)

	for i := 0; i < 32; i++ {
		v1 -= ((v0 << 4) + k2) ^ (v0 + sum) ^ ((v0 >> 5) + k3)
		v0 -= ((v1 << 4) + k0) ^ (v1 + sum) ^ ((v1 >> 5) + k1)
		sum -= delta
	}
	dst.fromInt(&v0, &v1)
}
