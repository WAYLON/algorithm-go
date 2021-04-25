package main

import (
	"sort"
)

/**
元素的 频数 是该元素在一个数组中出现的次数。
给你一个整数数组 nums 和一个整数 k 。在一步操作中，你可以选择 nums 的一个下标，并将该下标对应元素的值增加 1 。
执行最多 k 次操作后，返回数组中最高频元素的 最大可能频数 。

示例 1：
输入：nums = [1,2,4], k = 5
输出：3
解释：对第一个元素执行 3 次递增操作，对第二个元素执 2 次递增操作，此时 nums = [4,4,4] 。
4 是数组中最高频元素，频数是 3 。

示例 2：
输入：nums = [1,4,8,13], k = 5
输出：2
解释：存在多种最优解决方案：
- 对第一个元素执行 3 次递增操作，此时 nums = [4,4,8,13] 。4 是数组中最高频元素，频数是 2 。
- 对第二个元素执行 4 次递增操作，此时 nums = [1,8,8,13] 。8 是数组中最高频元素，频数是 2 。
- 对第三个元素执行 5 次递增操作，此时 nums = [1,4,13,13] 。13 是数组中最高频元素，频数是 2 。

示例 3：
输入：nums = [3,9,6], k = 2
输出：1

提示：
1 <= nums.length <= 105
1 <= nums[i] <= 105
1 <= k <= 105
*/

/**
 * 滑动窗口最右侧的元素位置 j 作为最高频元素，
 * 右指针 j 每向右移动一位，窗口内的增量为 右指针与前一位的差 * 窗口的大小减一 即 (nums[j] - nums[j - 1]) * (j - i)
 * 如果窗口内的值大于k，缩小左指针 i， 窗口内的值相应减小 最右侧与最左侧的差值 即 nums[j] - nums[i]
 */
func maxFrequency(nums []int, k int) int {
	sort.Ints(nums)
	n := len(nums)
	i := 0
	j := 1
	res := 1
	window := 0
	for j < n {
		window += (nums[j] - nums[j-1]) * (j - i)
		if window <= k {
			res = max(res, j-i+1)
		}
		for window > k {
			window -= nums[j] - nums[i]
			i++
		}
		j++
	}
	return res
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

/**
class Solution {
    public int maxFrequency(int[] nums, int k) {
        Arrays.sort(nums);
        int n = nums.length;
        int i = 0;
        int j = 1;
        int window = 0;
        int res = 1;
        // [i, j)
        while (j < n) {
            window += (nums[j] - nums[j - 1]) * (j - i);
            if (window <= k) {
                res = Math.max(res, j - i + 1);
            }
            while (window > k) {
                window -= nums[j] - nums[i];
                i++;
            }
            j++;
        }
        return res;
    }
}
*/
func main() {

}
