package main

/**
输入一个整型数组，数组中的一个或连续多个整数组成一个子数组。求所有子数组的和的最大值。
要求时间复杂度为O(n)。

示例1:
输入: nums = [-2,1,-3,4,-1,2,1,-5,4]
输出: 6
解释:连续子数组[4,-1,2,1] 的和最大，为6。

提示：
1 <=arr.length <= 10^5
-100 <= arr[i] <= 100
*/

/**
方法1 动态规划 时间复杂度O(n) 空间复杂度O(n)
*/
func maxSubArray(nums []int) int {
	max := nums[0]
	//dp[i] 为以nums[i]为结尾的最大值
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		if dp[i-1] > 0 {
			dp[i] = dp[i-1] + nums[i]
		} else {
			dp[i] = nums[i]
		}
		if max < dp[i] {
			max = dp[i]
		}
	}
	return max
}

/**
方法2 滚动数组 时间复杂度O(n) 空间复杂度O(1)
*/
func maxSubArray2(nums []int) int {
	max := nums[0]
	//dp[i] 为以nums[i]为结尾的最大值
	dp := make([]int, 2)
	dp[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		if dp[(i-1)&1] > 0 {
			dp[i&1] = dp[(i-1)&1] + nums[i]
		} else {
			dp[i&1] = nums[i]
		}
		if max < dp[i&1] {
			max = dp[i&1]
		}
	}
	return max
}

/**
方法3 常数 时间复杂度O(n) 空间复杂度O(1)
*/
func maxSubArray3(nums []int) int {
	max := nums[0]
	//dp[i] 为以nums[i]为结尾的最大值
	pre := nums[0]
	suf := 0
	for i := 1; i < len(nums); i++ {
		if pre > 0 {
			suf = pre + nums[i]
		} else {
			suf = nums[i]
		}
		if max < suf {
			max = suf
		}
		pre = suf
	}
	return max
}

func main() {

}
