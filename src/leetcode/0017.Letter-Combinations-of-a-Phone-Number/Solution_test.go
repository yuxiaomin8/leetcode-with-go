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
	one []string
}

func Test_Problem(t *testing.T) {

	qs := []question{

		{
			para{"23"},
			ans{[]string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}},
		},

		{
			para{""},
			ans{[]string{}},
		},

		{
			para{"2"},
			ans{[]string{"a", "b", "c"}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 0017------------------------\n")

	for _, q := range qs {
		_, p := q.ans, q.para
		fmt.Printf("【input】:%v       【output】:%v\n", p, letterCombinations(p.one))
	}
	fmt.Printf("\n\n\n")
}
