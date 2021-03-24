package main

import "math"

/**
给你一个整数数组 nums ，数组中共有 n 个整数。132 模式的子序列 由三个整数 nums[i]、nums[j] 和 nums[k] 组成，并同时满足：i < j < k 和 nums[i] < nums[k] < nums[j] 。
如果 nums 中存在 132 模式的子序列 ，返回 true ；否则，返回 false 。
进阶：很容易想到时间复杂度为 O(n^2) 的解决方案，你可以设计一个时间复杂度为 O(n logn) 或 O(n) 的解决方案吗？

示例 1：
输入：nums = [1,2,3,4]
输出：false
解释：序列中不存在 132 模式的子序列。

示例 2：
输入：nums = [3,1,4,2]
输出：true
解释：序列中有 1 个 132 模式的子序列： [1, 4, 2] 。

示例 3：
输入：nums = [-1,3,2,0]
输出：true
解释：序列中有 3 个 132 模式的的子序列：[-1, 3, 2]、[-1, 3, 0] 和 [-1, 2, 0] 。
提示：
n == nums.length
1 <= n <= 104
-109 <= nums[i] <= 109
*/

func find132pattern(nums []int) bool {
	n := len(nums)
	if n < 3 {
		return false
	}

	//定义极小值来找最小
	min := math.MinInt32
	//维护一个递减栈
	var stack []int
	for i := n - 1; i >= 0; i-- {
		//若当前值小于最小值 则说明存在 1 3 2
		if nums[i] < min {
			return true
		}

		//当前元素大于栈顶元素 循环出栈 赋值最小值
		for len(stack) != 0 && stack[len(stack)-1] < nums[i] {
			min = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
		//将当前值压入栈
		stack = append(stack, nums[i])
	}
	return false

}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 6, 7}
	find132pattern(nums)
}
