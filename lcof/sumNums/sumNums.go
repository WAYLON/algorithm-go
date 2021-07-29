package main

import "math"

/**
求 1+2+...+n ，要求不能使用乘除法、for、while、if、else、switch、case等关键字及条件判断语句（A?B:C）。

示例 1：
输入: n = 3
输出:6

示例 2：
输入: n = 9
输出:45

限制：
1 <= n<= 10000
*/
/**
数学公式 （（首项+尾项）* N ) /2 或 (n的平方 + n) /2
*/
func sumNums(n int) int {
	return (int(math.Pow(float64(n), 2)) + n) >> 1
}

/**
递归短路
*/
func sumNums1(n int) int {
	_ = n > 0 &&
		func() bool {
			n += sumNums1(n - 1)
			return true
		}()
	return n
}

/**
java
class Solution {
    public int sumNums(int n) {
        int sum = n;
        boolean flag = n > 0 && (sum += sumNums(n - 1)) > 0;
        return sum;
    }
}

class Solution {
    public int sumNums(int n) {
        return (int) (Math.pow(n, 2) + n) >> 1;
    }
}

*/

func main() {

}
