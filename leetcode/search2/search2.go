package main

/**
已知存在一个按非降序排列的整数数组 nums ，数组中的值不必互不相同。
在传递给函数之前，nums 在预先未知的某个下标 k（0 <= k < nums.length）上进行了 旋转 ，
使数组变为 [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]]（下标 从 0 开始 计数）。
例如， [0,1,2,4,4,4,5,6,6,7] 在下标 5 处经旋转后可能变为 [4,5,6,6,7,0,1,2,4,4] 。
给你 旋转后 的数组 nums 和一个整数 target ，请你编写一个函数来判断给定的目标值是否存在于数组中。如果 nums 中存在这个目标值 target ，则返回 true ，否则返回 false 。

示例1：
输入：nums = [2,5,6,0,0,1,2], target = 0
输出：true

示例2：
输入：nums = [2,5,6,0,0,1,2], target = 3
输出：false

提示：
1 <= nums.length <= 5000
-104 <= nums[i] <= 104
题目数据保证 nums 在预先未知的某个下标上进行了旋转
-104 <= target <= 104

进阶：
这是 搜索旋转排序数组的延伸题目，本题中的nums 可能包含重复元素。
这会影响到程序的时间复杂度吗？会有怎样的影响，为什么？
*/
/**
二分法
*/
func search(nums []int, target int) bool {
	left := 0
	right := len(nums) - 1
	//1.注意可以等于
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return true
		}

		//特判 防止 [1,0,1,1,1] 0
		if nums[mid] == nums[left] {
			left++
			continue
		}

		//说明mid在前一个递增序列
		if nums[mid] >= nums[left] {
			//当前left<=target<mid
			if nums[left] <= target && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else { //说明mid后一个递增序列
			//当前mid<target<=right
			if nums[mid] < target && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return false
}

/**
恢复二段性
*/
func search1(nums []int, target int) bool {
	n := len(nums)
	left := 0
	right := n - 1
	// 恢复二段性
	for left < right && nums[0] == nums[right] {
		right--
	}

	// 第一次二分，找旋转点
	for left < right {
		mid := (left + right + 1) >> 1
		if nums[mid] >= nums[0] {
			left = mid
		} else {
			right = mid - 1
		}
	}

	idx := n
	if nums[right] >= nums[0] && right+1 < n {
		idx = right + 1
	}
	// 第二次二分，找目标值
	ans := find(nums, 0, idx-1, target)
	if ans != -1 {
		return true
	}
	ans = find(nums, idx, n-1, target)
	return ans != -1
}

func find(nums []int, l int, r int, t int) int {
	for l < r {
		mid := (l + r) >> 1
		if nums[mid] >= t {
			r = mid
		} else {
			l = mid + 1
		}
	}
	if nums[r] == t {
		return t
	}
	return -1
}

/**
java版
class Solution {
    public boolean search(int[] nums, int target) {
        int i = 0;
        int j = nums.length - 1;

        while (i <= j) {
            int mid = (i + j) / 2;
            if (nums[mid] == target) {
                return true;
            }
            if (nums[j] == nums[mid]) {
                j--;
                continue;
            }

            if (nums[mid] <= nums[j]) { //说明mid在后一半升序数组
                if (nums[mid] < target && target <= nums[j]) {
                    i = mid + 1;
                } else {
                    j = mid - 1;
                }
            } else { //说明mid在前一半升序数组
                if (nums[mid] > target && target >= nums[i]) {
                    j = mid - 1;
                } else {
                    i = mid + 1;
                }
            }
        }
        return false;
    }
}

class Solution {
    public boolean search(int[] nums, int t) {
        int n = nums.length;
        int l = 0, r = n - 1;
        // 恢复二段性
        while (l < r && nums[0] == nums[r]) r--;

        // 第一次二分，找旋转点
        while (l < r) {
            int mid = l + r + 1 >> 1;
            if (nums[mid] >= nums[0]) {
                l = mid;
            } else {
                r = mid - 1;
            }
        }

        int idx = n;
        if (nums[r] >= nums[0] && r + 1 < n) idx = r + 1;

        // 第二次二分，找目标值
        int ans = find(nums, 0, idx - 1, t);
        if (ans != -1) return true;
        ans = find(nums, idx, n - 1, t);
        return ans != -1;
    }

    private int find(int[] nums, int l, int r, int t) {
        while (l < r) {
            int mid = l + r >> 1;
            if (nums[mid] >= t) {
                r = mid;
            } else {
                l = mid + 1;
            }
        }
        return nums[r] == t ? r : -1;
    }
}
*/
func main() {
}
