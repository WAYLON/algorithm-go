package main

/**
给定一个 没有重复 数字的序列，返回其所有可能的全排列。

示例:
输入: [1,2,3]
输出:
[
  [1,2,3],
  [1,3,2],
  [2,1,3],
  [2,3,1],
  [3,1,2],
  [3,2,1]
]
从 [1, 2, 3] 到 [1, 3, 2] ，深度优先遍历是这样做的，从 [1, 2, 3] 回到 [1, 2] 的时候，
需要撤销刚刚已经选择的数 3，因为在这一层只有一个数 3 我们已经尝试过了，因此程序回到上一层，
需要撤销对 2 的选择，好让后面的程序知道，选择 3 了以后还能够选择 2。
执行深度优先遍历，从较深层的结点返回到较浅层结点的时候，需要做「状态重置」，即「回到过去」
*/
func permute(nums []int) [][]int {
	var res [][]int
	n := len(nums)
	used := make([]bool, n)
	temp := []int{}
	var dfs func(depth int)
	dfs = func(depth int) {
		if depth == n {
			tmp := make([]int, n)
			copy(tmp, temp)
			res = append(res, tmp)
		}
		for i := 0; i < n; i++ {
			if !used[i] {
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

/**
java版
class Solution {
    public List<List<Integer>> permute(int[] nums) {
        int len = nums.length;
        List<List<Integer>> res = new ArrayList<>();
        boolean[] used = new boolean[len];
        List<Integer> temp = new ArrayList<>();
        dfs(res, len, 0, nums, temp, used);
        return res;
    }

*/
/**
 * @param res  动态数组保存的全排列
 * @param len  数组的长度
 * @param depth  深度
 * @param nums 数组
 * @param temp 路径
 * @param used 状态变量
 */ /*

    private void dfs(List<List<Integer>> res, int len, int depth, int[] nums, List<Integer> temp, boolean[] used) {
        if (depth == len) {
            res.add(new ArrayList<>(temp));
            return;
        }
        for (int i = 0; i < len; i++) {
            if (!used[i]) {
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
