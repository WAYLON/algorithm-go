package main

/**
给你一个字符串 s，找到 s 中最长的回文子串。

示例 1：
输入：s = "babad"
输出："bab"
解释："aba" 同样是符合题意的答案。

示例 2：
输入：s = "cbbd"
输出："bb"

示例 3：
输入：s = "a"
输出："a"

示例 4：
输入：s = "ac"
输出："a"

提示：
1 <= s.length <= 1000
s 仅由数字和英文字母（大写和/或小写）组成
*/
/**
方法1 -暴力法
*/
func longestPalindromeSubstring(s string) string {
	n := len(s)
	if n <= 1 {
		return s
	}
	b := []byte(s)
	//记录最长子串
	maxLongest := 1
	//记录起始下标
	startIndex := 0
	//遍历查找所有子串
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if j-i+1 > maxLongest && palindrome(b, i, j) {
				maxLongest = j - i + 1
				startIndex = i
			}
		}
	}
	return string(b[startIndex : startIndex+maxLongest])
}

func palindrome(s []byte, i int, j int) bool {
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

/**
方法2 -动态规划
*/
func longestPalindromeSubstring1(s string) string {
	n := len(s)
	if n <= 1 {
		return s
	}
	b := []byte(s)
	//记录最长子串
	maxLongest := 1
	//记录起始下标
	startIndex := 0

	//dp[i][j] 表示s[i][j] 是否是回文串 构造dp数组 并将对角线赋值为true
	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
		dp[i][i] = true
	}

	//遍历上半部分对角线 从上面开始
	for j := 1; j < n; j++ {
		for i := 0; i < j; i++ {
			if b[i] != b[j] {
				dp[i][j] = false
			} else {
				if j-i < 3 {
					dp[i][j] = true
				} else {
					dp[i][j] = dp[i+1][j-1]
				}
			}
			// 只要 dp[i][j] == true 成立，就表示子串 s[i..j] 是回文，此时记录回文长度和起始位置
			if dp[i][j] && j-i+1 > maxLongest {
				maxLongest = j - i + 1
				startIndex = i
			}

		}
	}
	return string(b[startIndex : startIndex+maxLongest])
}

/**
方法3 -中心扩散法
*/
func longestPalindromeSubstring2(s string) string {
	n := len(s)
	if n <= 1 {
		return s
	}
	b := []byte(s)
	//记录最长子串
	maxLongest := 1
	res := string(b[:1])

	// 中心位置枚举到 len - 2 即可
	for i := 0; i < n-1; i++ {
		oddStr := centerSpread(b, i, i)
		evenStr := centerSpread(b, i, i+1)
		maxLenStr := evenStr
		if len(oddStr) > len(evenStr) {
			maxLenStr = oddStr
		}
		if len(maxLenStr) > maxLongest {
			maxLongest = len(maxLenStr)
			res = maxLenStr
		}
	}
	return res
}

func centerSpread(b []byte, left int, right int) string {
	// left = right 的时候，此时回文中心是一个字符，回文串的长度是奇数
	// right = left + 1 的时候，此时回文中心是一个空隙，回文串的长度是偶数
	i := left
	j := right
	for i >= 0 && j < len(b) {
		if b[i] == b[j] {
			i--
			j++
		} else {
			break
		}
	}
	// 这里要小心，跳出 while 循环时，恰好满足 s.charAt(i) != s.charAt(j)，因此不能取 i，不能取 j
	return string(b[i+1 : j])
}

func main() {

}
