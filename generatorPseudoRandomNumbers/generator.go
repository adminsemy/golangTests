package generatorPseudoRandomNumbers

import (
	"math/rand"
	"time"
)

func Generating(n int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(n)
}