package reverselinkedlist

/*
leetcode 206 反转一个SingleListNode
题目描述
反转一个SingleListNode。

示例:

输入: 1->2->3->4->5->NULL
输出: 5->4->3->2->1->NULL
*/
type SingleListNode struct {
	Index int
	Next  *SingleListNode
}

// 新建一个SingleListNode
func NewList() *SingleListNode {
	n1 := new(SingleListNode)
	n1.Index = 1
	n2 := new(SingleListNode)
	n2.Index = 2
	n3 := new(SingleListNode)
	n3.Index = 3
	n4 := new(SingleListNode)
	n4.Index = 4
	n5 := new(SingleListNode)
	n5.Index = 5
	n6 := new(SingleListNode)
	n6.Index = 6
	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	n4.Next = n5
	n5.Next = n6
	return n1
}

// https://blog.csdn.net/Charliewolf/article/details/82622014
// 通过迭代方案
func Fun_反转单链表1(head *SingleListNode) *SingleListNode {
	// 先声明两个变量
	// 前一个节点
	var pre *SingleListNode
	// 后一个节点
	var next *SingleListNode

	for head != nil {
		// 保存头节点的下一个节点
		next = head.Next
		// 将头节点指向前一个节点
		head.Next = pre
		// 更新前一个节点
		pre = head
		// 更新头节点
		head = next
	}
	return pre
}

// https://studygolang.com/articles/15711
// 迭代
func Fun_反转单链表2(head *SingleListNode) *SingleListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var prev *SingleListNode
	cur := head
	// 当下一个值的Next不为nil时，即没有循环到最后一个时
	for cur != nil {
		// 1. cur.Next = prev 将下一个和上一个对调
		// 2. prev = cur 将index后移一位，将当前值赋值给前一个值
		// 3. cur = cur.Next 将下一个值的链表作为当前值，达到位移到下一个位置迭代
		cur.Next, prev, cur = prev, cur, cur.Next
	}
	return prev
}

var newList *SingleListNode
var endNode *SingleListNode

// 递归方式
// 进阶：你可以迭代或递归地反转链表。你能否用两种方法解决这道题？
// http://www.voidcn.com/article/p-kvpkwcjz-bpc.html
func Fun_递归反转单链表(head *SingleListNode) {
	if head == nil {
		return
	}

	if nil == head.Next {
		endNode = head
		newList = endNode
		return
	}

	Fun_递归反转单链表(head.Next)
	endNode.Next = head
	endNode = head
	endNode.Next = nil
}

// 思路
// 使用迭代的方式反转链表大家已经很熟了，其实利用递归调用栈的特性，我们也可以轻松做到链表反转。
// 链表反转后，原链表的最后一个结点，会变成新表的头结点。因此我们可以设递归函数总是返回当前链表的最后一个结点，这样最深的一层递归调用就是原链表的尾结点，也就是新表的头结点。此后每一次递归调用结束，调用栈都会回到原表的倒数第二个结点，也就是新表的正数第二结点。确定这些后，我们只需要把每一次递归的遍历出的当前结点按顺序保存起来就好了。
func ReverseList(node *SingleListNode) *SingleListNode {
	Fun_递归反转单链表(node)
	return newList
}
