package main

import (
	"fmt"

	"github.com/adminsemy/golangTests/kmp"
)

func main() {
	res := kmp.FindRepeateds("abc", "abcabeabcacabcabd")
	fmt.Println(res)
}
