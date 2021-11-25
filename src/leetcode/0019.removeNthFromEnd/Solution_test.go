package leetcode

import (
	"fmt"
	"github.com/yuxiaomin8/leetcode-with-go/src/structures"
	"testing"
)

//type ListNode struct {
//	Val int
//	Next *ListNode
//}

type question struct {
	para
	ans
}

// para 是参数
// one 代表第一个参数
type para struct {
	one []int
	n   int
}

// ans 是答案
// one 代表第一个答案
type ans struct {
	one []int
}

func Test_Problem(t *testing.T) {

	qs := []question{

		{
			para{[]int{1}, 3},
			ans{[]int{1}},
		},

		{
			para{[]int{1, 2}, 2},
			ans{[]int{2}},
		},

		{
			para{[]int{1}, 1},
			ans{[]int{}},
		},

		{
			para{[]int{1, 2, 3, 4, 5}, 10},
			ans{[]int{1, 2, 3, 4, 5}},
		},

		{
			para{[]int{1, 2, 3, 4, 5}, 1},
			ans{[]int{1, 2, 3, 4}},
		},

		{
			para{[]int{1, 2, 3, 4, 5}, 2},
			ans{[]int{1, 2, 3, 5}},
		},

		{
			para{[]int{1, 2, 3, 4, 5}, 3},
			ans{[]int{1, 2, 4, 5}},
		},
		{
			para{[]int{1, 2, 3, 4, 5}, 4},
			ans{[]int{1, 3, 4, 5}},
		},

		{
			para{[]int{1, 2, 3, 4, 5}, 5},
			ans{[]int{2, 3, 4, 5}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 0019------------------------\n")

	for _, q := range qs {
		_, p := q.ans, q.para
		fmt.Printf("【input】:%v       【output】:%v\n", p, structures.List2Ints(removeNthFromEnd(structures.Ints2List(p.one), p.n)))
	}
	fmt.Printf("\n\n\n")
}
