package main

import (
	"bytes"
	"strconv"
)

/**
输入一个非负整数数组，把数组里所有数字拼接起来排成一个数，打印能拼接出的所有数字中最小的一个。

示例 1:
输入: [10,2]
输出: "102"

示例n2:
输入: [3,30,34,5,9]
输出: "3033459"

提示:
0 < nums.length <= 100

说明:
输出结果可能非常大，所以你需要返回一个字符串而不是整数
拼接起来的数字可能会有前导 0，最后结果不需要去掉前导 0
*/
func minNumber(nums []int) string {
	n := len(nums)
	strs := make([]string, n)
	for i := 0; i < n; i++ {
		strs[i] = strconv.Itoa(nums[i])
	}
	quickSort(strs, 0, n-1)
	var buffer bytes.Buffer
	for i := 0; i < n; i++ {
		buffer.WriteString(strs[i])
	}
	return buffer.String()
}

func quickSort(strs []string, low int, high int) {
	if low >= high {
		return
	}
	i := low
	j := high
	sentinel := strs[low]
	for i < j {
		for i < j && (sentinel+strs[j]) <= (strs[j]+sentinel) {
			j--
		}
		for i < j && (sentinel+strs[i]) >= (strs[i]+sentinel) {
			i++
		}
		if i < j {
			strs[i], strs[j] = strs[j], strs[i]
		}
	}
	strs[low], strs[i] = strs[i], strs[low]
	quickSort(strs, low, i-1)
	quickSort(strs, j+1, high)
}

/**
Java版
class Solution {
    public String minNumber(int[] nums) {
        String[] strs = new String[nums.length];
        for (int i = 0; i < nums.length; i++)
            strs[i] = String.valueOf(nums[i]);
        quickSort(strs, 0, strs.length - 1);
        StringBuilder res = new StringBuilder();
        for (String s : strs)
            res.append(s);
        return res.toString();
    }

    void quickSort(String[] strs, int l, int r) {
        if (l >= r) return;
        int i = l, j = r;
        String tmp = strs[i];
        while (i < j) {
            while ((strs[j] + strs[l]).compareTo(strs[l] + strs[j]) >= 0 && i < j) j--;
            while ((strs[i] + strs[l]).compareTo(strs[l] + strs[i]) <= 0 && i < j) i++;
            tmp = strs[i];
            strs[i] = strs[j];
            strs[j] = tmp;
        }
        strs[i] = strs[l];
        strs[l] = tmp;
        quickSort(strs, l, i - 1);
        quickSort(strs, i + 1, r);
    }
}

*/
func main() {
	minNumber([]int{3, 30, 34, 5, 9})
}
