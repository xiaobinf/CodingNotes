package remove_elements_203

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

// 递归解法 O(N) O(N)
//func removeElements(head *ListNode, val int) *ListNode {
//	if head == nil {
//		return head
//	}
//	if head.Val == val {
//		return removeElements(head.Next, val)
//	}
//	head.Next = removeElements(head.Next, val)
//	return head
//}

// 迭代解法 O(N) 空间复杂度O(1)
func removeElements(head *ListNode, val int) *ListNode {
	// 定义一个指向头节点的哑节点
	dummyHead := &ListNode{Next: head}
	for tmp := dummyHead; tmp.Next != nil; {
		if tmp.Next.Val == val {
			tmp.Next = tmp.Next.Next
		} else {
			tmp = tmp.Next
		}
	}
	return dummyHead.Next
}
