package sample

/**
Definition for singly-linked list.
**/
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	var start1, start2, root, prev *ListNode
	var carry int
	start1 = l1
	start2 = l2
	for start1 != nil || start2 != nil {
		var num1, num2, sum int
		if start1 != nil {
			num1 = start1.Val
			start1 = start1.Next
		}
		if start2 != nil {
			num2 = start2.Val
			start2 = start2.Next
		}
		sum = num1 + num2 + carry
		if sum/10 >= 1 {
			// calculate carry
			carry = sum / 10
			sum = sum % 10
		} else {
			// reset carry
			carry = 0
		}
		if root == nil {
			root = &ListNode{Val: sum}
			prev = root
		} else {
			prev.Next = &ListNode{Val: sum}
			prev = prev.Next
		}
	}
	if carry > 0 {
		// add last node
		prev.Next = &ListNode{Val: carry}
		prev = prev.Next
	}
	return root
}
