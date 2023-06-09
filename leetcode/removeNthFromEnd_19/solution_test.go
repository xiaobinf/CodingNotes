type ListNode struct {
	Val int
	Next *ListNode
}

 func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// 设置前置节点，防止头节点被删除
	dummy := &ListNode{0, head}
	var slow,fast = dummy, head
	for i:=0;i<n;i++ {
		fast = fast.Next
	}

	for fast!=nil  {
		slow = slow.Next
		fast = fast.Next
	}

	slow.Next = slow.Next.Next
	return dummy.Next
 }