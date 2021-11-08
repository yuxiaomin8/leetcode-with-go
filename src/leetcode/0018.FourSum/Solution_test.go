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
	one [][]int
}

func Test_Problem(t *testing.T) {

	qs := []question{
		{
			para{[]int{1, 0, -1, 0, -2, 2}, 8},
			ans{[][]int{{-2, -1, 1, 2}, {-2, 0, 0, 2}, {-1, 0, 0, 1}}},
		},

		{
			para{[]int{1, 0, -1, 0, -2, 2, 0, 0, 0, 0}, 1},
			ans{[][]int{{-1, 0, 0, 2}, {-2, 0, 1, 2}, {0, 0, 0, 1}}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 0018------------------------\n")

	for _, q := range qs {
		_, p := q.ans, q.para
		fmt.Printf("【input】:%v       【output】:%v\n", p, fourSum(p.one, p.two))
	}
	fmt.Printf("\n\n\n")
}
