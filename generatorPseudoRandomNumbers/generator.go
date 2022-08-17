package generatorPseudoRandomNumbers

import (
	"math/rand"
	"time"
)

//Получаем число, до которого надо взять случайное
//число и возваращаем псевдо-случайное число
func Generating(n int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(n)
}