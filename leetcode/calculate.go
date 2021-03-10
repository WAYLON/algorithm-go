package main

/**
实现一个基本的计算器来计算一个简单的字符串表达式 s 的值。

示例 1：
输入：s = "1 + 1"
输出：2

示例 2：
输入：s = " 2-1 + 2 "
输出：3
示例 3：

输入：s = "(1+(4+5+2)-3)+(6+8)"
输出：23

提示：
1 <= s.length <= 3* 105
s 由数字、'+'、'-'、'('、')'、和 ' ' 组成
s 表示一个有效的表达式
*/
func calculate(s string) int {
	stack := []int{}
	length := len(s)
	sign, res := 1, 0
	for i := 0; i < length; i++ {
		ch := s[i]
		if '0' <= ch && ch <= '9' {
			cur := int(ch - '0')
			for i+1 < length && '0' <= s[i+1] && s[i+1] <= '9' {
				cur = cur*10 + int(s[i+1]-'0')
				i++
			}
			res = res + sign*cur
		} else if ch == '+' {
			sign = 1
		} else if ch == '-' {
			sign = -1
		} else if ch == '(' {
			stack = append(stack, res)
			res = 0
			stack = append(stack, sign)
			sign = 1
		} else if ch == ')' {
			num1 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			num2 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res = num1*res + num2
		}
	}
	return res
}

func main() {
	calculate("2147483647")
}
