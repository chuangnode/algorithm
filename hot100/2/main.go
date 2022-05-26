package main
//两数相加
type ListNode struct {
	Val int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var result = &ListNode{Val: 0}
	var cur = result
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
		sum, carry = sum % 10, sum / 10
		if result != nil {
			cur.Next = &ListNode{
				Val: sum,
			}
			cur = cur.Next
		} else {
			result = &ListNode{Val: sum}
			cur = result
		}
	}
	if carry > 0 {
		cur.Next = &ListNode{Val: carry}
	}
	return result.Next
}

func main()  {
	/*l1 := &ListNode{
		val: 2,
		Next: &ListNode{
			val: 4,
			Next: &ListNode{
				val: 3,
				Next: nil,
			},
		},
	}
	l2 := &ListNode{
		val: 5,
		Next: &ListNode{
			val: 6,
			Next: &ListNode{
				val: 4,
				Next: nil,
			},
		},
	}*/

}
