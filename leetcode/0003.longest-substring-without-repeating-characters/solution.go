package leetcode

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	left, right, res := 0, 0, 0
	maps := make(map[byte]int, len(s))
	for left < len(s) {
		if i, ok := maps[s[left]]; ok && i >= right {
			right = i + 1
		}
		maps[s[left]] = left
		left++
		if res > left - right{
			res = left -right
		}
	}
	return res
}