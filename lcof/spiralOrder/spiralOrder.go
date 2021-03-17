package main

/**
输入一个矩阵，按照从外向里以顺时针的顺序依次打印出每一个数字。

示例 1：
输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
输出：[1,2,3,6,9,8,7,4,5]

示例 2：
输入：matrix =[[1,2,3,4],[5,6,7,8],[9,10,11,12]]
输出：[1,2,3,4,8,12,11,10,9,5,6,7]
限制：
0 <= matrix.length <= 100
0 <= matrix[i].length<= 100

*/

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}
	left := 0
	right := len(matrix[0]) - 1
	top := 0
	bottom := len(matrix) - 1
	res := make([]int, (right+1)*(bottom+1))
	x := 0
	for {
		for i := left; i <= right; i++ {
			res[x] = matrix[top][i]
			x++
		}

		top++
		if top > bottom {
			break
		}
		for i := top; i <= bottom; i++ {
			res[x] = matrix[i][right]
			x++
		}

		right--
		if right < left {
			break
		}

		for i := right; i >= left; i-- {
			res[x] = matrix[bottom][i]
			x++
		}

		bottom--
		if top > bottom {
			break
		}
		for i := bottom; i >= top; i-- {
			res[x] = matrix[i][left]
			x++
		}

		left++
		if right < left {
			break
		}
	}
	return res
}

func main() {

}
