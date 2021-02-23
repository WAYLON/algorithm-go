package main

/**
今天，书店老板有一家店打算试营业customers.length分钟。每分钟都有一些顾客（customers[i]）会进入书店，所有这些顾客都会在那一分钟结束后离开。
在某些时候，书店老板会生气。 如果书店老板在第 i 分钟生气，那么 grumpy[i] = 1，否则 grumpy[i] = 0。 当书店老板生气时，那一分钟的顾客就会不满意，不生气则他们是满意的。
书店老板知道一个秘密技巧，能抑制自己的情绪，可以让自己连续X 分钟不生气，但却只能使用一次。
请你返回这一天营业下来，最多有多少客户能够感到满意的数量。

示例：
输入：customers = [1,0,1,2,1,1,7,5], grumpy = [0,1,0,1,0,1,0,1], X = 3
输出：16
解释：
书店老板在最后 3 分钟保持冷静。
感到满意的最大客户数量 = 1 + 1 + 1 + 1 + 7 + 5 = 16.

提示：
1 <= X <=customers.length ==grumpy.length <= 20000
0 <=customers[i] <= 1000
0 <=grumpy[i] <= 1

*/
func maxSatisfied(customers []int, grumpy []int, X int) int {
	total := 0
	n := len(customers)
	//1.求出老板所有不生气时候的顾客
	for i := 0; i < n; i++ {
		if grumpy[i] == 0 {
			total += customers[i]
		}
	}

	//2.求出老板第一个窗口的生气顾客数
	increase := 0
	for i := 0; i < X; i++ {
		increase += customers[i] * grumpy[i]
	}
	var maxIncrease = increase

	//3.滑动窗口 求出大小为X的窗口中老板生气顾客数的最大值
	for i := X; i < n; i++ {
		increase = increase - customers[i-X]*grumpy[i-X] + customers[i]*grumpy[i]
		if increase > maxIncrease {
			maxIncrease = increase
		}
	}
	return total + maxIncrease

}

func main() {

}
