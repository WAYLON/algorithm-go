package main

/**
给定一个字符串 s，将 s 分割成一些子串，使每个子串都是回文串。
返回 s 所有可能的分割方案。

示例:
输入:"aab"
输出:
[
  ["aa","b"],
  ["a","a","b"]
]

*/
func partition(s string) (ans [][]string) {
	n := len(s)
	f := make([][]bool, n)
	for i := range f {
		f[i] = make([]bool, n)
		for j := range f[i] {
			f[i][j] = true
		}
	}
	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			f[i][j] = s[i] == s[j] && f[i+1][j-1]
		}
	}

	var splits []string
	var dfs func(int)
	dfs = func(i int) {
		if i == n {
			ans = append(ans, append([]string(nil), splits...))
			return
		}
		for j := i; j < n; j++ {
			if f[i][j] {
				splits = append(splits, s[i:j+1])
				dfs(j + 1)
				splits = splits[:len(splits)-1]
			}
		}
	}
	dfs(0)
	return
}

/**
中心扩散法
*/
func partition1(s string) [][]string {
	var ans [][]string
	n := len(s)
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
	}

	//进行预处理，利用中心扩展 将所有回文子串的位置存储到 dp 中
	prePro := func(s string, left int, right int, dp [][]bool) {
		for left >= 0 && right < len(s) && s[left] == s[right] {
			dp[left][right] = true
			left--
			right++
		}
	}

	for i := 0; i < n; i++ {
		prePro(s, i, i, dp)
		prePro(s, i, i+1, dp)
	}

	var splits []string
	var dfs func(int)
	dfs = func(i int) {
		if i == n {
			ans = append(ans, append([]string(nil), splits...))
			return
		}
		for j := i; j < n; j++ {
			if dp[i][j] {
				splits = append(splits, s[i:j+1])
				dfs(j + 1)
				splits = splits[:len(splits)-1]
			}
		}
	}
	dfs(0)
	return ans
}
func main() {
	partition1("aab")
}
