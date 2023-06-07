package sortList_148

import "testing"

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	// 分开链表
	if head == nil || head.Next == nil {
		return head
	}

	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	mid := slow.Next
	slow.Next = nil
	left := sortList(head)
	right := sortList(mid)
	res := &ListNode{
		Val:  0,
		Next: nil,
	}
	h := res

	for left != nil && right != nil {
		if left.Val < right.Val {
			h.Next = left
			left = left.Next
		} else {
			h.Next = right
			right = right.Next
		}
		h = h.Next
	}

	if left != nil {
		h.Next = left
	} else {
		h.Next = right
	}

	return res.Next
}

func TestName(t *testing.T) {

}
