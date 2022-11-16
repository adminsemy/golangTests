/*
Алгоритм Кнута Морриса Прата по нахождению в строке подстроки
*/

package kmp

import (
	"fmt"
)

var pi []int

func CreatePi(str string) {
	var i, iPrefics int
	strRune := []rune(str)
	pi = make([]int, len(strRune))
	i = 1
	for {
		if i == len(strRune) {
			break
		}
		switch {
		case strRune[i] == strRune[iPrefics]:
			pi[i] = iPrefics + 1
			iPrefics++
			i++
		case iPrefics == 0:
			pi[i] = 0
			i++
		default:
			iPrefics = pi[iPrefics-1]
		}

	}
	fmt.Println(pi)
}
