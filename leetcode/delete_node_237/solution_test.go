package delete_node_237

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

func deleteNode(node *ListNode) {
	next := node.Next
	node.Next = next.Next
	node.Val = next.Val
}
