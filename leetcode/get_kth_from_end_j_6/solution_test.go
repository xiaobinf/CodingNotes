package get_kth_from_end_j_6

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

func getKthFromEnd(head *ListNode, k int) *ListNode {
  s, f := head, head
  for k > 0 {
    f = f.Next
    k--
  }

  for f != nil {
    s = s.Next
    f = f.Next
  }

  return s
}
