package main

/**
给定一个二维矩阵，计算其子矩形范围内元素的总和，该子矩阵的左上角为 (row1,col1) ，右下角为 (row2,col2)。
上图子矩阵左上角(row1, col1) = (2, 1)，右下角(row2, col2) = (4, 3)，该子矩形内元素的总和为 8。

示例:
给定 matrix = [
  [3, 0, 1, 4, 2],
  [5, 6, 3, 2, 1],
  [1, 2, 0, 1, 5],
  [4, 1, 0, 1, 7],
  [1, 0, 3, 0, 5]
]

sumRegion(2, 1, 4, 3) -> 8
sumRegion(1, 1, 2, 2) -> 11
sumRegion(1, 2, 2, 4) -> 12

说明:
你可以假设矩阵不可变。
会多次调用sumRegion方法。
你可以假设row1 ≤ row2 且col1 ≤ col2。

*/
type NumMatrix struct {
	sums [][]int
}

func Constructor1(matrix [][]int) NumMatrix {
	var array NumMatrix
	m := len(matrix)
	if m > 0 {
		n := len(matrix[0])
		sums := make([][]int, m+1)
		for i := 0; i < m+1; i++ {
			sums[i] = make([]int, n+1)
		}
		array = NumMatrix{sums: sums}

		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				array.sums[i+1][j+1] = array.sums[i][j+1] + array.sums[i+1][j] - array.sums[i][j] + matrix[i][j]
			}
		}
	}
	return array
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	return this.sums[row2+1][col2+1] - this.sums[row2+1][col1] - this.sums[row1][col2+1] + this.sums[row1][col1]
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */

func main() {

}
