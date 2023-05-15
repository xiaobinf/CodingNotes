package swapPairs_24

import "testing"

/**
* Definition for singly-linked list.
* type ListNode struct {
*     Val int
*     Next *ListNode
* }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

// 两两交换链表的节点
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	dummy := &ListNode{Next: head}
	p := dummy
	for p.Next != nil && p.Next.Next != nil {
		node1 := p.Next
		node2 := p.Next.Next
		p.Next, node1.Next, node2.Next = node2, node2.Next, node1
		p = node1
	}

	return dummy.Next
}

func swapPairs1(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	prev := dummy
	for prev.Next != nil && prev.Next.Next != nil {
		node1 := prev.Next
		node2 := node1.Next
		prev.Next, node1.Next, node2.Next = node2, node2.Next, node1
		prev = node1
	}
	return dummy.Next
}

func TestName(t *testing.T) {

}
