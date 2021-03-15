package main

/**
54. 螺旋矩阵
给你一个 m 行 n 列的矩阵 matrix ，请按照 顺时针螺旋顺序 ，返回矩阵中的所有元素。

示例 1：
输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
输出：[1,2,3,6,9,8,7,4,5]

示例 2：
输入：matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
输出：[1,2,3,4,8,12,11,10,9,5,6,7]

提示：
m == matrix.length
n == matrix[i].length
1 <= m, n <= 10
-100 <= matrix[i][j] <= 100
*/

func spiralOrder(matrix [][]int) []int {
	var res []int
	m := len(matrix)
	n := len(matrix[0])
	//层数 保证数组不为空的情况下都有一层
	count := 0
	if m > n {
		count = (n + 1) / 2
	} else {
		count = (m + 1) / 2
	}
	i := 0
	for i < count {
		//从左向右
		for j := i; j < n-i; j++ {
			res = append(res, matrix[i][j])
		}
		//从上到下
		for j := i + 1; j < m-i; j++ {
			res = append(res, matrix[j][n-1-i])
		}
		//从右向左
		for j := (n - 1) - (i + 1); j >= i && (m-1-i != i); j-- {
			res = append(res, matrix[(m-1)-i][j])
		}
		//从下到上
		for j := (m - 1) - (i + 1); j >= i+1 && (n-1-i) != i; j-- {
			res = append(res, matrix[j][i])
		}
		i++
	}
	return res
}

func main() {

}
