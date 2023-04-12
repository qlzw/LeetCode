/*
 * @lc app=leetcode.cn id=344 lang=golang
 *
 * [344] 反转字符串
 */

// @lc code=start
func reverseString(s []byte) {
	l := 0
	r := len(s) - 1
	for l <= r {
		s[l], s[r] = s[r], s[l]
		l++
		r--
	}
}

// @lc code=end

