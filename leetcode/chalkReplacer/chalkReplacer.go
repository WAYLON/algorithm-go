package main

/**
一个班级里有n个学生，编号为 0到 n - 1。每个学生会依次回答问题，编号为 0的学生先回答，然后是编号为 1的学生，以此类推，直到编号为 n - 1的学生，然后老师会重复这个过程，重新从编号为 0的学生开始回答问题。
给你一个长度为 n且下标从 0开始的整数数组chalk和一个整数k。一开始粉笔盒里总共有k支粉笔。当编号为i的学生回答问题时，他会消耗 chalk[i]支粉笔。如果剩余粉笔数量 严格小于chalk[i]，那么学生 i需要 补充粉笔。
请你返回需要 补充粉笔的学生 编号。

示例 1：
输入：chalk = [5,1,5], k = 22
输出：0
解释：学生消耗粉笔情况如下：
- 编号为 0 的学生使用 5 支粉笔，然后 k = 17 。
- 编号为 1 的学生使用 1 支粉笔，然后 k = 16 。
- 编号为 2 的学生使用 5 支粉笔，然后 k = 11 。
- 编号为 0 的学生使用 5 支粉笔，然后 k = 6 。
- 编号为 1 的学生使用 1 支粉笔，然后 k = 5 。
- 编号为 2 的学生使用 5 支粉笔，然后 k = 0 。
编号为 0 的学生没有足够的粉笔，所以他需要补充粉笔。

示例 2：
输入：chalk = [3,4,1,2], k = 25
输出：1
解释：学生消耗粉笔情况如下：
- 编号为 0 的学生使用 3 支粉笔，然后 k = 22 。
- 编号为 1 的学生使用 4 支粉笔，然后 k = 18 。
- 编号为 2 的学生使用 1 支粉笔，然后 k = 17 。
- 编号为 3 的学生使用 2 支粉笔，然后 k = 15 。
- 编号为 0 的学生使用 3 支粉笔，然后 k = 12 。
- 编号为 1 的学生使用 4 支粉笔，然后 k = 8 。
- 编号为 2 的学生使用 1 支粉笔，然后 k = 7 。
- 编号为 3 的学生使用 2 支粉笔，然后 k = 5 。
- 编号为 0 的学生使用 3 支粉笔，然后 k = 2 。
编号为 1 的学生没有足够的粉笔，所以他需要补充粉笔。


提示：
chalk.length == n
1 <= n <= 105
1 <= chalk[i] <= 105
1 <= k <= 109
*/

/**
前缀和 + 二分
*/
func chalkReplacer(chalk []int, k int) int {
	n := len(chalk)
	sum := make([]int64, n+1)
	for i := 1; i <= n; i++ {
		sum[i] = sum[i-1] + int64(chalk[i-1])
	}
	k = int(int64(k) % sum[n])

	l, r := 1, n
	for l < r {
		mid := (l + r + 1) >> 1
		if sum[mid] <= int64(k) {
			l = mid
		} else {
			r = mid - 1
		}
	}

	if sum[r] <= int64(k) {
		return r
	} else {
		return r - 1
	}
}

/**
模拟
*/
func chalkReplacer2(chalk []int, k int) int {
	n := len(chalk)
	var sum int64 = 0
	for i := 1; i <= n; i++ {
		sum += int64(chalk[i-1])
	}
	k = int(int64(k) % sum)
	i := 0
	for k >= chalk[i] {
		k -= chalk[i]
		i = (i + 1) % n
	}
	return i
}

/** java

class Solution {
    public int chalkReplacer(int[] chalk, int k) {
        long count = 0;
        for (int c : chalk) {
            count += c;
        }
        k = (int)(k % count);
        int i = 0;
        while (k >= chalk[i]) {
            k -= chalk[i];
            i = (i + 1) % chalk.length;
        }
        return i;
    }
}

class Solution {
    public int chalkReplacer(int[] chalk, int k) {
        int n = chalk.length;
        long[] sums = new long[n + 1];
        for (int i = 1; i <= n; i++) {
            sums[i] = sums[i - 1] + chalk[i - 1];
        }
        k = (int) (k % sums[n]);

        int l = 1, r = n;
        while (l < r) {
            int mid = l + r + 1 >> 1;
            if (sums[mid] <= k) {
                l = mid;
            } else {
                r = mid - 1;
            }
        }
        return sums[r] <= k ? r : r - 1;
    }
}

*/
func main() {

}
