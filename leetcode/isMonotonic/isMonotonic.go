package main

/**
如果数组是单调递增或单调递减的，那么它是单调的。
如果对于所有 i <= j，A[i] <= A[j]，那么数组 A 是单调递增的。 如果对于所有 i <= j，A[i]> = A[j]，那么数组 A 是单调递减的。
当给定的数组 A是单调数组时返回 true，否则返回 false。

示例 1：
输入：[1,2,2,3]
输出：true

示例 2：
输入：[6,5,4,4]
输出：true
提示：
1 <= A.length <= 50000
-100000 <= A[i] <= 100000

*/
func isMonotonic(A []int) bool {
	inc := true
	dec := true

	for i := 1; i < len(A); i++ {
		if A[i] < A[i-1] {
			inc = false
		} else if A[i] > A[i-1] {
			dec = false
		}
		if !inc && !dec {
			return false
		}

	}
	return true
}

func main() {

}
