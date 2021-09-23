package main

/**
给定一个整数，写一个函数来判断它是否是 3的幂次方。如果是，返回 true ；否则，返回 false 。
整数 n 是 3 的幂次方需满足：存在整数 x 使得 n == 3x

示例 1：
输入：n = 27
输出：true

示例 2：
输入：n = 0
输出：false

示例 3：
输入：n = 9
输出：true

示例 4：
输入：n = 45
输出：false

提示：
-231 <= n <= 231 - 1

进阶：
你能不使用循环或者递归来完成本题吗？

*/
func isPowerOfThree(n int) bool {
	if n <= 0 {
		return false
	}
	for n%3 == 0 {
		n /= 3
	}
	return n == 1
}

/**
1162261467 为3的19次幂 其他3次方整数一定是他的公约数 附3是质数所以可行  即int 范围即 2^31-1 ~ -2^31 即：2147483647 ~ -2147483648
*/
func isPowerOfThree1(n int) bool {
	return n > 0 && 1162261467%n == 0
}

/*var str = ",1,"

func isPowerOfThree2(n int) bool {
	cur := 1
	for cur <= math.MaxInt32/3 {
		cur *= 3
		str = str + "," + strconv.Itoa(cur) + ","
	}

	return n > 0 && strings.Contains(str, ","+strconv.Itoa(n)+",")
}*/

/** java

class Solution {
    public boolean isPowerOfThree(int n) {
        if (n <= 0) return false;
        while (n % 3 == 0) n /= 3;
        return n == 1;
    }
}

class Solution {
    public boolean isPowerOfThree(int n) {
        return n > 0 && 1162261467 % n == 0;
    }
}

class Solution {
    static Set<Integer> set = new HashSet<>();
    static {
        int cur = 1;
        set.add(cur);
        while (cur <= Integer.MAX_VALUE / 3) {
            cur *= 3;
            set.add(cur);
        }
    }
    public boolean isPowerOfThree(int n) {
        return n > 0 && set.contains(n);
    }
}

*/
func main() {

}
