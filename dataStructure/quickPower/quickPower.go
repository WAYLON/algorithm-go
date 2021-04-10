package main

/**
快速幂模板
*/
func quickPower(x float64, n int) float64 {
	res := 1.0
	if n == 0 {
		return res
	}
	//因为测试用例中有-2147483648  超过了int32的最小值 即-2147483647 - 2147483648  符号位占一位
	b := int64(n)

	if b < 0 {
		x = 1 / x
		b = -b
	}

	for b > 0 {
		if b&1 == 1 { //二进制最后一位为1时即奇数，将待乘数乘以x
			res *= x
		}
		x *= x //二进制最后一位为1时即偶数，将底数平方
		b >>= 1
	}
	return res
}

func main() {
	quickPower(2, 5)
}
