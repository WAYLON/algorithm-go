package main

/**
统计一个数字在排序数组中出现的次数。

示例 1:
输入: nums = [5,7,7,8,8,10], target = 8
输出: 2

示例2:
输入: nums = [5,7,7,8,8,10], target = 6
输出: 0

限制：
0 <= 数组长度 <= 50000

注意：本题与主站 34 题相同（仅返回值不同）：https://leetcode-cn.com/problems/find-first-and-last-position-of-element-in-sorted-array/
*/
/**
两次2分 找左和右
*/
func search(nums []int, target int) int {
	left := 0
	right := len(nums)
	var mid int
	for left < right {
		mid = (left + right) >> 1
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	i := left
	right = len(nums)
	for left < right {
		mid = (left + right) >> 1
		if nums[mid] <= target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	j := left
	return j - i
}

/**
错误解法
*/
func search1(nums []int, target int) int {
	res := 0
	n := len(nums)
	left := 0
	right := n - 1

	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			i := mid
			j := mid + 1
			for i >= 0 && nums[i] == target {
				i--
				res++
			}
			for j < n && nums[j] == target {
				j++
				res++
			}
			break
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return res
}

/**

class Solution {
    public int search(int[] nums, int target) {
        // 搜索右边界 right
        int i = 0, j = nums.length - 1;
        while (i <= j) {
            int m = (i + j) / 2;
            if (nums[m] <= target) i = m + 1;
            else j = m - 1;
        }
        int right = i;
        // 若数组中无 target ，则提前返回
        if (j >= 0 && nums[j] != target) return 0;
        // 搜索左边界 right
        i = 0;
        j = nums.length - 1;
        while (i <= j) {
            int m = (i + j) / 2;
            if (nums[m] < target) i = m + 1;
            else j = m - 1;
        }
        int left = j;
        return right - left - 1;
    }
}

class Solution {
    public int search(int[] nums, int target) {
        int res = 0;
        int n = nums.length;
        if (n == 0) {
            return res;
        }

        int left = 0;
        int right = nums.length - 1;

        while (left <= right) {
            int mid = left + (right - left) / 2;

            if (nums[mid] == target) {
                int i = mid;
                int j = mid + 1;
                while (i >= 0 && nums[i--] == target) {
                    res++;
                }
                while (j < n && nums[j++] == target) {
                    res++;
                }
                break;
            } else if (nums[mid] > target) {
                right = mid - 1;
            } else {
                left = mid + 1;
            }

        }

        return res;
    }

    public static void main(String[] args) {
        new Solution().search(new int[]{1, 2, 3}, 3);
    }
}
*/
func main() {

}
