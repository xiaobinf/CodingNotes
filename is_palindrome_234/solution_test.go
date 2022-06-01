package is_palindrome_234

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

func isPalindrome(head *ListNode) bool {
  if head.Next == nil {
    return true
  }
  // 先找到中间节点mid
  slow, fast := head, head
  for fast != nil && fast.Next != nil {
    slow = slow.Next
    fast = fast.Next.Next
  }

  var mid *ListNode
  var odd bool
  if fast == nil {
    mid = slow
    odd = false
  } else {
    mid = slow.Next
    odd = true
  }

  // 翻转前半部分
  var pre *ListNode
  cur := head
  for cur != mid {
    next := cur.Next
    cur.Next = pre
    pre = cur
    cur = next
  }

  cur = pre
  // 双指针 中间往两边遍历
  if odd {
    cur = cur.Next
  }
  // 注意结束条件
  for cur != nil {
    if cur.Val != mid.Val {
      return false
    }
    cur = cur.Next
    mid = mid.Next
  }
  return true
}
