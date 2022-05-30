package middle_node_876

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

func middleNode(head *ListNode) *ListNode {
  if head.Next == nil {
    return head
  }
  if head.Next.Next == nil {
    return head.Next
  }
  p := head
  q := head
  for q != nil && q.Next != nil {
    p = p.Next
    q = q.Next.Next
  }
  return p
}
