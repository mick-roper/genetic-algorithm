package algorithm

import "unsafe"

const (
	genes         = "abcdefghijklmnopqrxstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789 []{}1234567890;:'/?,.`~<>"
	letterIdxBits = 6
	letterIdxMask = 1<<letterIdxBits - 1
	letterIdxMax  = 63 / letterIdxBits
)

func randStringBytes(n int) string {
	b := make([]byte, n)

	for i, cache, remain := n-1, r.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = r.Int63(), letterIdxMax
		}

		if idx := int(cache & letterIdxMask); idx < len(genes) {
			b[i] = genes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return *(*string)(unsafe.Pointer(&b))
}

func createGenome(n int) string {
	return randStringBytes(n)
}

func randomGene() string {
	return randStringBytes(1)
}
