/*
 * @lc app=leetcode.cn id=375 lang=golang
 *
 * [375] 猜数字大小 II
 */

// @lc code=start
var cache [210][210]int

func getMoneyAmount(n int) int {
	return getAllAmount(1, n)
}

func getAllAmount(left int, right int) int {
	if left >= right {
		return 0
	}
	if cache[left][right] != 0 {
		return cache[left][right]
	}
	res := 0x3f3f3f3f
	for money := left; money <= right; money++ {
		max := Max(getAllAmount(left, money-1), getAllAmount(money+1, right)) + money
		res = Min(res, max)
	}
	cache[left][right] = res
	return res
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// @lc code=end

// 思路：逆向来看，确保你获胜的最小现金数这个结果应该是从某个数字开始猜的。
// 可以枚举从所有数字为第一个猜到的数字，然后考虑最坏情况，目标应该要花费更多的一边，可以开始dfs判断
// 用数组缓存dfs的结果，减少重复计算
// 然后结果应该所有数字开始路线中的最小值