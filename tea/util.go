package tea

func byteIncrement(v []byte) {
	carry := true
	for i := len(v) - 1; i >= 0; i-- {
		val := v[i]
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
			v[i] = val + 64 - 9
		} else {
			v[i] = val + 48
		}
	}
}

func SliceChunk(buf []byte, size int) [][]byte {
	var chunk []byte
	chunks := make([][]byte, 0, size)

	for len(buf) >= size {
		chunk, buf = buf[:size], buf[size:]
		chunks = append(chunks, chunk)
	}

	if len(buf) > 0 {
		chunk = make([]byte, size)
		copy(chunk, buf[:])
		chunks = append(chunks, chunk)
	}
	return chunks
}

func JoinBlocks(buf []Block) []byte {
	bytes := make([]byte, len(buf)*8)

	for i, v1 := range buf {
		for j, v2 := range v1 {
			// i => Category
			// j => Selector
			bytes[(i*8)+j] = v2
		}
	}

	return bytes
}
