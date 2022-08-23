package main

import (
	"github.com/adminsemy/golangTests/structures"
)

func main() {
	slice := []int{1,2,3,3,4,5,5}
	result := structures.CreateListNode(slice)
	mapList := make(map[int]*structures.ListNode)
	structures.DeleteDublicates(&result,mapList)
	structures.DoubleList(result)

}