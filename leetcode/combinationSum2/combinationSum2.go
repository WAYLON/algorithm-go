package main

import "sort"

/**
给定一个数组candidates和一个目标数target，找出candidates中所有可以使数字和为target的组合。
candidates中的每个数字在每个组合中只能使用一次。

说明：
所有数字（包括目标数）都是正整数。
解集不能包含重复的组合。

示例1:
输入: candidates =[10,1,2,7,6,1,5], target =8,
所求解集为:
[
  [1, 7],
  [1, 2, 5],
  [2, 6],
  [1, 1, 6]
]

示例2:
输入: candidates =[2,5,2,1,2], target =5,
所求解集为:
[
 [1,2,2],
 [5]
]
*/
/**
方法1 回溯+剪枝
	从每一层的第 2 个结点开始，都不能再搜索产生同一层结点已经使用过的 candidate 里的元素。
	因candidates中的每个数字在每个组合中只能使用一次。故上次选择的数不能跟本次选择的数相同
	剪枝提速:根据上面画树形图的经验，如果 target 减去一个数得到负数，那么减去一个更大的树依然是负数，同样搜索不到结果。基于这个想法，我们可以对输入数组进行排序，添加相关逻辑达到进一步剪枝的目的；
	排列问题，讲究顺序（即 [2, 2, 3] 与 [2, 3, 2] 视为不同列表时），需要记录哪些数字已经使用过，此时用 used 数组；
	组合问题，不讲究顺序（即 [2, 2, 3] 与 [2, 3, 2] 视为相同列表时），需要按照某种顺序搜索，此时使用 begin 变量。

*/
func combinationSum2(candidates []int, target int) [][]int {
	var res [][]int
	sort.Ints(candidates)
	var dfs func(begin int, target int, temp []int)
	dfs = func(begin int, target int, temp []int) {
		tmp := make([]int, len(temp))
		copy(tmp, temp)
		if target < 0 {
			return
		} else if target == 0 {
			res = append(res, tmp)
		}
		for i := begin; i < len(candidates); i++ {
			// 重点理解这里剪枝，前提是候选数组已经有序，
			if target-candidates[i] < 0 {
				break
			}
			//剪枝 当且仅当此次选择的值跟上次选择的值不同的时候才可以继续选择  否则跳过当前元素
			if i > begin && candidates[i-1] == candidates[i] {
				continue
			}

			tmp = append(tmp, candidates[i])
			dfs(i+1, target-candidates[i], tmp)
			tmp = tmp[:len(tmp)-1]
		}
	}

	dfs(0, target, []int{})
	return res
}

/**
java版
class Solution {
    public List<List<Integer>> combinationSum2(int[] candidates, int target) {
        Arrays.sort(candidates);
        List<List<Integer>> res = new ArrayList<>();
        dfs(res, candidates, target, 0, new ArrayList<Integer>());
        return res;
    }

    private void dfs(List<List<Integer>> res, int[] candidates, int target, int begin, List<Integer> temp) {
        if (target < 0) {
            return;
        } else if (target == 0) {
            res.add(new ArrayList<>(temp));
            return;
        }
        for (int i = begin; i < candidates.length; i++) {
            // 重点理解这里剪枝，前提是候选数组已经有序，
            if (target - candidates[i] < 0) {
                break;
            }
            //剪枝 当且仅当此次选择的值跟上次选择的值不同的时候才可以继续选择  否则跳过当前元素
            if (i > begin && candidates[i - 1] == candidates[i]) {
                continue;
            }
            temp.add(candidates[i]);
            dfs(res, candidates, target - candidates[i], i, temp);
            temp.remove(temp.size() - 1);
        }

    }
}
*/
func main() {

}
