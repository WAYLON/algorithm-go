package main

/**
给你一个字符串 s，请你将 s 分割成一些子串，使每个子串都是回文。
返回符合要求的 最少分割次数 。

示例 1：
输入：s = "aab"
输出：1
解释：只需一次分割就可将s 分割成 ["aa","b"] 这样两个回文子串。

示例 2：
输入：s = "a"
输出：0

示例 3：
输入：s = "ab"
输出：1

提示：
1 <= s.length <= 2000
s 仅由小写英文字母组成

*/

func minCut(s string) int {
	//特例
	l := len(s)
	if l < 2 {
		return 0
	}
	//dp数组记录每一个子串的回文性
	dp := make([][]bool, l)
	for i := range dp {
		dp[i] = make([]bool, l)
	}
	bytes := []byte(s)
	//1.动态规划填充dp数组
	/*for i := 0; i < l; i++ {
		for j := i; j >= 0; j-- {
			if i == j {
				dp[j][i] = true
			} else {
				if i-j+1 == 2 {
					dp[j][i] = bytes[i] == bytes[j]
				} else {
					dp[j][i] = bytes[i] == bytes[j] && dp[j+1][i-1]
				}
			}
		}
	}*/
	//2.方法2
	for right := 0; right < l; right++ {
		for left := 0; left <= right; left++ {
			if bytes[left] == bytes[right] && ((right-left <= 2) || dp[left+1][right-1]) {
				dp[left][right] = true
			}
		}
	}

	// f(i) 代表考虑下标为 i 的字符为结尾的最小分割次数
	f := make([]int, l)
	for i := 1; i < l; i++ {
		if dp[0][i] {
			f[i] = 0
		} else {
			// 如果无法直接构成回文
			// 那么对于第 i 个字符，有使用分割次数，或者不使用分割次数两种选择
			f[i] = f[i-1] + 1
			for j := 1; j < i; j++ {
				if dp[j][i] {
					if f[i] > f[j-1]+1 {
						f[i] = f[j-1] + 1
					}
				}
			}
		}
	}

	return f[l-1]
}

func main() {
	minCut("aabaaassdds")
}
