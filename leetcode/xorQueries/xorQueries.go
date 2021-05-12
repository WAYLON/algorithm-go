package main

/**
有一个正整数数组arr，现给你一个对应的查询数组queries，其中queries[i] = [Li,Ri]。
对于每个查询i，请你计算从Li到Ri的XOR值（即arr[Li] xor arr[Li+1] xor ... xor arr[Ri]）作为本次查询的结果。
并返回一个包含给定查询queries所有结果的数组。

示例 1：
输入：arr = [1,3,4,8], queries = [[0,1],[1,2],[0,3],[3,3]]
输出：[2,7,14,8]
解释：
数组中元素的二进制表示形式是：
1 = 0001
3 = 0011
4 = 0100
8 = 1000
查询的 XOR 值为：
[0,1] = 1 xor 3 = 2
[1,2] = 3 xor 4 = 7
[0,3] = 1 xor 3 xor 4 xor 8 = 14
[3,3] = 8

示例 2：
输入：arr = [4,8,2,10], queries = [[2,3],[1,3],[0,0],[0,3]]
输出：[8,0,4,4]

提示：
1 <= arr.length <= 3 *10^4
1 <= arr[i] <= 10^9
1 <= queries.length <= 3 * 10^4
queries[i].length == 2
0 <= queries[i][0] <= queries[i][1] < arr.length
*/

/**
前缀和
*/
func xorQueries(arr []int, queries [][]int) []int {
	n := len(arr)
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = dp[i-1] ^ arr[i-1]
	}
	m := len(queries)
	res := make([]int, m)
	for i := 0; i < m; i++ {
		x := queries[i][0]
		y := queries[i][1]
		res[i] = dp[x] ^ dp[y+1]
	}
	return res
}

/**
树状数组
*/
func xorQueries2(arr []int, queries [][]int) []int {
	n := len(arr)
	m := len(queries)
	c := make([]int, 100009)

	var lowbit func(x int) int
	lowbit = func(x int) int {
		return x & -x
	}

	var add func(x int, u int)
	add = func(x int, u int) {
		for i := x; i <= n; i += lowbit(i) {
			c[i] ^= u
		}
	}

	var query func(x int) int
	query = func(x int) int {
		ans := 0
		for i := x; i > 0; i -= lowbit(i) {
			ans ^= c[i]
		}
		return ans
	}
	for i := 1; i <= n; i++ {
		add(i, arr[i-1])
	}
	res := make([]int, m)
	for i := 0; i < m; i++ {
		l := queries[i][0] + 1
		r := queries[i][1] + 1
		res[i] = query(r) ^ query(l-1)
	}

	return res
}

/**
class Solution {
    int n;
    int[] c = new int[100009];
    int lowbit(int x) {
        return x & -x;
    }
    void add(int x, int u) {
        for (int i = x; i <= n; i += lowbit(i)) c[i] ^= u;
    }
    int query(int x) {
        int ans = 0;
        for (int i = x; i > 0; i -= lowbit(i)) ans ^= c[i];
        return ans;
    }
    public int[] xorQueries(int[] arr, int[][] qs) {
        n = arr.length;
        int m = qs.length;
        for (int i = 1; i <= n; i++) add(i, arr[i - 1]);
        int[] ans = new int[m];
        for (int i = 0; i < m; i++) {
            int l = qs[i][0] + 1, r = qs[i][1] + 1;
            ans[i] = query(r) ^ query(l - 1);
        }
        return ans;
    }
}

class Solution {
    public int[] xorQueries(int[] arr, int[][] queries) {
        int n = arr.length;
        int[] dp = new int[n + 1];
        for (int i = 1; i <= n; i++) {
            dp[i] = dp[i - 1] ^ arr[i - 1];
        }
        int m = queries.length;
        int[] res = new int[m];
        for (int i = 0; i < m; i++) {
            int x = queries[i][0];
            int y = queries[i][1];
            res[i] = dp[x] ^ dp[y + 1];
        }
        return res;
    }
}
*/
func main() {

}
