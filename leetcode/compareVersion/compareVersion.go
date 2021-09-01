package main

import (
	"strconv"
	"strings"
)

/**
给你两个版本号 version1 和 version2 ，请你比较它们。
版本号由一个或多个修订号组成，各修订号由一个 '.' 连接。每个修订号由 多位数字 组成，可能包含 前导零 。每个版本号至少包含一个字符。修订号从左到右编号，下标从 0 开始，最左边的修订号下标为 0 ，下一个修订号下标为 1 ，以此类推。例如，2.5.33 和 0.1 都是有效的版本号。
比较版本号时，请按从左到右的顺序依次比较它们的修订号。比较修订号时，只需比较 忽略任何前导零后的整数值 。也就是说，修订号 1 和修订号 001 相等 。如果版本号没有指定某个下标处的修订号，则该修订号视为 0 。例如，版本 1.0 小于版本 1.1 ，因为它们下标为 0 的修订号相同，而下标为 1 的修订号分别为 0 和 1 ，0 < 1 。
返回规则如下：
如果version1>version2返回1，
如果version1<version2 返回 -1，
除此之外返回 0。

示例 1：
输入：version1 = "1.01", version2 = "1.001"
输出：0
解释：忽略前导零，"01" 和 "001" 都表示相同的整数 "1"

示例 2：
输入：version1 = "1.0", version2 = "1.0.0"
输出：0
解释：version1 没有指定下标为 2 的修订号，即视为 "0"

示例 3：
输入：version1 = "0.1", version2 = "1.1"
输出：-1
解释：version1 中下标为 0 的修订号是 "0"，version2 中下标为 0 的修订号是 "1" 。0 < 1，所以 version1 < version2

示例 4：
输入：version1 = "1.0.1", version2 = "1"
输出：1

示例 5：
输入：version1 = "7.5.2.4", version2 = "7.5.3"
输出：-1

提示：
1 <= version1.length, version2.length <= 500
version1 和 version2 仅包含数字和 '.'
version1 和 version2 都是 有效版本号
version1 和 version2 的所有修订号都可以存储在 32 位整数 中
*/

/**
方法1：双指针
*/
func compareVersion(version1 string, version2 string) int {
	n := len(version1)
	m := len(version2)
	i := 0
	j := 0
	for i < n || j < m {
		a := 0
		for ; i < n && version1[i] != '.'; i++ {
			a = a*10 + int(version1[i]-'0')
		}

		//跳过.
		i++

		b := 0
		for ; j < m && version2[j] != '.'; j++ {
			b = b*10 + int(version2[j]-'0')
		}

		//跳过.
		j++

		if a != b {
			if a > b {
				return 1
			} else {
				return -1
			}
		}

	}

	return 0
}

/**
方法2：字符串模拟
*/
func compareVersion2(version1 string, version2 string) int {
	ss1 := strings.Split(version1, ".")
	ss2 := strings.Split(version2, ".")
	n := len(ss1)
	m := len(ss2)
	i, j := 0, 0
	for i < n || j < m {
		a, b := 0, 0
		if i < n {
			a, _ = strconv.Atoi(ss1[i])
			i++
		}
		if j < m {
			b, _ = strconv.Atoi(ss2[j])
			j++
		}

		if a != b {
			if a > b {
				return 1
			} else {
				return -1
			}
		}
	}
	return 0
}

/**
class Solution {
    public int compareVersion(String version1, String version2) {
        int n = version1.length(), m = version2.length();
        int i = 0, j = 0;
        while (i < n || j < m) {
            int x = 0;
            for (; i < n && version1.charAt(i) != '.'; ++i) {
                x = x * 10 + version1.charAt(i) - '0';
            }
            ++i; // 跳过点号
            int y = 0;
            for (; j < m && version2.charAt(j) != '.'; ++j) {
                y = y * 10 + version2.charAt(j) - '0';
            }
            ++j; // 跳过点号
            if (x != y) {
                return x > y ? 1 : -1;
            }
        }
        return 0;
    }
}

class Solution {
    public int compareVersion(String v1, String v2) {
        String[] ss1 = v1.split("\\."), ss2 = v2.split("\\.");
        int n = ss1.length, m = ss2.length;
        int i = 0, j = 0;
        while (i < n || j < m) {
            int a = 0, b = 0;
            if (i < n) a = Integer.parseInt(ss1[i++]);
            if (j < m) b = Integer.parseInt(ss2[j++]);
            if (a != b) return a > b ? 1 : -1;
        }
        return 0;
    }
}
*/
func main() {
	compareVersion2("1.0001.1", "1.001.1")
}
