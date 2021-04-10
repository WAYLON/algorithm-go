package main

import "math"

/**
编写一个程序，找出第 n 个丑数。
丑数就是质因数只包含2, 3, 5 的正整数。

示例:
输入: n = 10
输出: 12
解释: 1, 2, 3, 4, 5, 6, 8, 9, 10, 12 是前 10 个丑数。

说明:
1是丑数。
n不超过1690。
*/
func nthUglyNumber(n int) int {
	nums := make([]int, n)
	nums[0] = 1
	i2 := 0
	i3 := 0
	i5 := 0
	for i := 1; i < n; i++ {
		a2 := nums[i2] * 2
		a3 := nums[i3] * 3
		a5 := nums[i5] * 5
		ugly := int(math.Min(float64(a2), math.Min(float64(a3), float64(a5))))
		if ugly == a2 {
			i2++
		}
		if ugly == a3 {
			i3++
		}
		if ugly == a5 {
			i5++
		}
		nums[i] = ugly
	}
	return nums[n-1]
}

/**
java版
class Solution {
    public int nthUglyNumber(int n) {
        int[] nums = new int[n];
        nums[0] = 1;
        int i2 = 0, i3 = 0, i5 = 0;
        for (int i = 1; i < n; i++) {
            int a2 = nums[i2] * 2;
            int a3 = nums[i3] * 3;
            int a5 = nums[i5] * 5;
            int ugly = Math.min(a2, Math.min(a3, a5));
            if (ugly == a2) {
                i2++;
            }
            if (ugly == a3) {
                i3++;
            }
            if (ugly == a5) {
                i5++;
            }
            nums[i] = ugly;
        }

        return nums[n - 1];
    }
}
*/
func main() {

}
