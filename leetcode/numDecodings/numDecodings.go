package main

/**
一条包含字母A-Z 的消息通过以下映射进行了 编码 ：
'A' -> 1
'B' -> 2
...
'Z' -> 26
要 解码 已编码的消息，所有数字必须基于上述映射的方法，反向映射回字母（可能有多种方法）。例如，"11106" 可以映射为：
"AAJF" ，将消息分组为 (1 1 10 6)
"KJF" ，将消息分组为 (11 10 6)
注意，消息不能分组为 (1 11 06) ，因为 "06" 不能映射为 "F" ，这是由于 "6" 和 "06" 在映射中并不等价。
给你一个只含数字的 非空 字符串 s ，请计算并返回 解码 方法的 总数 。
题目数据保证答案肯定是一个 32 位 的整数。

示例 1：

输入：s = "12"
输出：2
解释：它可以解码为 "AB"（1 2）或者 "L"（12）。

示例 2：

输入：s = "226"
输出：3
解释：它可以解码为 "BZ" (2 26), "VF" (22 6), 或者 "BBF" (2 2 6) 。

示例 3：

输入：s = "0"
输出：0
解释：没有字符映射到以 0 开头的数字。
含有 0 的有效映射是 'J' -> "10" 和 'T'-> "20" 。
由于没有字符，因此没有有效的方法对此进行解码，因为所有数字都需要映射。

示例 4：

输入：s = "06"
输出：0
解释："06" 不能映射到 "F" ，因为字符串含有前导 0（"6" 和 "06" 在映射中并不等价）。

提示：
1 <= s.length <= 100
s 只包含数字，并且可能包含前导零。
*/

/**
状态转移方程
f[i]=f[i−1],1⩽a≤9
f[i]=f[i−2],10⩽b⩽26
f[i]=f[i−1]+f[i−2],1⩽a≤9,10⩽b⩽26
*/
func numDecodings(s string) int {
	n := len(s)
	dp := make([]int, n+1)
	s = " " + s
	dp[0] = 1
	for i := 1; i <= n; i++ {
		// a : 代表「当前位置」单独形成 item
		// b : 代表「当前位置」与「前一位置」共同形成 item
		a := s[i] - '0'
		b := (s[i-1]-'0')*10 + (s[i] - '0')
		// 如果 a 属于有效值，那么 f[i] 可以由 f[i - 1] 转移过来
		if 1 <= a && a <= 9 {
			dp[i] = dp[i-1]
		}
		// 如果 b 属于有效值，那么 f[i] 可以由 f[i - 2] 或者 f[i - 1] & f[i - 2] 转移过来
		if 10 <= b && b <= 26 {
			dp[i] += dp[i-2]
		}
	}
	return dp[n]
}

/**
空间优化
*/
func numDecodings2(s string) int {
	n := len(s)
	dp := make([]int, 3)
	s = " " + s
	dp[0] = 1
	for i := 1; i <= n; i++ {
		dp[i%3] = 0
		// a : 代表「当前位置」单独形成 item
		// b : 代表「当前位置」与「前一位置」共同形成 item
		a := s[i] - '0'
		b := (s[i-1]-'0')*10 + (s[i] - '0')
		// 如果 a 属于有效值，那么 f[i] 可以由 f[i - 1] 转移过来
		if 1 <= a && a <= 9 {
			dp[i%3] = dp[(i-1)%3]
		}
		// 如果 b 属于有效值，那么 f[i] 可以由 f[i - 2] 或者 f[i - 1] & f[i - 2] 转移过来
		if 10 <= b && b <= 26 {
			dp[i%3] += dp[(i-2)%3]
		}
	}
	return dp[n%3]
}

/**
Java版
class Solution {
    public int numDecodings(String s) {
        int n = s.length();
        s = " " + s;
        char[] cs = s.toCharArray();
        int[] f = new int[n + 1];
        f[0] = 1;
        for (int i = 1; i <= n; i++) {
            // a : 代表「当前位置」单独形成 item
            // b : 代表「当前位置」与「前一位置」共同形成 item
            int a = cs[i] - '0', b = (cs[i - 1] - '0') * 10 + (cs[i] - '0');
            // 如果 a 属于有效值，那么 f[i] 可以由 f[i - 1] 转移过来
            if (1 <= a && a <= 9) f[i] = f[i - 1];
            // 如果 b 属于有效值，那么 f[i] 可以由 f[i - 2] 或者 f[i - 1] & f[i - 2] 转移过来
            if (10 <= b && b <= 26) f[i] += f[i - 2];
        }
        return f[n];
    }
}

class Solution {
    public int numDecodings(String s) {
        int n = s.length();
        s = " " + s;
        char[] cs = s.toCharArray();
        int[] f = new int[3];
        f[0] = 1;
        for (int i = 1; i <= n; i++) {
            f[i % 3] = 0;
            int a = cs[i] - '0', b = (cs[i - 1] - '0') * 10 + (cs[i] - '0');
            if (1 <= a && a <= 9) f[i % 3] = f[(i - 1) % 3];
            if (10 <= b && b <= 26) f[i % 3] += f[(i - 2) % 3];
        }
        return f[n % 3];
    }
}

*/
func main() {

}
