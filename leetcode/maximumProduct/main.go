package main

import (
	"math"
	"sort"
)

/**
给定一个整型数组，在数组中找出由三个数组成的最大乘积，并输出这个乘积。
*/
func maximumProduct(nums []int) int {
	sort.Ints(nums)
	n := len(nums)
	return max(nums[0]*nums[1]*nums[n-1], nums[n-3]*nums[n-2]*nums[n-1])
}

func maximumProduct2(nums []int) int {
	// 最小的和第二小的
	min1, min2 := math.MaxInt64, math.MaxInt64
	// 最大的、第二大的和第三大的
	max1, max2, max3 := math.MinInt64, math.MinInt64, math.MinInt64

	for _, x := range nums {
		if x < min1 {
			min2 = min1
			min1 = x
		} else if x < min2 {
			min2 = x
		}

		if x > max1 {
			max3 = max2
			max2 = max1
			max1 = x
		} else if x > max2 {
			max3 = max2
			max2 = x
		} else if x > max3 {
			max3 = x
		}
	}

	return max(min1*min2*max1, max1*max2*max3)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	arr := []int{4, 93, 84, 85, 80, 37, 81, 93, 27, 12}
	maximumProduct(arr)
	maximumProduct2(arr)
}
