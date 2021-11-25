package leetcode

func generateParenthesis(n int) []string {
	res := []string{}
	backtrack(n, n, "", &res)
	return res
}

func backtrack(left int, right int, str string, res *[]string) {
	if right == 0 {
		*res = append(*res, str)
	}
	if left > 0 {
		backtrack(left-1, right, str+"(", res)
	}
	if right > left {
		backtrack(left, right-1, str+")", res)
	}
}
