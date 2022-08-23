package main

import (
	"github.com/adminsemy/golangTests/structures"
)

func main() {
	slice := []int{1,2,2,2,3,4,5,5}
	result := structures.CreateListNode(slice)
	structures.DeleteDublicates(&result,nil)
	structures.DoubleList(result)

}