package main

/**
实现strStr()函数。
给你两个字符串haystack 和 needle ，请你在 haystack 字符串中找出 needle 字符串出现的第一个位置（下标从 0 开始）。如果不存在，则返回 -1 。

说明：
当needle是空字符串时，我们应当返回什么值呢？这是一个在面试中很好的问题。
对于本题而言，当needle是空字符串时我们应当返回 0 。这与 C 语言的strstr()以及 Java 的indexOf()定义相符。

示例 1：
输入：haystack = "hello", needle = "ll"
输出：2

示例 2：
输入：haystack = "aaaaa", needle = "bba"
输出：-1

示例 3：
输入：haystack = "", needle = ""
输出：0

提示：
0 <= haystack.length, needle.length <= 5 * 104
haystack 和 needle 仅由小写英文字符组成
*/

/**
朴素解法
*/
func strStr(haystack string, needle string) int {
	n := len(haystack)
	m := len(needle)
	// 枚举原串的「发起点」
	for i := 0; i <= n-m; i++ {
		// 从原串的「发起点」和匹配串的「首位」开始，尝试匹配
		a := i
		b := 0
		for b < m && haystack[a] == needle[b] {
			a++
			b++
		}
		// 如果能够完全匹配，返回原串的「发起点」下标
		if b == m {
			return i
		}
	}
	return -1
}

/**
KMP
*/
func strStr2(ss string, pp string) int {
	if len(pp) == 0 {
		return 0
	}

	// 分别读取原串和匹配串的长度
	n := len(ss)
	m := len(pp)
	// 原串和匹配串前面都加空格，使其下标从 1 开始
	ss = " " + ss
	pp = " " + pp

	// 构建 next 数组，数组长度为匹配串的长度（next 数组是和匹配串相关的）
	next := make([]int, m+1)
	// 构造过程 i = 2，j = 0 开始，i 小于等于匹配串长度 【构造 i 从 2 开始】
	j := 0
	for i := 2; i <= m; i++ {
		// 匹配不成功的话，j = next(j)
		for j > 0 && pp[i] != pp[j+1] {
			j = next[j]
		}
		// 匹配成功的话，先让 j++
		if pp[i] == pp[j+1] {
			j++
		}
		// 更新 next[i]，结束本次循环，i++
		next[i] = j
	}

	// 匹配过程，i = 1，j = 0 开始，i 小于等于原串长度 【匹配 i 从 1 开始】
	j = 0
	for i := 1; i <= n; i++ {
		// 匹配不成功 j = next(j)
		for j > 0 && ss[i] != pp[j+1] {
			j = next[j]
		}
		// 匹配成功的话，先让 j++，结束本次循环后 i++
		if ss[i] == pp[j+1] {
			j++
		}
		// 整一段匹配成功，直接返回下标
		if j == m {
			return i - m
		}
	}

	return -1
}

/**
朴素解法
class Solution {
    public int strStr(String ss, String pp) {
        int n = ss.length(), m = pp.length();
        char[] s = ss.toCharArray(), p = pp.toCharArray();
        // 枚举原串的「发起点」
        for (int i = 0; i <= n - m; i++) {
            // 从原串的「发起点」和匹配串的「首位」开始，尝试匹配
            int a = i, b = 0;
            while (b < m && s[a] == p[b]) {
                a++;
                b++;
            }
            // 如果能够完全匹配，返回原串的「发起点」下标
            if (b == m) return i;
        }
        return -1;
    }
}

class Solution {
    // KMP 算法
    // ss: 原串(string)  pp: 匹配串(pattern)
    public int strStr(String ss, String pp) {
        if (pp.isEmpty()) return 0;

        // 分别读取原串和匹配串的长度
        int n = ss.length(), m = pp.length();
        // 原串和匹配串前面都加空格，使其下标从 1 开始
        ss = " " + ss;
        pp = " " + pp;

        char[] s = ss.toCharArray();
        char[] p = pp.toCharArray();

        // 构建 next 数组，数组长度为匹配串的长度（next 数组是和匹配串相关的）
        int[] next = new int[m + 1];
        // 构造过程 i = 2，j = 0 开始，i 小于等于匹配串长度 【构造 i 从 2 开始】
        for (int i = 2, j = 0; i <= m; i++) {
            // 匹配不成功的话，j = next(j)
            while (j > 0 && p[i] != p[j + 1]) j = next[j];
            // 匹配成功的话，先让 j++
            if (p[i] == p[j + 1]) j++;
            // 更新 next[i]，结束本次循环，i++
            next[i] = j;
        }

        // 匹配过程，i = 1，j = 0 开始，i 小于等于原串长度 【匹配 i 从 1 开始】
        for (int i = 1, j = 0; i <= n; i++) {
            // 匹配不成功 j = next(j)
            while (j > 0 && s[i] != p[j + 1]) j = next[j];
            // 匹配成功的话，先让 j++，结束本次循环后 i++
            if (s[i] == p[j + 1]) j++;
            // 整一段匹配成功，直接返回下标
            if (j == m) return i - m;
        }

        return -1;
    }
}
*/
func main() {

}
