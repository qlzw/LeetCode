/*
 * @lc app=leetcode.cn id=125 lang=golang
 *
 * [125] 验证回文串
 */

// @lc code=start
func isPalindrome(s string) bool {
	left := 0
	right := len(s) - 1
	s = strings.ToLower(s)
	for left < right {
		for (s[left] < 'a' || s[left] > 'z') && (s[left] < '0' || s[left] > '9') && left < right {
			left++
		}
		for (s[right] < 'a' || s[right] > 'z') && (s[right] < '0' || s[right] > '9') && left < right {
			right--
		}
		if left == right {
			return true
		}

		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

// @lc code=end

