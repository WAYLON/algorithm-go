package main

import "strings"

/**
给你一个字符串 s 和一个整数 k ，请你找出 s 中的最长子串，要求该子串中的每一字符出现次数都不少于 k 。返回这一子串的长度。

示例 1：
输入：s = "aaabb", k = 3
输出：3
解释：最长子串为 "aaa" ，其中 'a' 重复了 3 次。

示例 2：
输入：s = "ababbc", k = 2
输出：5
解释：最长子串为 "ababb" ，其中 'a' 重复了 2 次， 'b' 重复了 3 次。

提示：
1 <= s.length <= 104
s 仅由小写英文字母组成
1 <= k <= 105

分治法

*/
func longestSubstring(s string, k int) (ans int) {
	if s == "" {
		return
	}

	cnt := [26]int{}
	for _, ch := range s {
		cnt[ch-'a']++
	}

	var split byte
	for i, c := range cnt[:] {
		if 0 < c && c < k {
			split = 'a' + byte(i)
			break
		}
	}
	if split == 0 {
		return len(s)
	}

	for _, subStr := range strings.Split(s, string(split)) {
		ans = max(ans, longestSubstring(subStr, k))
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {

}
