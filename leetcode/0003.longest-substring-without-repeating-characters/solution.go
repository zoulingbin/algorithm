/**
https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/
给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。
input: s="abcabcbb"
output: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。

 */
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