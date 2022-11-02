/*
Алгоритм Кнута Морриса Прата по нахождению в строке подстроки
*/

package kmp

import "fmt"

var pi []int

func CreatePi(str string) {
	var prefics, suffics rune
	var iPrefics int

	strRune := []rune(str)

	pi = make([]int, len(strRune))

	for i, symbol := range strRune {
		pi[i] = 0
		fmt.Println(pi)
		if i == 0 {
			prefics = symbol
			continue
		}
		suffics = symbol

		if suffics != prefics {
			if iPrefics == 0 {
				pi[i] = 0
			} else {
				iPrefics = pi[iPrefics-1]
				prefics = strRune[pi[iPrefics]]
			}
		}

		if suffics == prefics {
			pi[i] = iPrefics + 1
			iPrefics++
		}
		prefics = strRune[pi[iPrefics]]
	}
	fmt.Println(pi)
}
