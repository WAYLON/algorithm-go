package main

import (
	"math"
	"strconv"
)

/**
解法一 普通解法
*/
func printNumbers(n int) []int {
	size := int(math.Pow10(n) - 1)
	res := make([]int, size)
	for i := 0; i < size; i++ {
		res[i] = i + 1
	}
	return res
}

/**
解法二 大数相加法
*/
func printNumbers1(n int) []int {
	res := []int{}         //结果数组，用来返回答案
	str := make([]byte, n) //字符串，存储每个val
	dfs(0, n, str, &res)   //从str的第0个位置开始填起
	return res
}
func dfs(dep int, n int, str []byte, res *[]int) {
	if dep == n { //前n层已经填完，
		//注意我们是从str的0号位置开始填起，因此此时我们已经得到一个n位树
		val, _ := strconv.Atoi(string(str))
		if val != 0 { //避免将0加入到结果数组
			*res = append(*res, val)
		}
		return

	}
	for i := 0; i < 10; i++ {
		str[dep] = byte('0' + i)
		dfs(dep+1, n, str, res)
	}
}

func main() {
	printNumbers(2)
}
