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

//第二种  回溯
var digtalMap map[string]string = map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}
var combinations []string

func letterCombinations2(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	combinations = []string{}
	backtrack(digits, 0, "")
	return combinations
}

func backtrack(digits string, index int, combination string) {
	if len(digits) == index {
		combinations = append(combinations, combination)
	} else {
		digit := string(digits[index])
		letters := digtalMap[digit]
		letterCount := len(letters)
		for i := 0; i < letterCount; i++ {
			backtrack(digits, index+1, combination+string(letters[i]))
		}
	}

}
