package leetcode

// 最长公共前缀
func longestCommonPrefix(strs []string) string {
	if len(strs)==0 || len(strs[0])==0{
		return ""
	}
	for k, char := range strs[0] {
		for _, str := range strs {
			if len(str) <= k || rune(str[k]) != char {
				return strs[0][:k]
			}
		}
	}
	return strs[0]
}
