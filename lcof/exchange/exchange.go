package main

/**
输入一个整数数组，实现一个函数来调整该数组中数字的顺序，使得所有奇数位于数组的前半部分，所有偶数位于数组的后半部分。

示例：
输入：nums =[1,2,3,4]
输出：[1,3,2,4]
注：[3,1,2,4] 也是正确的答案之一。

提示：
0 <= nums.length <= 50000
1 <= nums[i] <= 10000

*/

/**
方法1 双指针
*/
func exchange(nums []int) []int {
	left := 0
	right := len(nums) - 1
	for left < right {
		//找到偶数
		for left < right && nums[left]%2 == 1 {
			left++
		}
		//找到奇数
		for left < right && nums[right]%2 == 0 {
			right--
		}
		if left < right {
			nums[left], nums[right] = nums[right], nums[left]
		}
	}
	return nums
}

/**
方法2 快慢指针
*/
func exchange2(nums []int) []int {
	//负责寻找奇数
	fast := 0
	//负责记录奇数下一次插入的位置
	slow := 0
	for fast < len(nums) {
		if nums[fast]&1 == 1 {
			nums[fast], nums[slow] = nums[slow], nums[fast]
			slow++
		}
		fast++
	}
	return nums
}

func main() {

}
