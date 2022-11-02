package main

import (
	"fmt"

	"github.com/adminsemy/golangTests/bankomat"
)

func main() {
	fmt.Println(bankomat.GiveMeMoney(13, []int32{1, 4, 6}))
}
