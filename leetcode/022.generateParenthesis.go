package leetcode

func generateParenthesis(n int) []string {
	r := make(map[string]int)
	r["("] = 1
	res := perfixMatch(n, r)
	strs := []string{}
	for k, v := range res {
		if v == 0 {
			strs = append(strs, k)
		}
	}
	return strs
}

func perfixMatch(index int, strs map[string]int) map[string]int {
	r := make(map[string]int)
	for k, v := range strs {
		if v == 0 {
			if len(k) >= index*2 {
				return strs
			}
			r[k+"("] = v + 1
		} else if v > 0 {
			r[k+"("] = v + 1
			r[k+")"] = v - 1
		}
	}
	return perfixMatch(index, r)
}
