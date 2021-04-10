package main

import "math"

/**
给定两个整数，被除数dividend和除数divisor。将两数相除，要求不使用乘法、除法和 mod 运算符。
返回被除数dividend除以除数divisor得到的商。
整数除法的结果应当截去（truncate）其小数部分，例如：truncate(8.345) = 8 以及 truncate(-2.7335) = -2

示例1:
输入: dividend = 10, divisor = 3
输出: 3
解释: 10/3 = truncate(3.33333..) = truncate(3) = 3

示例2:
输入: dividend = 7, divisor = -3
输出: -2
解释: 7/-3 = truncate(-2.33333..) = -2

提示：
被除数和除数均为 32 位有符号整数。
除数不为0。
假设我们的环境只能存储 32 位有符号整数，其数值范围是 [−231, 231− 1]。本题中，如果除法结果溢出，则返回 231− 1。
*/

/**
二分法 超出时间限制
*/
func divide(dividend int, divisor int) int {
	x := int64(dividend)
	y := int64(divisor)
	if x == 0 {
		return 0
	}
	negative := (x ^ y) < 0 //用异或来计算是否符号相异

	if x < 0 {
		x = -x
	}
	if y < 0 {
		y = -y
	}

	var l int64 = 0
	var r = x

	for l < r {
		mid := l + r + 1>>1
		if quickMul(mid, y) <= x {
			l = mid
		} else {
			r = mid - 1
		}
	}
	ans := l
	if negative { //符号相异取反
		ans = -l
	}

	if ans > math.MaxInt32 || ans < math.MinInt32 {
		return math.MaxInt32
	}
	return int(ans)
}

func quickMul(a int64, b int64) int64 {
	var res int64
	if b < 0 {
		a = -a
		b = -b
	}

	for b > 0 {
		if b&1 == 1 {
			res += a
		}
		a += a
		b >>= 1
	}
	return res
}

/**
 * 解题思路：这题是除法，所以先普及下除法术语
 * 商，公式是：(被除数-余数)÷除数=商，记作：被除数÷除数=商...余数，是一种数学术语。
 * 在一个除法算式里，被除数、余数、除数和商的关系为：(被除数-余数)÷除数=商，记作：被除数÷除数=商...余数，
 * 进而推导得出：商×除数+余数=被除数。
 *
 * 要求商，我们首先想到的是减法，能被减多少次，那么商就为多少，但是明显减法的效率太低
 *
 * 那么我们可以用位移法，因为计算机在做位移时效率特别高，向左移1相当于乘以2，向右位移1相当于除以2
 *
 * 我们可以把一个dividend（被除数）先除以2^n，n最初为31，不断减小n去试探,当某个n满足dividend/2^n>=divisor时，
 *
 * 表示我们找到了一个足够大的数，这个数*divisor是不大于dividend的，所以我们就可以减去2^n个divisor，以此类推
 *
 * 我们可以以100/3为例
 *
 * 2^n是1，2，4，8...2^31这种数，当n为31时，这个数特别大，100/2^n是一个很小的数，肯定是小于3的，所以循环下来，
 *
 * 当n=5时，100/32=3, 刚好是大于等于3的，这时我们将100-32*3=4，也就是减去了32个3，接下来我们再处理4，同样手法可以再减去一个3
 *
 * 所以一共是减去了33个3，所以商就是33
 *
 * 这其中得处理一些特殊的数，比如divisor是不能为0的，Integer.MIN_VALUE和Integer.MAX_VALUE
 *
 */
func divide2(dividend int, divisor int) int {
	if dividend == 0 {
		return 0
	}
	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}
	negative := (dividend ^ divisor) < 0 //用异或来计算是否符号相异
	if dividend < 0 {
		dividend = -dividend
	}
	if divisor < 0 {
		divisor = -divisor
	}
	result := 0
	for i := 31; i >= 0; i-- {
		if (dividend >> i) >= divisor { //找出足够大的数2^n*divisor
			result += 1 << i         //将结果加上2^n
			dividend -= divisor << i //将被除数减去2^n*divisor
		}
	}
	if negative { //符号相异取反
		return -result
	}
	return result
}

/**
java版
class Solution {
    public int divide(int a, int b) {
        long x = a, y = b;
        boolean isNeg = false;
        if ((x > 0 && y < 0) || (x < 0 && y > 0)) isNeg = true;
        if (x < 0) x = -x;
        if (y < 0) y = -y;
        long l = 0, r = x;
        while (l < r) {
            long mid = l + r + 1 >> 1;
            if (mul(mid, y) <= x) {
                l = mid;
            } else {
                r = mid - 1;
            }
        }
        long ans = isNeg ? -l : l;
        if (ans > Integer.MAX_VALUE || ans < Integer.MIN_VALUE) return Integer.MAX_VALUE;
        return (int) ans;
    }

    long mul(long a, long k) {
        long ans = 0;
        while (k > 0) {
            if ((k & 1) == 1) ans += a;
            k >>= 1;
            a += a;
        }
        return ans;
    }
}

class Solution {
    public int divide(int dividend, int divisor) {
        if (dividend == 0) {
            return 0;
        }
        if (dividend == Integer.MIN_VALUE && divisor == -1) {
            return Integer.MAX_VALUE;
        }
        boolean negative;
        negative = (dividend ^ divisor) <0;//用异或来计算是否符号相异
        long t = Math.abs((long) dividend);
        long d= Math.abs((long) divisor);
        int result = 0;
        for (int i=31; i>=0;i--) {
            if ((t>>i)>=d) {//找出足够大的数2^n*divisor
                result+=1<<i;//将结果加上2^n
                t-=d<<i;//将被除数减去2^n*divisor
            }
        }
        return negative ? -result : result;//符号相异取反
    }
}
*/
func main() {

}
