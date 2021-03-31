package main

import "sort"

/**
给你一个整数数组 nums ，其中可能包含重复元素，请你返回该数组所有可能的子集（幂集）
解集 不能 包含重复的子集。返回的解集中，子集可以按 任意顺序 排列。

示例 1：
输入：nums = [1,2,2]
输出：[[],[1],[1,2],[1,2,2],[2],[2,2]]

示例 2：
输入：nums = [0]
输出：[[],[0]]

提示：
1 <= nums.length <= 10
-10 <= nums[i] <= 10
*/
/**
回溯 方法1
*/
func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	var ret [][]int
	sl := make([]int, len(nums))
	dfs(nums, &ret, sl, 0, 0)
	return ret
}

func dfs(nums []int, ret *[][]int, sl []int, start int, layer int) {
	dst := make([]int, layer)
	copy(dst, sl)
	*ret = append(*ret, dst)
	for i := start; i < len(nums); i++ {
		if i > start && nums[i] == nums[i-1] {
			continue
		}
		sl[layer] = nums[i]
		dfs(nums, ret, sl, i+1, layer+1)
	}
}

/**
回溯 方法2
*/
func subsetsWithDup2(nums []int) [][]int {
	var res [][]int
	//排序
	sort.Ints(nums)

	var backTrack func(i int, temp []int)
	backTrack = func(i int, temp []int) {
		tmp := make([]int, len(temp))
		copy(tmp, temp)
		res = append(res, tmp)
		for j := i; j < len(nums); j++ {
			//剪枝 当且仅当此次选择的值跟上次选择的值不同的时候才可以继续选择  否则跳过当前元素
			if j > i && nums[j-1] == nums[j] {
				continue
			}
			tmp = append(tmp, nums[j])
			backTrack(j+1, tmp)
			tmp = tmp[:len(tmp)-1]
		}

	}
	backTrack(0, []int{})

	return res
}

/**
java版

class Solution {
    public List<List<Integer>> subsetsWithDup(int[] nums) {
        List<List<Integer>> res = new ArrayList<>();
        Arrays.sort(nums);
        dfs(nums, res, 0, new ArrayList<>());
        return res;
    }

    private void dfs(int[] nums, List<List<Integer>> res, int i, List<Integer> temp) {
        res.add(new ArrayList<>(temp));
        for (int j = i; j < nums.length; j++) {
            if (j > i && nums[j - 1] == nums[j]) continue;
            temp.add(nums[j]);
            dfs(nums, res, j + 1, temp);
            temp.remove(temp.size() - 1);
        }
    }
}
*/
func main() {
	subsetsWithDup2([]int{2, 2, 1})
}
