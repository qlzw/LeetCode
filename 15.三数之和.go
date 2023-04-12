package leetcode

import "sort"

/*
 * @lc app=leetcode.cn id=15 lang=golang
 *
 * [15] 三数之和
 */
// 思路：排序+双指针
// @lc code=start
func threeSum(nums []int) [][]int {
	n := len(nums)
	sort.Ints(nums)
	ans := make([][]int, 0)
	// 优化1：不需要完整遍历
	for i := 0; i < n-2; i++ {
		// 优化2: 提前返回
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left := i + 1
		right := n - 1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == 0 {
				ans = append(ans, []int{nums[i], nums[left], nums[right]})
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				left++
				right--
			} else if sum > 0 {
				right--
			} else {
				left++
			}

		}
	}
	return ans
}

// @lc code=end
