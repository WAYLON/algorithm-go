package main

/**
给定平面上n 对 互不相同 的点points ，其中 points[i] = [xi, yi] 。回旋镖 是由点(i, j, k) 表示的元组 ，其中i和j之间的距离和i和k之间的距离相等（需要考虑元组的顺序）。
返回平面上所有回旋镖的数量。

示例 1：
输入：points = [[0,0],[1,0],[2,0]]
输出：2
解释：两个回旋镖为 [[1,0],[0,0],[2,0]] 和 [[1,0],[2,0],[0,0]]

示例 2：
输入：points = [[1,1],[2,2],[3,3]]
输出：2

示例 3：
输入：points = [[1,1]]
输出：0

提示：
n ==points.length
1 <= n <= 500
points[i].length == 2
-104 <= xi, yi <= 104
所有点都 互不相同
*/
func numberOfBoomerangs(points [][]int) int {
	n := len(points)
	res := 0
	for i := 0; i < n; i++ {
		m := make(map[int]int)
		for j := 0; j < n; j++ {
			if j == i {
				continue
			}
			x := points[i][0] - points[j][0]
			y := points[i][1] - points[j][1]
			z := x*x + y*y
			m[z] += 1
		}
		for _, v := range m {
			res += v * (v - 1)
		}
	}

	return res
}

/**
class Solution {
    public int numberOfBoomerangs(int[][] points) {
        int n = points.length;
        int ans = 0;
        for (int i = 0; i < n; i++) {
            Map<Integer, Integer> map = new HashMap<>();
            for (int j = 0; j < n; j++) {
                if (i == j) continue;
                int x = points[i][0] - points[j][0], y = points[i][1] - points[j][1];
                int dist = x * x + y * y;
                map.put(dist, map.getOrDefault(dist, 0) + 1);
            }
            for (int dist : map.keySet()) {
                int cnt = map.get(dist);
                ans += cnt * (cnt - 1);
            }
        }
        return ans;
    }
}
*/
func main() {

}
