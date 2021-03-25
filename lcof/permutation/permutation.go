package main

/**
输入一个字符串，打印出该字符串中字符的所有排列。
你可以以任意顺序返回这个字符串数组，但里面不能有重复元素。

示例:
输入：s = "abc"
输出：["abc","acb","bac","bca","cab","cba"]

限制：
1 <= s 的长度 <= 8
*/
/**
Golang中的切片与Java中的ArrayList集合类似，进行的是引用传递
*/
func permutation(s string) []string {
	var res []string
	b := []byte(s)
	dfs(&res, b, 0)
	return res
}

func dfs(res *[]string, b []byte, i int) {
	if i == len(b)-1 {
		*res = append(*res, string(b))
		return
	}
	dict := map[byte]bool{}
	for j := i; j < len(b); j++ {
		if dict[b[j]] {
			continue
		}
		dict[b[j]] = true
		b[i], b[j] = b[j], b[i]
		dfs(res, b, i+1)
		b[i], b[j] = b[j], b[i]
	}
}

func main() {
	permutation("abc")
}
