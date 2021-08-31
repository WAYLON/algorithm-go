package main

/**
给定一个整数数组 nums，求出数组从索引i到j（i≤j）范围内元素的总和，包含i、j两点。
实现 NumArray 类：
NumArray(int[] nums) 使用数组 nums 初始化对象
int sumRange(int i, int j) 返回数组 nums 从索引i到j（i≤j）范围内元素的总和，包含i、j两点（也就是 sum(nums[i], nums[i + 1], ... , nums[j])）

示例：
输入：
["NumArray", "sumRange", "sumRange", "sumRange"]
[[[-2, 0, 3, -5, 2, -1]], [0, 2], [2, 5], [0, 5]]
输出：
[null, 1, -1, -3]
解释：
NumArray numArray = new NumArray([-2, 0, 3, -5, 2, -1]);
numArray.sumRange(0, 2); // return 1 ((-2) + 0 + 3)
numArray.sumRange(2, 5); // return -1 (3 + (-5) + 2 + (-1))
numArray.sumRange(0, 5); // return -3 ((-2) + 0 + 3 + (-5) + 2 + (-1))

提示：
0 <= nums.length <= 104
-105<= nums[i] <=105
0 <= i <= j < nums.length
最多调用 104 次 sumRange 方法

*/
type NumArray struct {
	sums []int
}

func Constructor(nums []int) NumArray {
	n := len(nums)
	array := NumArray{sums: make([]int, n+1)}
	for i := 0; i < n; i++ {
		array.sums[i+1] = array.sums[i] + nums[i]
	}
	return array
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.sums[right+1] - this.sums[left]
}

/**
java

class NumArray {
    private int[] sums;

    public NumArray(int[] nums) {
        int n = nums.length;
        sums = new int[n + 1];
        for (int i = 0; i < nums.length; i++) {
            sums[i + 1] = sums[i] + nums[i];
        }

    }

    public int sumRange(int left, int right) {
        return sums[right + 1] - sums[left];
    }
}

*/
func main() {

}
