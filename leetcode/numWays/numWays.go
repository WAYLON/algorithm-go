package main

/**
有一个长度为arrLen的数组，开始有一个指针在索引0 处。
每一步操作中，你可以将指针向左或向右移动 1 步，或者停在原地（指针不能被移动到数组范围外）。
给你两个整数steps 和arrLen ，请你计算并返回：在恰好执行steps次操作以后，指针仍然指向索引0 处的方案数。
由于答案可能会很大，请返回方案数 模10^9 + 7 后的结果。

示例 1：
输入：steps = 3, arrLen = 2
输出：4
解释：3 步后，总共有 4 种不同的方法可以停在索引 0 处。
向右，向左，不动
不动，向右，向左
向右，不动，向左
不动，不动，不动

示例 2：
输入：steps = 2, arrLen = 4
输出：2
解释：2 步后，总共有 2 种不同的方法可以停在索引 0 处。
向右，向左
不动，不动

示例 3：
输入：steps = 4, arrLen = 2
输出：8

提示：
1 <= steps <= 500
1 <= arrLen<= 10^6
*/
func numWays(steps, arrLen int) int {
	const mod = 1e9 + 7
	maxColumn := min(arrLen-1, steps)
	dp := make([][]int, steps+1)
	for i := range dp {
		dp[i] = make([]int, maxColumn+1)
	}
	dp[0][0] = 1
	for i := 1; i <= steps; i++ {
		for j := 0; j <= maxColumn; j++ {
			dp[i][j] = dp[i-1][j]
			if j-1 >= 0 {
				dp[i][j] = (dp[i][j] + dp[i-1][j-1]) % mod
			}
			if j+1 <= maxColumn {
				dp[i][j] = (dp[i][j] + dp[i-1][j+1]) % mod
			}
		}
	}
	return dp[steps][0]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

/**
class Solution {
    public int numWays(int steps, int arrLen) {
        int  min = Math.min(steps / 2 + 1, arrLen);
        int[][] result = new int[2][min + 2];
        result[0][1] = 1;
        for(int i = 1; i <= steps; i++){
            for(int j = 1; j <= min; j++){
                result[i & 1][j] =( ( result[(i - 1) & 1][j - 1] + result[(i - 1) & 1][j] ) % (1000000007) + result[(i - 1) & 1][j + 1] ) % (1000000007);
            }
        }
        return result[steps & 1][1];
    }
}
*/
func main() {

}
