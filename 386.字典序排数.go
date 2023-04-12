/*
 * @lc app=leetcode.cn id=386 lang=golang
 *
 * [386] 字典序排数
 */

// @lc code=start

func lexicalOrder(n int) []int {
	var res []int
	j := 1
	for i := 0; i < n; i++ {
		res = append(res, j)
		if j*10 <= n {
			j = j * 10
		} else {
			for j%10 == 9 || j+1 > n {
				j = j / 10
			}
			j++
		}
	}
	return res
}

// @lc code=end

// 思路：利用个位达到9或者超过n来是j=j/10达成回溯

