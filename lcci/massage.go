package main

/**
一个有名的按摩师会收到源源不断的预约请求，每个预约都可以选择接或不接。在每次预约服务之间要有休息时间，因此她不能接受相邻的预约。给定一个预约请求序列，替按摩师找到最优的预约集合（总预约时间最长），返回总的分钟数。
注意：本题相对原题稍作改动

示例 1：
输入： [1,2,3,1]
输出： 4
解释： 选择 1 号预约和 3 号预约，总时长 = 1 + 3 = 4。

示例 2：
输入： [2,7,9,3,1]
输出： 12
解释： 选择 1 号预约、 3 号预约和 5 号预约，总时长 = 2 + 9 + 1 = 12。

示例 3：
输入： [2,1,4,5,3,1,1,3]
输出： 12
解释： 选择 1 号预约、 3 号预约、 5 号预约和 8 号预约，总时长 = 2 + 4 + 3 + 3 = 12。
*/

/**
二维数组动态规划
*/
func massage(nums []int) int {
	l := len(nums)
	if l == 0 {
		return 0
	}
	if l == 1 {
		return nums[0]
	}

	dp := make([][]int, l)
	for i := range dp {
		dp[i] = make([]int, 2) // [0]定义为接，[1]定义为不接
	}
	dp[0][0] = 0
	dp[0][1] = nums[0]

	for i := 1; i < l; i++ {
		dp[i][0] = Max(dp[i-1][0], dp[i-1][1])
		dp[i][1] = dp[i-1][0] + nums[i]
	}
	return Max(dp[l-1][0], dp[l-1][1])
}

/**
一维数组动态规划
*/
func massage1(nums []int) int {
	l := len(nums)
	if l == 0 {
		return 0
	}
	if l == 1 {
		return nums[0]
	}

	dp := make([]int, l)
	dp[0] = nums[0]
	dp[1] = Max(nums[0], nums[1])

	for i := 2; i < l; i++ {
		dp[i] = Max(dp[i-1], dp[i-2]+nums[i])
	}
	return dp[l-1]
}

/**
一维滚动数组动态规划
*/
func massage2(nums []int) int {
	l := len(nums)
	if l == 0 {
		return 0
	}
	if l == 1 {
		return nums[0]
	}

	dp := make([]int, 2)
	dp[0] = nums[0]
	dp[1] = Max(nums[0], nums[1])
	//n对2取模相当于 n&1
	for i := 2; i < l; i++ {
		dp[i%2] = Max(dp[(i-1)%2], dp[(i-2)%2]+nums[i])
	}
	return dp[(l-1)%2]
}

func Max(n1 int, n2 int) int {
	if n1 > n2 {
		return n1
	}
	return n2
}

func main() {

}
