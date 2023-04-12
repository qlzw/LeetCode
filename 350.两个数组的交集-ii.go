/*
 * @lc app=leetcode.cn id=350 lang=golang
 *
 * [350] 两个数组的交集 II
 */

// @lc code=start
func intersect(nums1 []int, nums2 []int) []int {
	m := make(map[int]int, 0)
	for _, num := range nums1 {
		m[num]++
	}
	res := make([]int, 0)
	for _, v := range nums2 {
		if m[v] > 0 {
			m[v]--
			res = append(res, v)
		}
	}
	return res
}

// @lc code=end

