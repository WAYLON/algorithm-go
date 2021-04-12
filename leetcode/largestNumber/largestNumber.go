package main

import (
	"bytes"
	"sort"
	"strconv"
)

/**
给定一组非负整数 nums，重新排列每个数的顺序（每个数不可拆分）使之组成一个最大的整数。
注意：输出结果可能非常大，所以你需要返回一个字符串而不是整数。

示例 1：
输入：nums = [10,2]
输出："210"

示例2：
输入：nums = [3,30,34,5,9]
输出："9534330"

示例 3：
输入：nums = [1]
输出："1"

示例 4：
输入：nums = [10]
输出："10"

提示：
1 <= nums.length <= 100
0 <= nums[i] <= 109
*/
func largestNumber(nums []int) string {
	n := len(nums)
	// 处理边界情况，如果数组为空或长度为0
	if nums == nil || n == 0 {
		return "" // 返回空的字符串
	}
	// 否则定义一个长度为n的字符串数组
	strs := make([]string, n)
	// 遍历整数数组,把整数转成字符串,并存储到字符串数组
	for i := range nums {
		strs[i] = strconv.Itoa(nums[i])
	}
	// 接下来对字符串进行排序,注意这里要自定义对比规则
	sort.Slice(strs, func(i, j int) bool {
		// 把s1和s2的长度赋值给n1和n2方便使用
		n1, n2 := len(strs[i]), len(strs[j])
		// 然后下标从0遍历到n1+n2-1，模拟s12和s21的对比
		for c := 0; c < n1+n2; c++ {
			// 模拟从s12取第i个字符的操作
			var c1 byte
			if c < n1 {
				c1 = strs[i][c]
			} else {
				c1 = strs[j][c-n1]
			}

			// 模拟从s21取第i个字符的操作
			var c2 byte
			if c < n2 {
				c2 = strs[j][c]
			} else {
				c2 = strs[i][c-n2]
			}

			// 对比取出的两个字符
			if c1 == c2 {
				continue
			}
			return c1 > c2
		}
		return true
	})
	// 首先检查排序后第0个元素是不是字符串0
	if strs[0] == "0" {
		return "0" // 说明后面都是0，直接返回0即可
	}
	// 把数组中的字符串依次拼接起来
	var buffer bytes.Buffer
	for i := range strs {
		buffer.WriteString(strs[i])
	}
	return buffer.String()
}

/**
Java版
class Solution {
    public String largestNumber(int[] nums) {
        int n = nums.length;
        String[] strings = new String[n];
        for (int i = 0; i < n; i++) {
            strings[i] = String.valueOf(nums[i]);
        }
        Arrays.sort(strings, (a, b) -> (b + a).compareTo(a + b));

        if (("0".equals(strings[0]))) return "0";

        StringBuilder sb = new StringBuilder();
        for (String num : strings) {
            sb.append(num);
        }

        return sb.toString();
    }
}
*/
func main() {

}
