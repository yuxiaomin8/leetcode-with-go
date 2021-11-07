package Offer

import (
	"fmt"
	"testing"
)

type question10 struct {
	para10
	ans10
}

// para 是参数
// one 代表第一个参数
type para10 struct {
	one int
}

// ans 是答案
// one 代表第一个答案
type ans10 struct {
	one int
}

func Test_Problem10(t *testing.T) {

	qs := []question10{

		{
			para10{0},
			ans10{0},
		},

		{
			para10{1},
			ans10{1},
		},

		{
			para10{2},
			ans10{1},
		},
		{
			para10{3},
			ans10{2},
		},
		{
			para10{4},
			ans10{3},
		},
		{
			para10{5},
			ans10{5},
		},
		{
			para10{43},
			ans10{5},
		},
		{
			para10{44},
			ans10{5},
		},
		{
			para10{45},
			ans10{5},
		},
		{
			para10{46},
			ans10{5},
		},
	}

	fmt.Printf("------------------------Offer Problem 10------------------------\n")

	for _, q := range qs {
		_, p := q.ans10, q.para10
		fmt.Printf("【input】:%v       【output】:%v\n", p, fib(p.one))
	}
	fmt.Printf("\n\n\n")
}