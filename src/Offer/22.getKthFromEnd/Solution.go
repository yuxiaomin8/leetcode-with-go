package Offer
type ListNode struct{
	Val int
	Next *ListNode
}
func getKthFromEnd(head *ListNode,k int)*ListNode{
	slow,fast:=head,head
	for ;k>0&&fast!=nil;k--{
		fast=fast.Next
	}
	for fast!=nil{
		slow,fast=slow.Next,fast.Next
	}
	return slow
}