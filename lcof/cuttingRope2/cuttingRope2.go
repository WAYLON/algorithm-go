package main

/**
给你一根长度为 n 的绳子，请把绳子剪成整数长度的 m 段（m、n都是整数，n>1并且m>1），每段绳子的长度记为 k[0],k[1]...k[m-1] 。
请问 k[0]*k[1]*...*k[m-1] 可能的最大乘积是多少？例如，当绳子的长度是8时，我们把它剪成长度分别为2、3、3的三段，此时得到的最大乘积是18。

示例 1：
输入: 2
输出: 1
解释: 2 = 1 + 1, 1 × 1 = 1

示例2:
输入: 10
输出: 36
解释: 10 = 3 + 3 + 4, 3 ×3 ×4 = 36
提示：

2 <= n <= 58
*/

/**
数学推导法 即
当 b = 0 b=0 时，直接返回 3^a3
当 b = 1 b=1 时，要将一个 1 + 3 转换为 2+2，因此返回 3^{a-1}×4；
当 b = 2 b=2 时，返回 3^a×2。

快速幂
*/
func cuttingRope2(n int) int {
	if n <= 3 {
		return n - 1
	}

	b := n % 3
	rem := int64(1)
	x := int64(3)
	p := int64(1000000007)
	for a := n/3 - 1; a > 0; a >>= 1 {
		if (a & 1) == 1 {
			rem = (x * rem) % p
		}
		x = (x * x) % p
	}

	if b == 0 {
		return (int)(rem * 3 % p)
	} else if b == 1 {
		return (int)(rem * 4 % p)
	}

	return (int)(rem * 6 % p)
}

func main() {
	cuttingRope2(1000)
}
