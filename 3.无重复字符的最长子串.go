package main
/*
 * @lc app=leetcode.cn id=3 lang=golang
 *
 * [3] 无重复字符的最长子串
 */
// @lc code=start
func lengthOfLongestSubstring(s string) int {
	max := 0
	for i := 0; i < len(s); i++ {
		length := 0
		for j := i; j < len(s); j++ {

			if s[i] == s[j] {
				length++
			} else {
				break
			}

		}

		if max < length {
			max = length
		}

	}

	return max
}
// @lc code=end

