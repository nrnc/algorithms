package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0
	ans := &ListNode{
		Val: 0,
	}
	head := ans
	for l1 != nil && l2 != nil {
		sum := l1.Val + l2.Val + carry
		val := sum % 10
		carry = sum / 10
		ans.Next = &ListNode{Val: val}
		ans = ans.Next
		l1 = l1.Next
		l2 = l2.Next
	}
	if l1 == nil {
		for l2 != nil {
			sum := l2.Val + carry
			val := sum % 10
			carry = sum / 10
			ans.Next = &ListNode{Val: val}
			ans = ans.Next
			l2 = l2.Next
		}
	}
	if l2 == nil {
		for l1 != nil {
			sum := l1.Val + carry
			val := sum % 10
			carry = sum / 10
			ans.Next = &ListNode{Val: val}
			ans = ans.Next
			l1 = l1.Next
		}
	}
	if carry != 0 {
		ans.Next = &ListNode{Val: carry}
	}
	return head.Next
}
