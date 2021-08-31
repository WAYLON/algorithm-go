package main

/**
这里有n个航班，它们分别从 1 到 n 进行编号。
有一份航班预订表bookings ，表中第i条预订记录bookings[i] = [firsti, lasti, seatsi]意味着在从 firsti到 lasti （包含 firsti 和 lasti ）
的 每个航班 上预订了 seatsi个座位。
请你返回一个长度为 n 的数组answer，其中 answer[i] 是航班 i 上预订的座位总数。

示例 1：
输入：bookings = [[1,2,10],[2,3,20],[2,5,25]], n = 5
输出：[10,55,45,25,25]
解释：
航班编号        1   2   3   4   5
预订记录 1 ：   10  10
预订记录 2 ：       20  20
预订记录 3 ：       25  25  25  25
总座位数：      10  55  45  25  25
因此，answer = [10,55,45,25,25]

示例 2：
输入：bookings = [[1,2,10],[2,2,15]], n = 2
输出：[10,25]
解释：
航班编号        1   2
预订记录 1 ：   10  10
预订记录 2 ：       15
总座位数：      10  25
因此，answer = [10,25]

提示：
1 <= n <= 2 * 104
1 <= bookings.length <= 2 * 104
bookings[i].length == 3
1 <= firsti <= lasti <= n
1 <= seatsi <= 104
*/

/**
解题思路：
换一种思路理解题意，将问题转换为：某公交车共有 n 站，第 i 条记录 bookings[i] = [i, j, k] 表示在 i 站上车 k 人，乘坐到 j 站，在 j+1 站下车，需要按照车站顺序返回每一站车上的人数
根据 1 的思路，定义 counter[] 数组记录每站的人数变化，counter[i] 表示第 i+1 站。遍历 bookings[]：bookings[i] = [i, j, k] 表示在 i 站增加 k 人即 counters[i-1] += k，在 j+1 站减少 k 人即 counters[j] -= k
遍历（整理）counter[] 数组，得到每站总人数： 每站的人数为前一站人数加上当前人数变化 counters[i] += counters[i - 1]
*/

func corpFlightBookings(bookings [][]int, n int) []int {
	//记录站的变化
	res := make([]int, n)
	for _, v := range bookings {
		//在v[0]站上车的人
		res[v[0]-1] += v[2]
		if v[1] < n {
			//在v[1]站下车的人
			res[v[1]] -= v[2]
		}
	}

	//每站的人数为前一站人数加上当前人数变化
	for i := 1; i < n; i++ {
		res[i] += res[i-1]
	}

	return res
}

/**
java
class Solution {
    public int[] corpFlightBookings(int[][] bookings, int n) {
        //每站的人数变化
        int[] res = new int[n];
        for (int[] booking : bookings) {
            res[booking[0] - 1] += booking[2];
            if (booking[1] < n) {
                res[booking[1]] -= booking[2];
            }
        }

        for (int i = 1; i < n; i++) {
            res[i] += res[i - 1];
        }

        return res;
    }
}
*/
func main() {

}
