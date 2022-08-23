package structures

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func CreateListNode (arr []int) ListNode {
	list := ListNode{}
	ptr := &list
	for _, v := range arr {
		ptr.Val = v
		ptr.Next = new(ListNode)
		ptr = ptr.Next
	}

	return list
}

func DoubleList (l ListNode) {
	if l.Next ==  nil {
		return
	}
	fmt.Println(l.Val)
	DoubleList(*l.Next)
}

func DeleteDublicates(l, parent *ListNode) {
	if l.Next ==  nil {
		return
	}
	if parent == nil {
		DeleteDublicates(l.Next, l)
		return
	}
	if l.Val == parent.Val {
		parent.Next = l.Next
	}
	DeleteDublicates(l.Next, l)
}