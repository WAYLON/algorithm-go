package main

import (
	"sort"
	"strings"
)

/**
给定两个字符串 s1 和 s2，请编写一个程序，确定其中一个字符串的字符重新排列后，能否变成另一个字符串。

示例 1：
输入: s1 = "abc", s2 = "bca"
输出: true

示例 2：
输入: s1 = "abc", s2 = "bad"
输出: false

说明：
0 <= len(s1) <= 100
0 <= len(s2) <= 100

*/
func CheckPermutation(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	tmp1 := strings.Split(s1, "")
	tmp2 := strings.Split(s2, "")
	sort.Strings(tmp1)
	res1 := strings.Join(tmp1, "")
	sort.Strings(tmp2)
	res2 := strings.Join(tmp2, "")
	return res1 == res2

}
func main() {

}
