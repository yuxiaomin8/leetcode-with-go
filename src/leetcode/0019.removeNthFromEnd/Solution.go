package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

/***
解法一
*/
//func getLength(head *ListNode)(length int){
//	for ;head!=nil;head=head.Next{
//		length++
//	}
//	return
//}
//
//func removeNthFromEnd(head *ListNode, n int) *ListNode {
//	length:=getLength(head)
//	dummy:=&ListNode{0,head}
//	cur:=dummy
//	for i:=0;i<length-n;i++{
//		cur=cur.Next
//	}
//	cur.Next=cur.Next.Next
//	return dummy.Next
//}
/**
解法2
双指针算法
*/
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{0, head}
	first, second := head, dummy
	for i := 0; i < n; i++ {
		first = first.Next
	}
	for ; first != nil; first = first.Next {
		second = second.Next
	}
	return second
}
