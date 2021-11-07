package leetcode

import (
	"fmt"
	"testing"
)

type question struct {
	para
	ans
}

// para 是参数
// one 代表第一个参数
type para struct {
	one []int
	two int
}

// ans 是答案
// one 代表第一个答案
type ans struct {
	one int
}

func Test_Problem(t *testing.T) {

	qs := []question{

		{
			para{[]int{2, 7, 4, 1, 8, 1}, 2},
			ans{1},
		},

		{
			para{[]int{21, 26, 31, 33, 40}, 100},
			ans{5},
		},

		{
			para{[]int{1, 2, 2, 5}, 9},
			ans{1},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 0016------------------------\n")

	for _, q := range qs {
		_, p := q.ans, q.para
		fmt.Printf("【input】:%v       【output】:%v\n", p, ThreeSumClose(p.one, p.two))
	}
	fmt.Printf("\n\n\n")
}
