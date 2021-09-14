package main

/**
给定一个只包含三种字符的字符串：（，）和 *，写一个函数来检验这个字符串是否为有效字符串。有效字符串具有如下规则：
任何左括号 (必须有相应的右括号 )。
任何右括号 )必须有相应的左括号 (。
左括号 ( 必须在对应的右括号之前 )。
*可以被视为单个右括号 )，或单个左括号 (，或一个空字符串。
一个空字符串也被视为有效字符串。

示例 1:
输入: "()"
输出: True

示例 2:
输入: "(*)"
输出: True

示例 3:
输入: "(*))"
输出: True
注意:

字符串大小将在 [1，100] 范围内。
*/

/**
模拟栈 操作队尾元素
*/
func checkValidString(s string) bool {
	left := make([]int, 0)
	star := make([]int, 0)
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			left = append(left, i)
		} else if s[i] == '*' {
			star = append(star, i)
		} else {
			if len(left) == 0 && len(star) == 0 {
				return false
			}
			if len(left) != 0 {
				left = left[:len(left)-1]
			} else {
				star = star[:len(star)-1]
			}
		}
	}

	for len(left) != 0 && len(star) != 0 {
		if left[len(left)-1] > star[len(star)-1] {
			return false
		}
		left = left[:len(left)-1]
		star = star[:len(star)-1]
	}

	return len(left) == 0
}

/** java

class Solution {
    public boolean checkValidString(String s) {
        Deque<Integer> stackLeft = new ArrayDeque<>();
        Deque<Integer> stackstar = new ArrayDeque<>();
        for (int i = 0; i < s.length(); i++) {
            if (s.charAt(i) == '(') {
                stackLeft.addFirst(i);
            } else if (s.charAt(i) == '*') {
                stackstar.addFirst(i);
            } else {
                if (stackLeft.isEmpty() && stackstar.isEmpty()) return false;
                if (!stackLeft.isEmpty()) stackLeft.pop();
                else stackstar.pop();
            }
        }
        while(!stackLeft.isEmpty() && !stackstar.isEmpty()){
            if (stackLeft.peekFirst() > stackstar.peekFirst()) return false;
            stackLeft.pop();
            stackstar.pop();
        }

        return stackLeft.isEmpty();
    }
}

*/
func main() {
	checkValidString("((((()(()()()*()(((((*)()*(**(())))))(())()())(((())())())))))))(((((())*)))()))(()((*()*(*)))(*)()")
}
