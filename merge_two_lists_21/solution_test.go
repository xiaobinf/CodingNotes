package merge_two_lists_21

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

// 本质上找到递归结束的条件 子问题的定义
//func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
//	if list1 == nil {
//		return list2
//	}
//	if list2 == nil {
//		return list1
//	}
//	if list1.Val < list2.Val {
//		list1.Next = mergeTwoLists(list1.Next, list2)
//		return list1
//	} else {
//		list2.Next = mergeTwoLists(list1, list2.Next)
//		return list2
//	}
//}

// 非递归方式
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	head := &ListNode{
		Val:  0,
		Next: nil,
	}
	cur := head
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			cur.Next = list1
			list1 = list1.Next
			cur = cur.Next
		} else {
			cur.Next = list2
			list2 = list2.Next
			cur = cur.Next
		}
	}

	if list1 == nil {
		cur.Next = list2
	} else {
		cur.Next = list1
	}

	return head.Next
}
