type ListNode struct {
	Val int
	Next *ListNode
}


// addTwoNumbers 两个链表的数相加
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var head,tail *ListNode
	var carry int
	for l1!=nil || l2!=nil {
		var n1, n2 int
		if l1!=nil {
			n1 = l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}

		sum := n1 + n2 + carry
		// 计算当前位和进位
		local := sum%10
		carry = sum/10
		if head == nil {
			head = &ListNode {
				Val: local,
				Next: nil,
			}
			tail = head
		} else {
			tail.Next = &ListNode{
				Val: local,
				Next: nil,
			}
			tail = tail.Next
		}
	}

	if carry > 0 {
		tail.Next = &ListNode{
			Val: carry,
			Next: nil, 	
		}
        tail = tail.Next
	}

	return head
}




func addTwoNumbers(l1, l2 *ListNode) (head *ListNode) {
    var tail *ListNode
    carry := 0
    for l1 != nil || l2 != nil {
        n1, n2 := 0, 0
        if l1 != nil {
            n1 = l1.Val
            l1 = l1.Next
        }
        if l2 != nil {
            n2 = l2.Val
            l2 = l2.Next
        }
        sum := n1 + n2 + carry
        sum, carry = sum%10, sum/10
        if head == nil {
            head = &ListNode{Val: sum}
            tail = head
        } else {
            tail.Next = &ListNode{Val: sum}
            tail = tail.Next
        }
    }
    if carry > 0 {
        tail.Next = &ListNode{Val: carry}
    }
    return
}
