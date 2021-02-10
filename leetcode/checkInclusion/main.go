package checkInclusion

/**
给定两个字符串s1和s2，写一个函数来判断 s2 是否包含 s1的排列。
换句话说，第一个字符串的排列之一是第二个字符串的子串。

示例1:
输入: s1 = "ab" s2 = "eidbaooo"
输出: True
解释: s2 包含 s1 的排列之一 ("ba").

示例2:
输入: s1= "ab" s2 = "eidboaoo"
输出: False

注意：
输入的字符串只包含小写字母
两个字符串的长度都在 [1, 10,000] 之间

*/
func checkInclusion(s1 string, s2 string) bool {
	dirS1, dirS2 := [26]int{}, [26]int{}
	n := len(s1)
	if n > len(s2) {
		return false
	}
	for i := range s1 {
		dirS1[s1[i]-'a']++
	}
	for i := range s2[:n] {
		dirS2[s2[i]-'a']++
	}
	if dirS1 == dirS2 {
		return true
	}
	left := 0
	right := n - 1
	for right < len(s2)-1 {
		right++
		dirS2[s2[right]-'a']++
		dirS2[s2[left]-'a']--
		left++
		if dirS2 == dirS1 {
			return true
		}
	}
	return false
}

func main() {

}
