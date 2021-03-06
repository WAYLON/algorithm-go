package main

import "container/list"

/**
 地上有一个m行n列的方格，从坐标 [0,0] 到坐标 [m-1,n-1] 。一个机器人从坐标 [0, 0] 的格子开始移动，它每次可以向左、右、上、下移动一格（不能移动到方格外），也不能进入行坐标和列坐标的数位之和大于k的格子。例如，当k为18时，机器人能够进入方格 [35, 37] ，因为3+5+3+7=18。但它不能进入方格 [35, 38]，因为3+5+3+8=19。请问该机器人能够到达多少个格子？

示例 1：
输入：m = 2, n = 3, k = 1
输出：3

示例 2：
输入：m = 3, n = 1, k = 0
输出：1
提示：

1 <= n,m <= 100
0 <= k<= 20
*/

/**
方法1 dfs
*/

var n1, m1, k1 int
var visited [][]bool

func movingCount(m int, n int, k int) int {
	m1 = m
	n1 = n
	k1 = k
	visited = make([][]bool, m)
	for i := 0; i < len(visited); i++ {
		visited[i] = make([]bool, n)
	}

	return dfs(0, 0, 0, 0)
}

func dfs(i int, j int, si int, sj int) int {
	if i >= m1 || j >= n1 || si+sj > k1 || visited[i][j] {
		return 0
	}
	visited[i][j] = true

	sj1 := sj + 1
	si1 := si + 1
	if (j+1)%10 == 0 {
		sj1 = sj - 8
	}
	if (i+1)%10 == 0 {
		si1 = si - 8
	}

	return 1 + dfs(i, j+1, si, sj1) + dfs(i+1, j, si1, sj)
}

/**
方法2 bfs
*/
func movingCount1(m int, n int, k int) int {
	queue := list.List{}
	visited := make([][]bool, m)
	for i := 0; i < len(visited); i++ {
		visited[i] = make([]bool, n)
	}
	queue.PushBack([]int{0, 0, 0, 0})
	res := 0
	for queue.Len() > 0 {
		back := queue.Back()
		queue.Remove(back)
		bfs := back.Value.([]int)
		i := bfs[0]
		j := bfs[1]
		si := bfs[2]
		sj := bfs[3]
		if i >= m || j >= n || si+sj > k || visited[i][j] {
			continue
		}
		res++
		visited[i][j] = true

		sj1 := sj + 1
		si1 := si + 1
		if (j+1)%10 == 0 {
			sj1 = sj - 8
		}
		if (i+1)%10 == 0 {
			si1 = si - 8
		}

		queue.PushBack([]int{i + 1, j, si1, sj})
		queue.PushBack([]int{i, j + 1, si, sj1})
	}
	return res
}

func main() {

}
