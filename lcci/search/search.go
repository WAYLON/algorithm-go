package main

/**
搜索旋转数组。给定一个排序后的数组，包含n个整数，但这个数组已被旋转过很多次了，次数不详。请编写代码找出数组中的某个元素，假设数组元素原先是按升序排列的。若有多个相同元素，返回索引值最小的一个。

示例1:
输入: arr = [15, 16, 19, 20, 25, 1, 3, 4, 5, 7, 10, 14], target = 5
输出: 8（元素5在该数组中的索引）

示例2:
输入：arr = [15, 16, 19, 20, 25, 1, 3, 4, 5, 7, 10, 14], target = 11
输出：-1 （没有找到）

提示:
arr 长度范围在[1, 1000000]之间
*/
/**
二分法
*/
func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	if right == -1 {
		return -1
	}
	//1.注意可以等于
	for left <= right {
		if nums[left] == target {
			//  1. 找到结果, 因为找的是最小的索引, 所以当start符合时直接返回
			return left
		}
		mid := (left + right) / 2
		if nums[mid] == target {
			//  2. mid的值是目标的值, 更新右边界
			right = mid
		} else if nums[mid] > nums[left] {
			//  3.1 左边是升序的
			if nums[left] <= target && target <= nums[mid] {
				//  去掉右边的数据
				right = mid
			} else {
				//  去掉左边的数据
				left = mid + 1
			}
		} else if nums[mid] < nums[left] {
			//  3.2 右边是升序的
			if nums[mid] <= target && target <= nums[right] {
				//  去掉左边的数据
				left = mid
			} else {
				//  去掉右边的数据
				right = mid - 1
			}
		} else {
			//  左边索引前移
			left++
		}
	}
	return -1
}

/**
恢复二段性
*/
func search2(nums []int, target int) int {
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
		return ans
	}
	ans = find(nums, idx, n-1, target)
	return ans
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
		return r
	}
	return -1
}

/**
java版
class Solution {
    public int search(int[] nums, int target) {
        if (nums == null || nums.length == 0) {
            //  特殊情况判断
            return -1;
        }

        //  声明左索引
        int left = 0;
        //  声明右索引
        int right = nums.length - 1;
        while (left <= right) {
            if (nums[left] == target) {
                //  1. 找到结果, 因为找的是最小的索引, 所以当start符合时直接返回
                return left;
            }
            //  计算中间位置的索引
            int mid = (left + right) / 2;
            if (nums[mid] == target) {
                //  2. mid的值是目标的值, 更新右边界
                right = mid;
            } else if (nums[mid] > nums[left]) {
                //  3.1 左边是升序的
                if (nums[left] <= target && target <= nums[mid]) {
                    //  去掉右边的数据
                    right = mid;
                } else {
                    //  去掉左边的数据
                    left = mid + 1;
                }
            } else if (nums[mid] < nums[left]) {
                //  3.2 右边是升序的
                if (nums[mid] <= target && target <= nums[right]) {
                    //  去掉左边的数据
                    left = mid;
                } else {
                    //  去掉右边的数据
                    right = mid - 1;
                }
            } else {
                //  左边索引前移
                left++;
            }
        }
        //  没找到目标
        return -1;
    }
}

class Solution {
    public int search(int[] arr, int target) {
        int len = arr.length;
        int L = 0;
        int R = len - 1;
        //进行预处理
        while (arr[R] == arr[0]) R--;
        //开始进行二分查找转折点,找到最后一个值大于arr[0]的元素的下标
        while (L < R) {
            int mid = (L + R + 1) >> 1;
            if (arr[mid] >= arr[0]) {
                L = mid;
            } else {
                R = mid - 1;
            }
        }
        //程序走到这里R已经保存了转折点的下标,进行转折点的左边和右边进行二分查找
        int idx = binSearch(arr, 0, R, target);//若找到返回下标否则返回-1
        if (idx == -1) {
            idx = binSearch(arr, R + 1, len - 1, target);
            return idx;
        }
        return idx;
    }

    private int binSearch(int[] arr, int L, int R, int target) {
        while (L < R) {
            int mid = (L + R) >> 1;
            if (arr[mid] >= target) {
                R = mid;
            } else {
                L = mid + 1;
            }
        }
        return arr[R] == target ? R : -1;
    }
}
*/
func main() {

}
