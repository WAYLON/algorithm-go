package main

/**
给定一个非负整数num。对于0 ≤ i ≤ num 范围中的每个数字i，计算其二进制数中的 1 的数目并将它们作为数组返回。

示例 1:
输入: 2
输出: [0,1,1]

示例2:
输入: 5
输出: [0,1,1,2,1,2]
进阶:

给出时间复杂度为O(n*sizeof(integer))的解答非常容易。但你可以在线性时间O(n)内用一趟扫描做到吗？
要求算法的空间复杂度为O(n)。
你能进一步完善解法吗？要求在C++或任何其他语言中不使用任何内置函数（如 C++ 中的__builtin_popcount）来执行此操作。
*/

/**
方法一：i & (i - 1)可以去掉i最右边的一个1（如果有），
因此 i & (i - 1）是比 i 小的，而且i & (i - 1)的1的个数已经在前面算过了
，所以i的1的个数就是 i & (i - 1)的1的个数加上1
*/
func countBits(num int) []int {
	res := make([]int, num+1)
	for i := 1; i <= num; i++ {
		res[i] = res[i&(i-1)] + 1
	}
	return res
}

/**
方法二：i >> 1会把最低位去掉，因此i >> 1 也是比i小的，同样也是在前面的数组里算过。
当 i 的最低位是0，则 i 中1的个数和i >> 1中1的个数相同
；当i的最低位是1，i 中1的个数是 i >> 1中1的个数再加1
*/
func countBits2(num int) []int {
	res := make([]int, num+1)
	for i := 0; i <= num; i++ {
		res[i] = res[i>>1] + (i & 1)
	}
	return res
}

func main() {

}
