package main

import "sort"

/**
给你一个由 无重复 正整数组成的集合 nums ，请你找出并返回其中最大的整除子集 answer ，子集中每一元素对 (answer[i], answer[j]) 都应当满足：
answer[i] % answer[j] == 0 ，或
answer[j] % answer[i] == 0
如果存在多个有效解子集，返回其中任何一个均可。

示例 1：
输入：nums = [1,2,3]
输出：[1,2]
解释：[1,3] 也会被视为正确答案。

示例 2：
输入：nums = [1,2,4,8]
输出：[1,2,4,8]

提示：
1 <= nums.length <= 1000
1 <= nums[i] <= 2 * 109
nums 中的所有整数 互不相同
*/
func largestDivisibleSubset(nums []int) []int {
	sort.Ints(nums)
	n := len(nums)
	//定义 f[i] 为考虑前 i 个数字，且以第 i 个数为结尾的最长「整除子集」长度。
	f := make([]int, n)
	//定义 g[i] 为记录 f[i] 是由哪个下标的状态转移而来，如果 f[i] = f[j] + 1, 则有 g[i] = j。
	g := make([]int, n)

	for i := 0; i < n; i++ {
		// 至少包含自身一个数，因此起始长度为 1，由自身转移而来
		l := 1
		prev := i
		for j := 0; j < i; j++ {
			if nums[i]%nums[j] == 0 {
				// 如果能接在更长的序列后面，则更新「最大长度」&「从何转移而来」
				if f[j]+1 > l {
					l = f[j] + 1
					prev = j
				}
			}
		}

		// 记录「最终长度」&「从何转移而来」
		f[i] = l
		g[i] = prev
	}

	// 遍历所有的 f[i]，取得「最大长度」和「对应下标」
	max := -1
	idx := -1
	for i := 0; i < n; i++ {
		if f[i] > max {
			idx = i
			max = f[i]
		}
	}

	// 使用 g[] 数组回溯出具体方案
	var ans []int
	for len(ans) != max {
		ans = append(ans, nums[idx])

		idx = g[idx]
	}
	return ans
}

/**
class Solution {
    public List<Integer> largestDivisibleSubset(int[] nums) {
        Arrays.sort(nums);

        int n = nums.length;
        //定义 f[i] 为考虑前 i 个数字，且以第 i 个数为结尾的最长「整除子集」长度。
        int[] f = new int[n];
        //定义 g[i] 为记录 f[i] 是由哪个下标的状态转移而来，如果 f[i] = f[j] + 1, 则有 g[i] = j。
        int[] g = new int[n];


        for (int i = 0; i < n; i++) {
            // 至少包含自身一个数，因此起始长度为 1，由自身转移而来
            int len = 1, prev = i;
            for (int j = 0; j < i; j++) {
                if (nums[i] % nums[j] == 0) {
                    // 如果能接在更长的序列后面，则更新「最大长度」&「从何转移而来」
                    if (f[j] + 1 > len) {
                        len = f[j] + 1;
                        prev = j;
                    }
                }
            }
            // 记录「最终长度」&「从何转移而来」
            f[i] = len;
            g[i] = prev;
        }

        // 遍历所有的 f[i]，取得「最大长度」和「对应下标」
        int max = -1, idx = -1;
        for (int i = 0; i < n; i++) {
            if (f[i] > max) {
                idx = i;
                max = f[i];
            }
        }

        // 使用 g[] 数组回溯出具体方案
        List<Integer> ans = new ArrayList<>();
        while (ans.size() != max) {
            ans.add(nums[idx]);
            idx = g[idx];
        }
        return ans;
    }
}
*/
func main() {
	largestDivisibleSubset([]int{1, 2, 3})
}
