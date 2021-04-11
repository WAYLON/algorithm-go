package main

import "math"

/**
给你一个长度为n的3 跑道道路，它总共包含n + 1个点，编号为0到n。一只青蛙从0号点第二条跑道出发，它想要跳到点n处。然而道路上可能有一些障碍。
给你一个长度为 n + 1的数组obstacles，其中obstacles[i]（取值范围从 0 到 3）表示在点 i处的obstacles[i]跑道上有一个障碍。如果obstacles[i] == 0，那么点i处没有障碍。任何一个点的三条跑道中最多有一个障碍。
比方说，如果obstacles[2] == 1，那么说明在点 2 处跑道 1 有障碍。
这只青蛙从点 i跳到点 i + 1且跑道不变的前提是点 i + 1的同一跑道上没有障碍。为了躲避障碍，这只青蛙也可以在同一个点处侧跳到 另外一条跑道（这两条跑道可以不相邻），但前提是跳过去的跑道该点处没有障碍。
比方说，这只青蛙可以从点 3 处的跑道 3 跳到点 3 处的跑道 1 。
这只青蛙从点 0 处跑道 2出发，并想到达点 n处的 任一跑道 ，请你返回 最少侧跳次数。
注意：点 0处和点 n处的任一跑道都不会有障碍。

示例 1：
输入：obstacles = [0,1,2,3,0]
输出：2
解释：最优方案如上图箭头所示。总共有 2 次侧跳（红色箭头）。
注意，这只青蛙只有当侧跳时才可以跳过障碍（如上图点 2 处所示）。

示例 2：
输入：obstacles = [0,1,1,3,3,0]
输出：0
解释：跑道 2 没有任何障碍，所以不需要任何侧跳。

示例 3：
输入：obstacles = [0,2,1,0,3,0]
输出：2
解释：最优方案如上图所示。总共有 2 次侧跳。

提示：
obstacles.length == n + 1
1 <= n <= 5 * 105
0 <= obstacles[i] <= 3
obstacles[0] == obstacles[n] == 0
*/
/**
动态规划
*/
func minSideJumps(obstacles []int) int {
	n := len(obstacles)
	//dp[i][j]表示到达第i处第j跑道的最少跳跃次数
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, 3)
	}
	dp[0][0] = 1
	dp[0][1] = 0
	dp[0][2] = 1
	for i := 1; i < n; i++ {
		//首先填充最大值
		dp[i][0] = math.MaxInt32
		dp[i][1] = math.MaxInt32
		dp[i][2] = math.MaxInt32

		//说明当前位置没有障碍物 可以进行平移
		if obstacles[i] != 1 {
			dp[i][0] = dp[i-1][0]
		}
		if obstacles[i] != 2 {
			dp[i][1] = dp[i-1][1]
		}
		if obstacles[i] != 3 {
			dp[i][2] = dp[i-1][2]
		}

		//也可以从其他位置侧跳而来 +1
		if obstacles[i] != 1 {
			dp[i][0] = min(dp[i][0], min(dp[i][1], dp[i][2])+1)
		}
		if obstacles[i] != 2 {
			dp[i][1] = min(dp[i][1], min(dp[i][0], dp[i][2])+1)
		}
		if obstacles[i] != 3 {
			dp[i][2] = min(dp[i][2], min(dp[i][0], dp[i][1])+1)
		}
	}
	return min(dp[n-1][2], min(dp[n-1][0], dp[n-1][1]))
}

func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}

/**
class Solution {
    public int minSideJumps(int[] obstacles) {
        int n = obstacles.length;
        int[][] dp = new int[n][3];

        //dp[i][j]表示到达第i处第j跑道的最少跳跃次数
        dp[0][0] = 1;
        dp[0][1] = 0;
        dp[0][2] = 1;

        for (int i = 1; i < n; ++i) {
            //首先填充最大值
            Arrays.fill(dp[i], Integer.MAX_VALUE - 1);
            //如果障碍物不在跑道1、2、3的位置上时 则dp[i][跑道] = 第i-1处跑道的位置的不跳跃位置
            if (obstacles[i] != 1) dp[i][0] = dp[i - 1][0];
            if (obstacles[i] != 2) dp[i][1] = dp[i - 1][1];
            if (obstacles[i] != 3) dp[i][2] = dp[i - 1][2];
            //或者dp[i][跑道] = 同位置其他跑道跳跃过来+1
            if (obstacles[i] != 1) dp[i][0] = Math.min(dp[i][0], Math.min(dp[i][1], dp[i][2]) + 1);
            if (obstacles[i] != 2) dp[i][1] = Math.min(dp[i][1], Math.min(dp[i][0], dp[i][2]) + 1);
            if (obstacles[i] != 3) dp[i][2] = Math.min(dp[i][2], Math.min(dp[i][0], dp[i][1]) + 1);
        }
        return Math.min(dp[n - 1][0], Math.min(dp[n - 1][1], dp[n - 1][2]));
    }
}
*/
func main() {

}
