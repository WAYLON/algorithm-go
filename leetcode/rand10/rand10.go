package main

import "math/rand"

/**
已有方法rand7可生成 1 到 7 范围内的均匀随机整数，试写一个方法rand10生成 1 到 10 范围内的均匀随机整数。
不要使用系统的Math.random()方法。

示例 1:
输入: 1
输出: [7]

示例 2:
输入: 2
输出: [8,4]

示例 3:
输入: 3
输出: [8,1,10]

提示:
rand7已定义。
传入参数:n表示rand10的调用次数。

进阶:
rand7()调用次数的期望值是多少?
你能否尽量少调用 rand7() ?
*/

/**
rand7()  						1  2  3  4  5  6  7
rand7() * 7						7  14 21 28 35 42 49
rand7() * 7 - (rand7() -1 )		7  14 21 28 35 42 49   - [0,6]

*/
func rand10() int {
	for {
		ans := rand7()*7 - (rand7() - 1)
		if ans <= 40 {
			return ans%10 + 1
		}
	}
}

func rand7() int {
	return rand.Intn(7)
}

/** java
 class Solution extends SolBase {
    public int rand10() {
        while (true) {
            int ans = rand7() * 7 - (rand7() - 1);
            if (ans <= 40) return ans % 10 + 1;
        }
    }
}
*/

func main() {

}
