package main

import (
	"sort"
)

/**
森林中，每个兔子都有颜色。其中一些兔子（可能是全部）告诉你还有多少其他的兔子和自己有相同的颜色。我们将这些回答放在answers数组里。
返回森林中兔子的最少数量。

示例:
输入: answers = [1, 1, 2]
输出: 5
解释:
两只回答了 "1" 的兔子可能有相同的颜色，设为红色。
之后回答了 "2" 的兔子不会是红色，否则他们的回答会相互矛盾。
设回答了 "2" 的兔子为蓝色。
此外，森林中还应有另外 2 只蓝色兔子的回答没有包含在数组中。
因此森林中兔子的最少数量是 5: 3 只回答的和 2 只没有回答的。

输入: answers = [10, 10, 10]
输出: 11

输入: answers = []
输出: 0
说明:

answers的长度最大为1000。
answers[i]是在[0, 999]范围内的整数。
*/
/**
方法1：模拟解法
*/
func numRabbits(answers []int) int {
	sort.Ints(answers)
	n := len(answers)
	ans := 0
	for i := 0; i < n; i++ {
		cnt := answers[i]
		ans += cnt + 1
		// 跳过「数值 cnt」后面的 cnt 个「数值 cnt」
		k := cnt
		for k > 0 && i+1 < n && answers[i] == answers[i+1] {
			k--
			i++
		}
	}
	return ans
}

/**
方法2：统计分配
*/
func numRabbits2(answers []int) int {
	N := 1009
	counts := make([]int, N)
	for _, v := range answers {
		counts[v]++
	}
	// counts[x] = cnt 代表在 cs 中数值 x 的数量为 cnt
	ans := counts[0]
	for i := 1; i < N; i++ {
		per := i + 1
		cnt := counts[i]
		k := cnt / per
		if k*per < cnt {
			k++
		}
		ans += k * per
	}
	return ans
}

func main() {

}
