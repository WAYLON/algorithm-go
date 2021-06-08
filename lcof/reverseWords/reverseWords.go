package main

import "strings"

/**
输入一个英文句子，翻转句子中单词的顺序，但单词内字符的顺序不变。为简单起见，标点符号和普通字母一样处理。例如输入字符串"I am a student. "，则输出"student. a am I"。

示例 1：
输入: "the sky is blue"
输出:"blue is sky the"

示例 2：
输入: " hello world! "
输出:"world! hello"
解释: 输入字符串可以在前面或者后面包含多余的空格，但是反转后的字符不能包括。

示例 3：
输入: "a good  example"
输出:"example good a"
解释: 如果两个单词间有多余的空格，将反转后单词间的空格减少到只含一个。

说明：
无空格字符构成一个单词。
输入字符串可以在前面或者后面包含多余的空格，但是反转后的字符不能包括。
如果两个单词间有多余的空格，将反转后单词间的空格减少到只含一个。
注意：本题与主站 151 题相同：https://leetcode-cn.com/problems/reverse-words-in-a-string/

注意：此题对比原题有改动
*/

func reverseWords(s string) string {
	//删除收尾空格并分割
	strs := strings.Split(strings.Trim(s, " "), " ")
	res := ""
	//倒序遍历
	for i := len(strs) - 1; i >= 0; i-- {
		if strs[i] == "" {
			continue
		}
		res += strs[i] + " "
	}
	//删除末尾空格
	return strings.Trim(res, " ")
}

func reverseWords1(s string) string {
	//删除收尾空格
	s = strings.Trim(s, " ")
	res := ""
	j := len(s) - 1
	i := j
	for i >= 0 {
		// 搜索首个空格
		for i >= 0 && s[i] != ' ' {
			i--
		}

		// 添加单词
		res += s[i+1:j+1] + " "

		// 跳过单词间空格
		for i >= 0 && s[i] == ' ' {
			i--
		}
		j = i // j 指向下个单词的尾字符
	}

	//删除末尾空格
	return strings.Trim(res, " ")
}

/**
class Solution {
    public String reverseWords(String s) {
        String[] strs = s.trim().split(" "); // 删除首尾空格，分割字符串
        StringBuilder res = new StringBuilder();
        for (int i = strs.length - 1; i >= 0; i--) { // 倒序遍历单词列表
            if (strs[i].equals("")) continue; // 遇到空单词则跳过
            res.append(strs[i]).append(" "); // 将单词拼接至 StringBuilder
        }
        return res.toString().trim(); // 转化为字符串，删除尾部空格，并返回
    }
}

public class Solution {
    public String reverseWords(String s) {
        s = s.trim(); // 删除首尾空格
        int j = s.length() - 1, i = j;
        StringBuilder res = new StringBuilder();
        while (i >= 0) {
            while (i >= 0 && s.charAt(i) != ' ') i--; // 搜索首个空格
            res.append(s.substring(i + 1, j + 1) + " "); // 添加单词
            while (i >= 0 && s.charAt(i) == ' ') i--; // 跳过单词间空格
            j = i; // j 指向下个单词的尾字符
        }
        return res.toString().trim(); // 转化为字符串并返回
    }
}
*/

func main() {
	reverseWords1("a good   example")
}
