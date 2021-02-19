package main

/**
给定一个由若干 0 和 1 组成的数组A，我们最多可以将K个值从 0 变成 1 。
返回仅包含 1 的最长（连续）子数组的长度。

示例 1：
输入：A = [1,1,1,0,0,0,1,1,1,1,0], K = 2
输出：6
解释：
[1,1,1,0,0,1,1,1,1,1,1]
粗体数字从 0 翻转到 1，最长的子数组长度为 6。

示例 2：
输入：A = [0,0,1,1,0,0,1,1,1,0,1,1,0,0,0,1,1,1,1], K = 3
输出：10
解释：
[0,0,1,1,1,1,1,1,1,1,1,1,0,0,0,1,1,1,1]
粗体数字从 0 翻转到 1，最长的子数组长度为 10。

提示：
1 <= A.length <= 20000
0 <= K <= A.length
A[i] 为0或1
	滑动窗口
*/

func longestOnes(A []int, K int) int {
	N := len(A)
	res := 0
	left, right := 0, 0
	zeros := 0
	for right < N {
		if A[right] == 0 {
			zeros++
		}
		for zeros > K {
			if A[left] == 0 {
				zeros--
			}
			left++
		}
		temp := right - left + 1
		if res < temp {
			res = right - left + 1
		}
		right++
	}
	return res
}

func main() {

}
