package algorithm

import (
	"math/rand"
	"time"
)

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

func randomInt(min, max int) int {
	return r.Intn(max-min) + min
}
