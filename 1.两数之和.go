/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */
// 思路：遍历时使用hashtable将元素索引存储，判断table[target-i]是否在数组中
// @lc code=start
func twoSum(nums []int, target int) []int {
	table := map[int]int{}
	for i, num := range nums {
		if p, ok := table[target-num]; ok {
			return []int{i, p}
		}
		// 后加入哈希表，防止把自己加入答案，例如6+6=12
		table[num] = i
	}
	return nil
}

// @lc code=end

