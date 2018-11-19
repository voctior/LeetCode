package leetcode

func letterCombinations(digits string) []string {
	if digits == "" {
		return []string{}
	}
	return prefixletterCombinations(digits, 0, nil)
}

func prefixletterCombinations(str string, length int, prefix []string) (r []string) {
	if length == 0 {
		return prefixletterCombinations(str, 1, itemStr(string(str[length])))
	}
	if length >= len(str) {
		return prefix
	}
	itemStrs := itemStr(string(str[length]))
	for _, itemp := range prefix {
		for _, items := range itemStrs {
			r = append(r, itemp+items)
		}
	}
	return prefixletterCombinations(str, length+1, r)
}

func itemStr(str string) (r []string) {
	switch str {
	case "2":
		r = []string{"a", "b", "c"}
	case "3":
		r = []string{"d", "e", "f"}
	case "4":
		r = []string{"g", "h", "i"}
	case "5":
		r = []string{"j", "k", "l"}
	case "6":
		r = []string{"m", "n", "o"}
	case "7":
		r = []string{"p", "q", "r", "s"}
	case "8":
		r = []string{"t", "u", "v"}
	case "9":
		r = []string{"w", "x", "y", "z"}
	default:
		r = []string{""}

	}
	return
}
