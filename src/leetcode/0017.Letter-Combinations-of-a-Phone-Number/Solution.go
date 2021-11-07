package leetcode

func letterCombinations(digits string) []string {
	length := len(digits)
	if length == 0 || length > 4 {
		return nil
	}
	m := make(map[byte][]string)
	m['2'] = []string{"a", "b", "c"}
	m['3'] = []string{"d", "e", "f"}
	m['4'] = []string{"g", "h", "i"}
	m['5'] = []string{"j", "k", "l"}
	m['6'] = []string{"m", "o", "n"}
	m['7'] = []string{"p", "q", "s"}
	m['8'] = []string{"t", "u", "v"}
	m['9'] = []string{"w", "x", "y", "z"}

	var ret []string
	for i := 0; i < len(digits); i++ {
		ret = getCombinations(ret, m[digits[i]])
	}
	return ret
}

func getCombinations(a, b []string) []string {
	if len(a) == 0 {
		return b
	}
	if len(b) == 0 {
		return a
	}
	var ret []string
	for _, v := range a {
		for _, bv := range b {
			ret = append(ret, v+bv)
		}
	}
	return ret
}

//第二种
