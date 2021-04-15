package main

import "strconv"

/**
给定一个数字，我们按照如下规则把它翻译为字符串：0 翻译成 “a” ，1 翻译成 “b”，……，11 翻译成 “l”，……，25 翻译成 “z”。一个数字可能有多个翻译。请编程实现一个函数，用来计算一个数字有多少种不同的翻译方法。

示例 1:
输入: 12258
输出: 5
解释: 12258有5种不同的翻译，分别是"bccfi", "bwfi", "bczi", "mcfi"和"mzi"

提示：
0 <= num < 231
*/

/**
方法1：字符串遍历
*/
func translateNum(num int) int {
	bytes := []byte(strconv.Itoa(num))
	n := len(bytes)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = int(bytes[i] - '0')
	}
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	for i := 2; i <= n; i++ {
		if arr[i-2] != 0 && (arr[i-2]*10+arr[i-1]) < 26 {
			dp[i] = dp[i-1] + dp[i-2]
		} else {
			dp[i] = dp[i-1]
		}
	}
	return dp[n]
}

/**
方法2：数字求余
*/
func translateNum2(num int) int {
	a := 1
	b := 1
	x, y := num%10, num%10
	for num != 0 {
		num /= 10
		x = num % 10
		temp := 10*x + y
		c := a
		if 10 <= temp && temp <= 25 {
			c = a + b
		}
		b = a
		a = c
		y = x
	}
	return a
}

/**
Java版
class Solution {
    public int translateNum(int num) {
        char[] chars = String.valueOf(num).toCharArray();
        int n = chars.length;
        int[] arr = new int[n];

        for (int i = 0; i < chars.length; i++) {
            arr[i] = chars[i] - '0';
        }

        int[] dp = new int[n + 1];
        dp[0] = 1;
        dp[1] = 1;
        for (int i = 2; i <= n; i++) {
            dp[i] = arr[i - 2] != 0 && (arr[i - 2] * 10 + arr[i - 1]) < 26 ? dp[i - 1] + dp[i - 2] : dp[i - 1];
        }

        return dp[n];
    }
}

class Solution {
    public int translateNum(int num) {
        int a = 1, b = 1, x, y = num % 10;
        while (num != 0) {
            num /= 10;
            x = num % 10;
            int tmp = 10 * x + y;
            int c = (tmp >= 10 && tmp <= 25) ? a + b : a;
            b = a;
            a = c;
            y = x;
        }
        return a;
    }
}
*/
func main() {

}
