package main

import "fmt"

/**
实现一个算法，确定一个字符串 s 的所有字符是否全都不同。

示例 1：
输入: s = "leetcode"
输出: false

示例 2：
输入: s = "abc"
输出: true

限制：
0 <= len(s) <= 100
如果你不使用额外的数据结构，会很加分。

把不同的字符与字符'a'之间的差值，转换成1左移的位数，
比如字符'c','c'-'a'=2,把1左移两位 ==> 100，d就是1000
把差值与一个数num进行与运算(&),num初始值为0，

如果之前没有相同的字符出现，&结果为0，把出现的字符用或运算(|)存储到num中，以便下一次比较
如果之前有相同的字符出现，&结果为1

*/
func isUnique(astr string) bool {
	num := 0
	for _, v := range astr {
		moveBit := v - 'a'
		if num&(1<<moveBit) != 0 {
			return false
		} else {
			num = num | (1 << moveBit)
		}
	}
	return true
}

/**
使用map，记录每个字符出现的次数，如果次数大于1,则表示有重复的字符。
*/
func isUnique1(astr string) bool {
	md := make(map[rune]int)
	for _, v := range astr {
		md[v]++
		if md[v] > 1 {
			return false
		}
	}
	return true
}

/**
利用map判断是否存在key值
*/
func isUnique2(astr string) bool {
	var str = make(map[string]bool)
	arr := []rune(astr)
	for i := 0; i < len(arr); i++ {
		fmt.Printf(string(arr[i]))
		if _, ok := str[string(arr[i])]; ok {
			return false
		}
		str[string(arr[i])] = true
	}
	return true
}

func main() {
	isUnique("leetcode")
}
