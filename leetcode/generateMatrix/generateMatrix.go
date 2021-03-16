package main

/**
给你一个正整数n ，生成一个包含 1 到n2所有元素，且元素按顺时针顺序螺旋排列的n x n 正方形矩阵 matrix 。
示例 1：
输入：n = 3
输出：[[1,2,3],[8,9,4],[7,6,5]]

示例 2：
输入：n = 1
输出：[[1]]

提示：
1 <= n <= 20
*/

func generateMatrix(n int) [][]int {
	res := make([][]int, n)
	for i := range res {
		res[i] = make([]int, n)
	}
	l, t := 0, 0
	r, b := n-1, n-1
	m := 1
	for m <= n*n {
		for i := l; i <= r; i++ {
			res[t][i] = m
			m++
		}
		t++
		for i := t; i <= b; i++ {
			res[i][r] = m
			m++
		}
		r--
		for i := r; i >= l; i-- {
			res[b][i] = m
			m++
		}
		b--
		for i := b; i >= t; i-- {
			res[i][l] = m
			m++
		}
		l++
	}
	return res
}

func main() {

}
