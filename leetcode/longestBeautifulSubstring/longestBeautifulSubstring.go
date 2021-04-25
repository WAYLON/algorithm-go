package main

/**
当一个字符串满足如下条件时，我们称它是 美丽的：
所有 5 个英文元音字母（'a'，'e'，'i'，'o'，'u'）都必须至少出现一次。
这些元音字母的顺序都必须按照 字典序升序排布（也就是说所有的 'a'都在 'e'前面，所有的 'e'都在 'i'前面，以此类推）
比方说，字符串"aeiou" 和"aaaaaaeiiiioou"都是 美丽的，但是"uaeio"，"aeoiu"和"aaaeeeooo"不是美丽的。
给你一个只包含英文元音字母的字符串word，请你返回word 中 最长美丽子字符串的长度。如果不存在这样的子字符串，请返回 0。
子字符串 是字符串中一个连续的字符序列。


示例 1：
输入：word = "aeiaaioaaaaeiiiiouuuooaauuaeiu"
输出：13
解释：最长子字符串是 "aaaaeiiiiouuu" ，长度为 13 。

示例 2：
输入：word = "aeeeiiiioooauuuaeiou"
输出：5
解释：最长子字符串是 "aeiou" ，长度为 5 。

示例 3：
输入：word = "a"
输出：0
解释：没有美丽子字符串，所以返回 0 。

提示：
1 <= word.length <= 5 * 105
word只包含字符'a'，'e'，'i'，'o'和'u'。
*/

/**
思路：判断字符种类和大小
这5个元音字符是升序的且字符串中仅包含这些元音字母，所以当字符串满足整体升序的同时字符种类达到5,那么这个字符串就是美丽的。
时间复杂度：O(n) 空间复杂度：O(1)
*/
func longestBeautifulSubstring(word string) int {
	ans := 0
	ty := 1
	l := 1
	for i := 1; i < len(word); i++ {
		if word[i] >= word[i-1] {
			l++
		}
		if word[i] > word[i-1] {
			ty++
		}
		if word[i] < word[i-1] {
			ty = 1
			l = 1
		}
		if ty == 5 {
			ans = max(ans, l)
		}
	}
	return ans
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

/**
思路：判断字符种类和大小
这5个元音字符是升序的且字符串中仅包含这些元音字母，所以当字符串满足整体升序的同时字符种类达到5,那么这个字符串就是美丽的。
时间复杂度：O(n) 空间复杂度：O(1)
class Solution {
    public int longestBeautifulSubstring(String word) {
        int ans = 0, type = 1, len = 1;
        for (int i = 1; i < word.length(); i++) {
            if (word.charAt(i) >= word.charAt(i - 1)) len++; //更新当前字符串长度
            if (word.charAt(i) > word.charAt(i - 1)) type++; //更新当前字符种类
            if (word.charAt(i) < word.charAt(i - 1)) {
                type = 1;
                len = 1;
            } //当前字符串不美丽，从当前字符重新开始
            if (type == 5) ans = Math.max(ans, len);  //更新最大字符串
        }
        return ans;
    }
}
*/
func main() {

}
