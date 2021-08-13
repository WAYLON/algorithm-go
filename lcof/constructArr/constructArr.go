package main

/**
给定一个数组 A[0,1,…,n-1]，请构建一个数组 B[0,1,…,n-1]，其中B[i] 的值是数组 A 中除了下标 i 以外的元素的积, 即B[i]=A[0]×A[1]×…×A[i-1]×A[i+1]×…×A[n-1]。不能使用除法。

示例:
输入: [1,2,3,4,5]                                 1 2 3 4 5
输出: [120,60,40,30,24]							 1 1 3 4 5
												 1 2 1 4 5
												 1 2 3 1 5
											 	 1 2 3 4 1
提示：
所有元素乘积之和不会溢出 32 位整数
a.length <= 100000
*/
func constructArr(a []int) []int {
	n := len(a)
	if n == 0 {
		return a
	}
	left := make([]int, n)
	right := make([]int, n)
	left[0] = 1
	for i := 1; i < n; i++ {
		left[i] = left[i-1] * a[i-1]
	}
	right[n-1] = 1
	for i := n - 2; i >= 0; i-- {
		right[i] = right[i+1] * a[i+1]
	}
	for i := 0; i < n; i++ {
		left[i] = left[i] * right[i]
	}

	return left
}

/**
class Solution {
    public int[] constructArr(int[] a) {
        if (a == null || a.length == 0) return a;
        int len = a.length;
        int[] left = new int[len];
        int[] right = new int[len];
        left[0] = right[len - 1] = 1;

        for (int i = 1; i < len; i++) {
            left[i] = left[i - 1] * a[i - 1];
        }
        for (int i = len - 2; i >= 0; i--) {
            right[i] = right[i + 1] * a[i + 1];
        }

        int[] ans = new int[len];
        for (int i = 0; i < len; i++) {
            ans[i] = left[i] * right[i];
        }
        return ans;
    }
}


*/
func main() {
	constructArr([]int{1, 2, 3, 4, 5})
}
