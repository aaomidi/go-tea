package tea

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

func JoinBytes(buf [][]byte) []byte {
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
