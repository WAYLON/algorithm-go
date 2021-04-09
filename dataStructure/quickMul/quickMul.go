package main

/**
快速乘模板
*/
func quickMul(a int, b int) int {
	res := 0
	if b < 0 {
		a = -a
		b = -b
	}

	for b > 0 {
		if b&1 == 1 {
			res += a
		}
		a += a
		b >>= 1
	}
	return res
}

func main() {
	quickMul(3, 5)
}
