package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	var pre *ListNode        //定义前驱节点
	var cur *ListNode = head //定义当前节点
	for cur != nil {
		var next *ListNode = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}
