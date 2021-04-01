package main

import "sort"

/**
给定一个可包含重复数字的序列 nums ，按任意顺序 返回所有不重复的全排列。

示例 1：
输入：nums = [1,1,2]
输出：
[[1,1,2],
 [1,2,1],
 [2,1,1]]

示例 2：
输入：nums = [1,2,3]
输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
*/
func permuteUnique(nums []int) [][]int {
	var res [][]int
	sort.Ints(nums)
	n := len(nums)
	used := make([]bool, n)
	var temp []int
	var dfs func(depth int)
	dfs = func(depth int) {
		if depth == n {
			tmp := make([]int, n)
			copy(tmp, temp)
			res = append(res, tmp)
		}

		for i := 0; i < n; i++ {
			//剪枝条件 当前元素没有被使用过并且（当前选择的数跟上次选择的数不一样|| 当前选择的数跟上次选择的数一样 但是上次选择的数已经使用了表明以后选择不到这个数了）
			if !used[i] && (i == 0 || nums[i-1] != nums[i] || used[i-1]) {
				temp = append(temp, nums[i])
				used[i] = true
				dfs(depth + 1)
				temp = temp[:len(temp)-1]
				used[i] = false
			}
		}

	}

	dfs(0)
	return res
}

/*
java版
class Solution {
    public List<List<Integer>> permuteUnique(int[] nums) {
        int len = nums.length;
        Arrays.sort(nums);
        List<List<Integer>> res = new ArrayList<>();
        boolean[] used = new boolean[len];
        List<Integer> temp = new ArrayList<>();
        dfs(res, len, 0, nums, temp, used);
        return res;
    }
*/
/**
 * @param res   动态数组保存的全排列
 * @param len   数组的长度
 * @param depth 深度
 * @param nums  数组
 * @param temp  路径
 * @param used  状态变量
 */ /*
    private void dfs(List<List<Integer>> res, int len, int depth, int[] nums, List<Integer> temp, boolean[] used) {
        if (depth == len) {
            res.add(new ArrayList<>(temp));
            return;
        }
        for (int i = 0; i < len; i++) {
            //剪枝条件 当前元素没有被使用过并且（当前选择的数跟上次选择的数不一样|| 当前选择的数跟上次选择的数一样 但是上次选择的数已经使用了表明以后选择不到这个数了）
            if (!used[i] && (i == 0 || nums[i - 1] != nums[i] || used[i - 1])) {
                temp.add(nums[i]);
                used[i] = true;
                dfs(res, len, depth + 1, nums, temp, used);
                temp.remove(temp.size() - 1);
                used[i] = false;
            }
        }
    }
}*/

func main() {

}
