/*
 * @lc app=leetcode.cn id=240 lang=golang
 *
 * [240] 搜索二维矩阵 II
 */

// @lc code=start
func searchMatrix(matrix [][]int, target int) bool {
	var i, j, h, w int
	h = len(matrix)
	w = len(matrix[0])

	i = h - 1
	j = 0

	res := false

	for i >= 0 && j < w {
		if matrix[i][j] == target {
			res = true
			break
		} else if matrix[i][j] < target {
			j++
		} else {
			i--
		}
	}
	return res
}

// @lc code=end

