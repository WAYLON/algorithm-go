package main

/**
整数数组 nums 按升序排列，数组中的值 互不相同 。
在传递给函数之前，nums 在预先未知的某个下标 k（0 <= k < nums.length）上进行了 旋转，使数组变为 [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]]（下标 从 0 开始 计数）。例如， [0,1,2,4,5,6,7] 在下标 3 处经旋转后可能变为[4,5,6,7,0,1,2] 。
给你 旋转后 的数组 nums 和一个整数 target ，如果 nums 中存在这个目标值 target ，则返回它的下标，否则返回-1。

示例 1：
输入：nums = [4,5,6,7,0,1,2], target = 0
输出：4

示例2：
输入：nums = [4,5,6,7,0,1,2], target = 3
输出：-1

示例 3：
输入：nums = [1], target = 0
输出：-1

提示：
1 <= nums.length <= 5000
-10^4 <= nums[i] <= 10^4
nums 中的每个值都 独一无二
题目数据保证 nums 在预先未知的某个下标上进行了旋转
-10^4 <= target <= 10^4

进阶：你可以设计一个时间复杂度为 O(log n) 的解决方案吗？
*/
/**
二分法
[4,5,6,7,0,1,2]
0
*/
func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		}
		//中间元素大于等于第一个元素 证明中间元素处于第一个递增区间 否则处于第二个递增区间
		if nums[mid] >= nums[left] {
			//如果当前target值处于这个区间的话 令right等于mid-1 否则 left等于mid+1
			if nums[mid] > target && target >= nums[left] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			//如果当前target值处于这个区间的话 令left等于mid+1 否则right等于mid-1
			if nums[mid] < target && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}

/**
java版
class Solution {
    public int search(int[] nums, int target) {
        int i = 0;
        int j = nums.length - 1;

        while (i <= j) {
            int mid = (i + j) / 2;
            if (nums[mid] == target) {
                return mid;
            } else if (nums[mid] <= nums[j]) { //说明mid在后一半升序数组
                if (nums[mid] < target && target <= nums[j]) {
                    i = mid + 1;
                } else {
                    j = mid - 1;
                }
            } else if (nums[mid] > nums[j]) { //说明mid在前一半升序数组
                if (nums[mid] > target && target >= nums[i]) {
                    j = mid - 1;
                } else {
                    i = mid + 1;
                }
            }
        }
        return -1;
    }
}
*/

func main() {
}
