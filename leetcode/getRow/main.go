package main

import "fmt"

/**
给定一个非负索引k，其中 k≤33，返回杨辉三角的第 k 行。
在杨辉三角中，每个数是它左上方和右上方的数的和。

示例:
输入: 3
输出: [1,3,3,1]

示例:
输入: 4
输出: [1,4,6,4,1]

进阶：
你可以优化你的算法到 O(k) 空间复杂度吗？
 * 获取杨辉三角的指定行
 * 直接使用组合公式C(n,i) = n!/(i!*(n-i)!)
 * 则第(i+1)项是第i项的倍数=(n-i)/(i+1);
*/
func getRow(rowIndex int) []int {
	res := make([]int, rowIndex+1)
	cur := 1
	for i := 0; i <= rowIndex; i++ {
		res[i] = cur
		cur = cur * (rowIndex - i) / (i + 1)
	}
	return res
}

func main() {
	row := getRow(4)
	fmt.Println(row)
}
