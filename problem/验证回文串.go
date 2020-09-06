package preoblem

/*
125. 验证回文串
给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。

说明：本题中，我们将空字符串定义为有效的回文串。

示例 1:

输入: "A man, a plan, a canal: Panama"
输出: true
示例 2:

输入: "race a car"
输出: false
题目来源： https://leetcode-cn.com/problems/valid-palindrome/

解题思路： 就是先整理原来的字符串， 在判断就行了
*/

func isPalindrome(s string) bool {
	var trim []rune
	for _, item := range s {
		if item >= 'a' && item <= 'z' {
			trim = append(trim, item)
		}
		if item >= 'A' && item <= 'Z' {
			trim = append(trim, 'a'+item-'A')
		}
		if item >= '0' && item <= '9' {
			trim = append(trim, item)
		}
	}

	l, r := 0, len(trim)-1
	for l < r {
		if trim[l] != trim[r] {
			return false
		}
		l++
		r--
	}
	return true
}
