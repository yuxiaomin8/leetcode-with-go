package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
递归
*/
func mergeTwoLists1(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}
	if l1.Val < l2.Val {
		l1.Next = mergeTwoLists1(l1.Next, l2)
		return l1
	}
	l2.Next = mergeTwoLists1(l1, l2.Next)
	return l2
}

func mergeTwoLists2(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	cur := head
	for l1 != nil || l2 != nil {
		if l1 == nil {
			cur.Next = l2
			break
		} else if l2 == nil {
			cur.Next = l1
			break
		} else {
			if l1.Val < l2.Val {
				cur.Next = l1
				l1 = l1.Next
			} else {
				cur.Next = l2
				l2 = l2.Next
			}
			cur = cur.Next
		}
	}
	return head.Next
}
