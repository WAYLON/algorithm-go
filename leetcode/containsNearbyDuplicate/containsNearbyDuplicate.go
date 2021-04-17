package main

/**
给定一个整数数组和一个整数k，判断数组中是否存在两个不同的索引i和j，使得nums [i] = nums [j]，并且 i 和 j的差的 绝对值 至多为 k。

示例1:
输入: nums = [1,2,3,1], k = 3
输出: true

示例 2:
输入: nums = [1,0,1,1], k = 1
输出: true

示例 3:
输入: nums = [1,2,3,1,2,3], k = 2
输出: false
*/
// 滑动窗口+查找表 Time: O(n), Space: O(k)
func containsNearbyDuplicate(nums []int, k int) bool {
	record := make(map[int]int) // 记录窗口的查找表
	for i := 0; i < len(nums); i++ {
		if _, ok := record[nums[i]]; ok {
			return true // 在窗口中找到这个元素
		}
		// 否则说明新的数与窗口中任意数都不重复
		record[nums[i]] = i // 将新的数添加到窗口中
		// 判断这个窗口有多大,保持record中最多有k个元素
		if len(record) == k+1 {
			// 删除这个窗口最左侧的数据
			delete(record, nums[i-k])
		}
	}
	return false // 循环结束后没有返回true即没有找到满足条件
}

//优化
func containsNearbyDuplicate1(nums []int, k int) bool {
	hash := make(map[int]int)
	for i, x := range nums {
		if j, ok := hash[x]; ok && i-j <= k {
			return true
		}
		hash[x] = i
	}
	return false
}

/**
class Solution {
    public boolean containsNearbyDuplicate(int[] nums, int k) {
        Set<Integer> set = new HashSet<>();
        for (int i = 0; i < nums.length; i++) {
            if (set.contains(nums[i])) {
                return true;
            }
            set.add(nums[i]);
            if (set.size() > k) {
                set.remove(nums[i - k]);
            }
        }
        return false;
    }
}
*/
func main() {

}
