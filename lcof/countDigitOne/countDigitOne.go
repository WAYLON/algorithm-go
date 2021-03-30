package main

/**
输入一个整数 n ，求1～n这n个整数的十进制表示中1出现的次数。
例如，输入12，1～12这些整数中包含1 的数字有1、10、11和12，1一共出现了5次。

示例 1：
输入：n = 12
输出：5

示例 2：
输入：n = 13
输出：6

限制：
1 <= n <2^31
*/

/**
 * 假设 n = 1204。此时十位是 0，我们把十位左边的叫做高位 high，所以这里 high = 12；把十位右边的叫做低位 low，所以这里 low = 4。
 * 十位能出现 1 的数字范围是：0010～1119，如果我们把十位遮住（也就是锁定十位，就可以不看它了），那就会变成：000～119，这里一共有：119 - 000 + 1 = 120 个数。
 * 是不是刚好等于 high * 10 也就是 high * digit？那我们就找到了当某一位等于 0 的时候，这一位一共可以出现的数字个数就是：high * digit。
 * 假设 n = 1214。此时十位是 1，同样我们有 high = 12，low = 4。十位能出现 1 的数字范围是：0010～1214，
 * 如果我们把十位遮住，那就会变成：000～124，这里一共有：124 - 000 + 1 = 125 个数。刚好等于 12 * 10 + 4 + 1 也就是 high * digit + low + 1。
 * 那我们就找到了当某一位等于 1 的时候，这一位一共可以出现的数字个数就是：high * digit + low + 1。
 * 假设 n = 1234。此时十位是 3，同样我们有 high = 12，low = 4。十位能出现 1 的数字范围是：0010～1219，
 * 如果我们把十位遮住，那就会变成：000～129，这里一共有：129 - 000 + 1 = 130 个数。刚好等于 (12 + 1) * 10 也就是 (high + 1) * digit。
 * 那我们就找到了当某一位等于 2/3/4/.../9 的时候，这一位一共可以出现的数字个数就是：(high + 1) * digit。
 *
 */
func countDigitOne(n int) int {
	res := 0
	//位数
	digit := 1
	//高位
	high := n / 10
	//当前位
	cur := n % 10
	//表示低位
	low := 0

	for high != 0 || cur != 0 {
		if cur == 0 {
			res += high * digit
		} else if cur == 1 {
			res += high*digit + low + 1
		} else {
			res += (high + 1) * digit
		}
		low += cur * digit
		cur = high % 10
		high /= 10
		digit *= 10
	}
	return res
}
func main() {

}