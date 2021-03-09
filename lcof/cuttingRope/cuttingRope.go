package main

import "math"

/**
给你一根长度为 n 的绳子，请把绳子剪成整数长度的 m 段（m、n都是整数，n>1并且m>1），每段绳子的长度记为 k[0],k[1]...k[m-1] 。
请问 k[0]*k[1]*...*k[m-1] 可能的最大乘积是多少？例如，当绳子的长度是8时，我们把它剪成长度分别为2、3、3的三段，此时得到的最大乘积是18。

示例 1：
输入: 2
输出: 1
解释: 2 = 1 + 1, 1 × 1 = 1

示例2:
输入: 10
输出: 36
解释: 10 = 3 + 3 + 4, 3 ×3 ×4 = 36
提示：

2 <= n <= 58
*/
/**
方法一 动态规划
首先思考怎么转移 即dp[i] = max(每次切割完不再切割即 j * (i-j)，每次切割完后面的j * dp[i-j])
其次思考初始值 dp[2] = 1
*/
func cuttingRope(n int) int {
	dp := make([]int, n+1)
	dp[2] = 1
	for i := 3; i < n+1; i++ {
		for j := 2; j < i; j++ {
			dp[i] = max(dp[i], max((i-j)*j, j*dp[i-j]))
		}
	}
	return dp[n]
}

func max(m int, n int) int {
	if m > n {
		return m
	}
	return n
}

/**
数学推导法 即
当 b = 0 b=0 时，直接返回 3^a3
当 b = 1 b=1 时，要将一个 1 + 3 转换为 2+2，因此返回 3^{a-1}×4；
当 b = 2 b=2 时，返回 3^a×2。
*/
func cuttingRope1(n int) int {
	if n <= 3 {
		return n - 1
	}
	a := n / 3
	b := n % 3

	if b == 0 {
		return int(math.Pow(float64(3), float64(a)))
	} else if b == 1 {
		return int(math.Pow(float64(3), float64(a)-1) * 4)
	}
	return int(math.Pow(float64(3), float64(a)) * 2)

}

func main() {

}
