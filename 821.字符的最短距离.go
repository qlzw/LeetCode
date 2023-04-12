/*
 * @lc app=leetcode.cn id=821 lang=golang
 *
 * [821] 字符的最短距离
 */

// @lc code=start
func shortestToChar(s string, c byte) []int {
	n := len(s)
	res := make([]int, n)
	pos := -n
	for i := 0; i < len(s); i++ {
		if s[i] == c {
			pos = i
		}
		res[i] = i - pos
	}
	pos = 2 * n
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == c {
			pos = i
		}
		if pos-i < res[i] {
			res[i] = pos - i
		}
	}
	return res
}

// @lc code=end

// 思路：距离问题可以考虑利用左右各遍历一次达成，当未找到
// 目标的时候可以给个大于数组长度的值或者无穷大
