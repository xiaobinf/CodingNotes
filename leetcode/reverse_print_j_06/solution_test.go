package reverse_print_j_06

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

func reversePrint(head *ListNode) []int {
	if head == nil {
		return []int{}
	}
	if head.Next == nil {
		return []int{head.Val}
	}
	// 辅助头节点
	p := &ListNode{}
	c := head
	n := head.Next
	// 三个指针 调整节点方向
	for n != nil {
		// 调整指针方向
		n = c.Next
		c.Next = p
		//整体后移
		p = c
		c = n
		n = n.Next
	}
	c.Next = p
	var ret []int
	// 忽略辅助头节点
	for c != nil && c.Next != nil {
		ret = append(ret, c.Val)
		c = c.Next
	}
	return ret
}
