package main

/**
给定一个非空且只包含非负数的整数数组nums，数组的度的定义是指数组里任一元素出现频数的最大值。
你的任务是在 nums 中找到与nums拥有相同大小的度的最短连续子数组，返回其长度。

示例 1：
输入：[1, 2, 2, 3, 1]
输出：2
解释：
输入数组的度是2，因为元素1和2的出现频数最大，均为2.
连续子数组里面拥有相同度的有如下所示:
[1, 2, 2, 3, 1], [1, 2, 2, 3], [2, 2, 3, 1], [1, 2, 2], [2, 2, 3], [2, 2]
最短连续子数组[2, 2]的长度为2，所以返回2.

示例 2：
输入：[1,2,2,3,1,4,2]
输出：6

提示：
nums.length在1到 50,000 区间范围内。
nums[i]是一个在 0 到 49,999 范围内的整数。

*/
type entry struct{ cnt, l, r int }

func findShortestSubArray(nums []int) (ans int) {
	mp := map[int]entry{}
	for i, v := range nums {
		if e, has := mp[v]; has {
			e.cnt++
			e.r = i
			mp[v] = e
		} else {
			mp[v] = entry{1, i, i}
		}
	}
	maxCnt := 0
	for _, e := range mp {
		if e.cnt > maxCnt {
			maxCnt, ans = e.cnt, e.r-e.l+1
		} else if e.cnt == maxCnt {
			ans = min(ans, e.r-e.l+1)
		}
	}
	return
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {

}
