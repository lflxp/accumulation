package reverselinkedlist

import (
	"fmt"
	"testing"
)

func printNode(head *SingleListNode) string {
	for head != nil {
		fmt.Println(head)
		head = head.Next
	}
	fmt.Println()
	return "finished"
}

func Test_反转链表1(t *testing.T) {
	origin := NewList()
	t.Log("origin", printNode(origin))

	result1 := 反转单链表1(origin)
	t.Log("result1", printNode(result1))
}

func Test_反转链表2(t *testing.T) {
	origin := NewList()
	t.Log("origin", printNode(origin))

	result2 := 反转单链表2(origin)
	t.Log("result2", printNode(result2))
}

func Test_递归反转链表(t *testing.T) {
	origin := NewList()
	t.Log("origin", printNode(origin))

	result2 := reverseList(origin)
	t.Log("result2", printNode(result2))
}
