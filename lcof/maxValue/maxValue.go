package main

/**
在一个 m*n 的棋盘的每一格都放有一个礼物，每个礼物都有一定的价值（价值大于 0）。你可以从棋盘的左上角开始拿格子里的礼物，
并每次向右或者向下移动一格、直到到达棋盘的右下角。给定一个棋盘及其上面的礼物的价值，请计算你最多能拿到多少价值的礼物？

示例 1:
输入:
[
 [1,3,1],
 [1,5,1],
 [4,2,1]
]
输出: 12
解释: 路径 1→3→5→2→1 可以拿到最多价值的礼物

提示：
0 < grid.length <= 200
0 < grid[0].length <= 200
*/

/**
方法1：自身dp
*/
func maxValue(grid [][]int) int {
	n := len(grid)
	m := len(grid[0])
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			max := 0
			if i-1 >= 0 {
				max = grid[i-1][j]
			}

			if j-1 >= 0 && max < grid[i][j-1] {
				max = grid[i][j-1]
			}
			grid[i][j] += max
		}
	}
	return grid[n-1][m-1]
}

/**
方法2：滚动数组
*/
func maxValue2(grid [][]int) int {
	n := len(grid)
	m := len(grid[0])
	dp := make([]int, m+1)
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			dp[j] = max(dp[j], dp[j-1]) + grid[i-1][j-1]
		}
	}
	return dp[m]
}

/**
方法3：正常dp
*/
func maxValue3(grid [][]int) int {
	n := len(grid)
	m := len(grid[0])
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			dp[i][j] = max(dp[i-1][j], dp[i][j-1]) + grid[i-1][j-1]
		}
	}
	return dp[n][m]
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

/**
Java版
class Solution {
    public int maxValue(int[][] grid) {
        //将自身当做dp数组
        int n = grid.length;
        int m = grid[0].length;

        for (int i = 0; i < n; i++) {
            for (int j = 0; j < m; j++) {
                grid[i][j] = Math.max(i - 1 >= 0 ? grid[i - 1][j] : 0, j - 1 >= 0 ? grid[i][j - 1] : 0) + grid[i][j];

            }
        }
        return grid[n - 1][m - 1];
    }
}

class Solution {
    public int maxValue(int[][] grid) {
        int n = grid.length;
        int m = grid[0].length;
        int[][] dp = new int[n + 1][m + 1];
        for (int i = 1; i < n + 1; i++) {
            for (int j = 1; j < m + 1; j++) {
                dp[i][j] = Math.max(dp[i - 1][j], dp[i][j - 1]) + grid[i - 1][j - 1];
            }
        }

        return dp[n][m];
    }
}

class Solution {
    public int maxValue(int[][] grid) {
        int m = grid.length, n = grid[0].length;
        int[] dp = new int[n + 1];
        for (int i = 1; i <= m; i++) {
            for (int j = 1; j <= n; j++) {
                dp[j] = Math.max(dp[j], dp[j - 1]) + grid[i - 1][j - 1];
            }
        }

        return dp[n];
    }
}
*/
func main() {
	maxValue2([][]int{{1, 2, 5}, {3, 2, 1}})
}
