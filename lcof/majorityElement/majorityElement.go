package main

import "sort"

/**
数组中有一个数字出现的次数超过数组长度的一半，请找出这个数字。
你可以假设数组是非空的，并且给定的数组总是存在多数元素。

示例1:
输入: [1, 2, 3, 2, 2, 2, 5, 4, 2]
输出: 2

限制：
1 <= 数组长度 <= 50000
*/

/**
方法1 摩尔投票法
*/
func majorityElement(nums []int) int {
	votes := 0
	sum := 0
	for _, v := range nums {
		if sum == 0 {
			votes = v
		}
		if v == votes {
			sum++
		} else {
			sum--
		}
	}
	// 验证 votes 是否为众数
	x := 0
	for _, v := range nums {
		if v == votes {
			x++
		}
	}
	if x > len(nums)/2 {
		return votes
	}

	return 0
}

/**
方法2 字典
*/
func majorityElement2(nums []int) int {
	m := make(map[int]int)
	for _, v := range nums {
		m[v] = m[v] + 1
	}
	for k, v := range m {
		if v > len(nums)/2 {
			return k
		}
	}

	return 0
}

/**
方法3 排序
*/
func majorityElement3(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)/2]
}

func main() {

}
