package tools

func ChunkByInterface(data []byte, size int) [][]byte {
	length := len(data)
	var res [][]byte
	for i := 0; i < length; i += size {
		end := i + size
		if end > length {
			end = length
		}
		res = append(res, data[i:end])
	}
	return res
}
