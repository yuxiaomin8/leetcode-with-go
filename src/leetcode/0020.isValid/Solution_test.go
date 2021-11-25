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
	one string
}

// ans 是答案
// one 代表第一个答案
type ans struct {
	one bool
}

func Test_Problem(t *testing.T) {

	qs := []question{
		{
			para{"()"},
			ans{true},
		},

		{
			para{"()[]{}"},
			ans{true},
		},
		{
			para{"(]"},
			ans{false},
		},
		{
			para{"([)]"},
			ans{false},
		},
		{
			para{"{[]}"},
			ans{true},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 0018------------------------\n")

	for _, q := range qs {
		_, p := q.ans, q.para
		fmt.Printf("【input】:%v       【output】:%v\n", p, isValid(p.one))
	}
	fmt.Printf("\n\n\n")
}
