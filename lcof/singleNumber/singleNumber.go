package main

/**
在一个数组 nums 中除一个数字只出现一次之外，其他数字都出现了三次。请找出那个只出现一次的数字。

示例 1：
输入：nums = [3,4,3,3]
输出：4

示例 2：
输入：nums = [9,1,7,9,7,9,7]
输出：1

限制：
1 <= nums.length <= 10000
1 <= nums[i] < 2^31
*/

func singleNumber(nums []int) int {
	k := make([]int, 32)
	for _, v := range nums {
		for i := 0; i < 32; i++ {
			k[i] += v & 1
			v >>= 1
		}
	}
	res := 0
	for i := 0; i < 32; i++ {
		res += (1 << i) * (k[i] % 3)
		//res |= (k[i] % 3) << i
	}
	return res
}

/**
class Solution {
    public int singleNumber(int[] nums) {
        int[] k = new int[32];
        for (int num : nums) {
            for (int i = 0; i < 32; i++) {
                k[i] += num & 1;
                num >>= 1;
            }
        }
        int res = 0;
        for (int i = 0; i < 32; i++) {
             res |= (k[i] % 3) << i;
        }
        return res;
    }
}
*/
func main() {

}
