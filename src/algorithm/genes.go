package algorithm

import "math"

const (
	genes        = "abcdefghijklmnopqrxstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789 "
	genesLength  = len(genes)
	fGenesLength = float64(genesLength)
)

func createGenome(n int) string {
	var c string
	for i := 0; i < n; i++ {
		c += string(genes[int64(math.Floor(r.Float64()*fGenesLength))])
	}
	return c
}

func randomGene() string {
	n := r.Intn(genesLength - 1)
	return string(genes[n])
}
