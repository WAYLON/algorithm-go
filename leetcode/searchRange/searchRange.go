package main

/**
给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。
如果数组中不存在目标值 target，返回[-1, -1]。

进阶：
你可以设计并实现时间复杂度为O(log n)的算法解决此问题吗？

示例 1：
输入：nums = [5,7,7,8,8,10], target = 8
输出：[3,4]

示例2：
输入：nums = [5,7,7,8,8,10], target = 6
输出：[-1,-1]

示例 3：
输入：nums = [], target = 0
输出：[-1,-1]

提示：
0 <= nums.length <= 105
-109<= nums[i]<= 109
nums是一个非递减数组
-109<= target<= 109
*/

/**
方法1
*/
func searchRange(nums []int, target int) []int {
	n := len(nums)
	res := []int{-1, -1}
	if n == 0 {
		return res
	}
	left := 0
	right := n - 1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			min := mid
			max := mid
			for 0 <= min-1 && nums[min-1] == nums[mid] {
				min--
			}
			for max+1 < n && nums[max+1] == nums[mid] {
				max++
			}
			res[0] = min
			res[1] = max
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
方法2 2次二分
*/
func searchRange2(nums []int, target int) []int {
	n := len(nums)
	res := []int{-1, -1}
	if n == 0 {
		return res
	}
	left := 0
	right := n - 1
	for left < right {
		mid := (left + right) >> 1
		if nums[mid] >= target {
			right = mid
		} else {
			left = mid + 1
		}
	}
	if nums[left] != target {
		return res
	} else {
		res[0] = left
		left = 0
		right = n - 1
		for left < right {
			mid := (left + right + 1) >> 1
			if nums[mid] <= target {
				left = mid
			} else {
				right = mid - 1
			}
		}
		res[1] = left
		return res
	}
}

/**
java版
class Solution {
    public int[] searchRange(int[] nums, int target) {
        int[] ans = new int[]{-1, -1};
        int n = nums.length;
        if (n == 0) return ans;

        int l = 0, r = n - 1;

        while (l < r) {
            //nums = [5,7,7,8,8,10], target = 8
            int mid = l + r >> 1;
            if (nums[mid] >= target) {
                r = mid;
            } else {
                l = mid + 1;
            }
        }

        if (nums[l] != target) {
            return ans;
        } else {
            ans[0] = l;
            l = 0; r = n - 1;
            while (l < r) {
                int mid = l + r + 1 >> 1;
                if (nums[mid] <= target) {
                    l = mid;
                } else {
                    r = mid - 1;
                }
            }
            ans[1] = l;
            return ans;
        }
    }
}

class Solution {
    public int[] searchRange(int[] nums, int target) {
        int[] res = new int[]{-1, -1};
        int n = nums.length;
        if (n == 0) return res;

        int left = 0;
        int right = n - 1;

        while (left <= right) {
            int mid = (left + right) / 2;
            if (nums[mid] == target) {
                int min = mid;
                int max = mid;
                while (0 <= min - 1 && nums[min - 1] == nums[mid]) {
                    min--;
                }
                while (max + 1 < n && nums[max + 1] == nums[mid]) {
                    max++;
                }
                res[0] = min;
                res[1] = max;
                break;

            } else if (nums[mid] < target) {
                left = mid + 1;
            } else {
                right = mid - 1;
            }
        }

        return res;
    }
}
*/
func main() {

}
