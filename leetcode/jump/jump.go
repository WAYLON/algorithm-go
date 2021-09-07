package main

/**
给你一个非负整数数组nums ，你最初位于数组的第一个位置。
数组中的每个元素代表你在该位置可以跳跃的最大长度。
你的目标是使用最少的跳跃次数到达数组的最后一个位置。
假设你总是可以到达数组的最后一个位置。

示例 1:
输入: nums = [2,3,1,1,4]
输出: 2
解释: 跳到最后一个位置的最小跳跃数是 2。
    从下标为 0 跳到下标为 1 的位置，跳1步，然后跳3步到达数组的最后一个位置。

示例 2:
输入: nums = [2,3,0,1,4]
输出: 2

提示:
1 <= nums.length <= 104
0 <= nums[i] <= 1000
*/
func jump(nums []int) int {
	res := 0
	start := 0
	end := 1
	for end < len(nums) {
		maxPos := 0
		for i := start; i < end; i++ {
			maxPos = max(maxPos, i+nums[i])
		}
		start = end
		end = maxPos + 1
		res++
	}

	return res
}

func jump1(nums []int) int {
	res := 0
	end := 0
	maxPos := 0
	for i := 0; i < len(nums)-1; i++ {
		//能跳到的最远距离
		maxPos = max(maxPos, i+nums[i])
		if i == end {
			end = maxPos
			res++
		}
	}

	return res
}

func max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}

/** java

class Solution {
    public int jump(int[] nums) {
        int res = 0;
        int start = 0;
        int end = 1;
        while (end < nums.length) {
            int maxPos = 0;
            for (int i = start; i < end; i++) {
                //能跳到的最远距离
                maxPos = Math.max(maxPos, i + nums[i]);
            }
            start = end;
            end = maxPos + 1;
            res++;
        }
        return res;
    }
}

class Solution {
    public int jump(int[] nums) {
        int res = 0;
        int end = 0;
        int maxPos = 0;
        for (int i = 0; i < nums.length - 1; i++) {
            //能跳到的最远距离
            maxPos = Math.max(maxPos, i + nums[i]);
            if (i == end) {
                end = maxPos;
                res++;
            }
        }

        return res;
    }
}
*/
func main() {

}
