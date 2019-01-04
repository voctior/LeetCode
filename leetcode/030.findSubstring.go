package leetcode

func findSubstring(s string, words []string) []int {
	result := []int{}

	if len(words) == 0 {
		return result
	}

	wordMap := map[string]int{}
	for _, v := range words {
		wordMap[v] = wordMap[v] + 1
	}
	wordLen := len(words[0])
	window := len(words) * wordLen

	for i := 0; i < wordLen; i++ {
		for j := i; j+window <= len(s); j = j + wordLen {
			tmpStr := s[j : j+window]
			temMap := make(map[string]int)

			for k := len(words) - 1; k >= 0; k-- {
				word := tmpStr[k*wordLen : (k+1)*wordLen]
				temMap[word]++
				if temMap[word] > wordMap[word] {
					j = j + k*wordLen
					break
				} else if k == 0 {
					//这里不用担心存在tempMap[word] < wordMap[word]的情况，
					//在单词窗口长度与words单词长度和相等的情况下，
					//若存在某个单词wordMap有而tempMap没有，
					//那么必然存在另外一个单词tempMap有而wordMap没有。
					//即必定会进入到上面if的语句中。
					result = append(result, j)
				}
			}
		}
	}
	return result
}
