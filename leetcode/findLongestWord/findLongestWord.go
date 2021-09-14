package main

import "sort"

/**
给你一个字符串 s 和一个字符串数组 dictionary 作为字典，找出并返回字典中最长的字符串，该字符串可以通过删除 s 中的某些字符得到。
如果答案不止一个，返回长度最长且字典序最小的字符串。如果答案不存在，则返回空字符串。

示例 1：
输入：s = "abpcplea", dictionary = ["ale","apple","monkey","plea"]
输出："apple"

示例 2：
输入：s = "abpcplea", dictionary = ["a","b","c"]
输出："a"

提示：
1 <= s.length <= 1000
1 <= dictionary.length <= 1000
1 <= dictionary[i].length <= 1000
s 和 dictionary[i] 仅由小写英文字母组成
*/

func findLongestWord(s string, dictionary []string) string {
	sort.Slice(dictionary, func(i, j int) bool {
		a, b := dictionary[i], dictionary[j]
		return len(a) > len(b) || len(a) == len(b) && a < b
	})
	for _, t := range dictionary {
		i := 0
		j := 0
		for i < len(s) && j < len(t) {
			if s[i] == t[j] {
				j++
			}
			i++
		}
		if j == len(t) {
			return t
		}
	}

	return ""
}

/** java
class Solution {
    public String findLongestWord(String s, List<String> dictionary) {
        String res = "";
        int sn = s.length();
        //1.排序
        Collections.sort(dictionary, (a,b)->{
             if (a.length() != b.length()) {
                 return b.length() - a.length();
             }
            return a.compareTo(b);
        });

        //2.循环双指针
        for (String ss : dictionary) {
            int ssn = ss.length();
            int a = 0;
            int b = 0;
            while (a < sn && b < ssn) {
                if (s.charAt(a) == ss.charAt(b)) {
                    b++;
                }
                a++;
            }

            if (b == ssn) {
                res = ss;
                break;
            }

        }
        return res;
    }
}

*/
func main() {

}
