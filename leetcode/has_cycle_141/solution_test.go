package has_cycle_141

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

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow, fast := head, head.Next
	for slow != fast {
		if fast == nil || fast.Next == nil {
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return true
}


// 复习 判断是否有环
func hasCycle(head *ListNode) bool {
	if head==nil ||head.Next==nil{
		return false
	}

	var slow, fast = head, head.Next
	for slow!=fast{
		// 判断是否走到了尾结点
		if fast==nil &&fast.Next==nil{
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next
	}

	return true
}