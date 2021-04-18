package main

/**
给你一个 有序数组nums，它由n个非负整数组成，同时给你一个整数maximumBit。你需要执行以下查询 n次：
找到一个非负整数k < 2maximumBit，使得nums[0] XOR nums[1] XOR ... XOR nums[nums.length-1] XOR k的结果 最大化。k是第 i个查询的答案。
从当前数组nums删除最后一个元素。
请你返回一个数组answer，其中answer[i]是第i个查询的结果。

示例 1：
输入：nums = [0,1,1,3], maximumBit = 2
输出：[0,3,2,3]
解释：查询的答案如下：
第一个查询：nums = [0,1,1,3]，k = 0，因为 0 XOR 1 XOR 1 XOR 3 XOR 0 = 3 。
第二个查询：nums = [0,1,1]，k = 3，因为 0 XOR 1 XOR 1 XOR 3 = 3 。
第三个查询：nums = [0,1]，k = 2，因为 0 XOR 1 XOR 2 = 3 。
第四个查询：nums = [0]，k = 3，因为 0 XOR 3 = 3 。

示例 2：
输入：nums = [2,3,4,7], maximumBit = 3
输出：[5,2,6,5]
解释：查询的答案如下：
第一个查询：nums = [2,3,4,7]，k = 5，因为 2 XOR 3 XOR 4 XOR 7 XOR 5 = 7。
第二个查询：nums = [2,3,4]，k = 2，因为 2 XOR 3 XOR 4 XOR 2 = 7 。
第三个查询：nums = [2,3]，k = 6，因为 2 XOR 3 XOR 6 = 7 。
第四个查询：nums = [2]，k = 5，因为 2 XOR 5 = 7 。

示例 3：
输入：nums = [0,1,2,2,5,7], maximumBit = 3
输出：[4,3,6,4,6,7]

提示：
nums.length == n
1 <= n <= 105
1 <= maximumBit <= 20
0 <= nums[i] < 2maximumBit
nums 中的数字已经按升序排好序。
*/
func getMaximumXor(nums []int, maximumBit int) []int {
	n := len(nums)
	res := make([]int, n)
	xor := 0
	max := (1 << maximumBit) - 1
	j := n - 1
	for i := 0; i < n; i++ {
		xor ^= nums[i]
		res[j] = max - xor
		j--
	}
	return res
}

/**
Java版
class Solution {
    public int[] getMaximumXor(int[] nums, int maximumBit) {
        int n = nums.length;
        int[] res = new int[n];
        int xor = 0;
        int max = (1 << maximumBit) - 1;
        for (int i = 0, j = n - 1; i < n; i++, j--) {
            xor ^= nums[i];
            res[j] = max - xor;
        }
        return res;
    }
}
*/
func main() {

}
