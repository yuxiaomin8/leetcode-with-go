package Offer

//type ListNode struct{
//	//Val int
//	//Next *ListNode
//}

type question22 struct {
	para22
	ans22
}

// para 是参数
// one 代表第一个参数
type para22 struct {
	one ListNode
	k int
}

// ans 是答案
// one 代表第一个答案
type ans22 struct {
	one ListNode
}

//func Test_Problem22(t *testing.T) {
//	qs := []question22{
//	}
//
//	fmt.Printf("------------------------Leetcode Problem 22------------------------\n")
//
//	for _, q := range qs {
//		_, p := q.para22, q.ans22
//		//fmt.Printf("【input】:%v       【output】:%v\n", p, getKthFromEnd(p.one))
//	}
//	fmt.Printf("\n\n\n")
//}
