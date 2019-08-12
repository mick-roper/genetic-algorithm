package algorithm

import "math"

const genes = "abcdefghijklmnopqrstuvwxyz1234567890ABCDEFGHIJKLMNOPQRSTUVWXYZ-=!@£$%^&*(_ .:',?/`~|"

func createGenome(min, max int) string {
	d := r.Intn(max) - min
	var c string
	for i := 0; i < d-min; i++ {
		c += string(genes[int64(math.Floor(r.Float64()*float64(len(genes))))])
	}
	return c
}

func randomGene() string {
	l := len(genes)
	n := r.Intn(l - 1)
	return string(genes[n])
}
