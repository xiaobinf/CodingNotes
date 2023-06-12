package get_intersection_node_160

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

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	p1 := headA
	p2 := headB
	// 当p1=nil p2=nil 不会进度循环
	for p1 != p2 {
		if p1 != nil {
			p1 = p1.Next
		} else {
			p1 = headB
		}

		if p2 != nil {
			p2 = p2.Next
		} else {
			p2 = headA
		}
	}
	return p1
}


// 复习 判断两个链表的相交节点 本质上是使两个指针走的距离相等
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	p1,p2:=headA, headB
	for p1!=p2{
		if p1!=nil{
			p1=p1.Next
		}else{
			p1=headB
		}

		if p2!=nil{
			p2=p2.Next
		}else{
			p2=headA
		}
	}

	return p1
}
