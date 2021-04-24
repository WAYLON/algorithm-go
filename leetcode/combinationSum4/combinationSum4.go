package main

/**
给你一个由 不同 整数组成的数组 nums ，和一个目标整数 target 。请你从 nums 中找出并返回总和为 target 的元素组合的个数。
题目数据保证答案符合 32 位整数范围。


示例 1：
输入：nums = [1,2,3], target = 4
输出：7
解释：
所有可能的组合为：
(1, 1, 1, 1)
(1, 1, 2)
(1, 2, 1)
(1, 3)
(2, 1, 1)
(2, 2)
(3, 1)
请注意，顺序不同的序列被视作不同的组合。

示例 2：
输入：nums = [9], target = 3
输出：0

提示：
1 <= nums.length <= 200
1 <= nums[i] <= 1000
nums 中的所有元素 互不相同
1 <= target <= 1000

进阶：如果给定的数组中含有负数会发生什么？问题会产生何种变化？如果允许负数出现，需要向题目中添加哪些限制条件？
*/

/**
最后一个数选择 nums[0]，方案数为 f[len - 1][target - nums[0]]
最后一个数选择 nums[1]，方案数为 f[len - 1][target - nums[1]]
最后一个数选择 nums[2]，方案数为 f[len - 1][target - nums[2]]
定义 f[i][j] 为组合长度为 i，凑成总和为 j 的方案数是多少。
*/
func combinationSum4(nums []int, target int) int {
	dp := make([][]int, target+1)
	for i := range dp {
		dp[i] = make([]int, target+1)
	}
	dp[0][0] = 1
	ans := 0
	for i := 1; i <= target; i++ {
		for j := 0; j <= target; j++ {
			for _, v := range nums {
				if v <= j {
					dp[i][j] += dp[i-1][j-v]
				}
			}
		}
		ans += dp[i][target]
	}
	return ans
}

//dp[i]=dp[i-nums[0]]+dp[i-nums[1]]+dp[i-nums[2]]+...
func combinationSum42(nums []int, target int) int {
	dp := make([]int, target+1)
	dp[0] = 1
	for i := 1; i <= target; i++ {
		for _, v := range nums {
			if v <= i {
				dp[i] += dp[i-v]
			}
		}
	}
	return dp[target]
}

/**
class Solution {
    public int combinationSum4(int[] nums, int target) {
        //dfs会超时
        //使用dp数组，dp[i]代表组合数为i时使用nums中的数能组成的组合数的个数
        //别怪我写的这么完整
        //dp[i]=dp[i-nums[0]]+dp[i-nums[1]]+dp[i-nums[2]]+...
        //举个例子比如nums=[1,3,4],target=7;
        //dp[7]=dp[6]+dp[4]+dp[3]
        //其实就是说7的组合数可以由三部分组成，1和dp[6]，3和dp[4],4和dp[3];
        int[] dp = new int[target + 1];
        //是为了算上自己的情况，比如dp[1]可以由dp【0】和1这个数的这种情况组成。
        dp[0] = 1;
        for (int i = 1; i <= target; i++) {
            for (int num : nums) {
                if (i >= num) {
                    dp[i] += dp[i - num];
                }
            }
        }
        return dp[target];
    }
}

class Solution {
    public int combinationSum4(int[] nums, int t) {
        // 因为 nums[i] 最小值为 1，因此构成答案的最大长度为 target
        //定义 f[i][j] 为组合长度为 i，凑成总和为 j 的方案数是多少。
        int len = t;
        int[][] f = new int[len + 1][t + 1];
        f[0][0] = 1;
        int ans = 0;
        for (int i = 1; i <= len; i++) {
            for (int j = 0; j <= t; j++) {
                for (int u : nums) {
                    if (j >= u) f[i][j] += f[i - 1][j - u];
                }
            }
            ans += f[i][t];
        }
        return ans;
    }
}
*/
func main() {

}
