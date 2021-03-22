package main

/**
把一个数组最开始的若干个元素搬到数组的末尾，我们称之为数组的旋转。输入一个递增排序的数组的一个旋转，输出旋转数组的最小元素。例如，数组[3,4,5,1,2] 为 [1,2,3,4,5] 的一个旋转，该数组的最小值为1。

示例 1：
输入：[3,4,5,1,2]
输出：1

示例 2：
输入：[2,2,2,0,1]
输出：0
*/

/**
二分法
*/
func minArray(numbers []int) int {
	i := 0
	j := len(numbers) - 1
	for i < j {
		mid := (i + j) / 2
		if numbers[mid] < numbers[j] {
			j = mid
		} else if numbers[mid] > numbers[j] {
			i = mid + 1
		} else {
			x := i
			for k := i + 1; k < j; k++ {
				if numbers[k] < numbers[x] {
					x = k
				}
			}
			return numbers[x]
		}

	}
	return numbers[i]
}

func main() {

}
