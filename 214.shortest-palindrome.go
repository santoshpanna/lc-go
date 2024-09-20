/*
 * @lc app=leetcode id=214 lang=golang
 *
 * [214] Shortest Palindrome
 *
 * https://leetcode.com/problems/shortest-palindrome/description/
 *
 * algorithms
 * Hard (34.03%)
 * Likes:    4009
 * Dislikes: 257
 * Total Accepted:    236.9K
 * Total Submissions: 638K
 * Testcase Example:  '"aacecaaa"'
 *
 * You are given a string s. You can convert s to a palindrome by adding
 * characters in front of it.
 *
 * Return the shortest palindrome you can find by performing this
 * transformation.
 *
 *
 * Example 1:
 * Input: s = "aacecaaa"
 * Output: "aaacecaaa"
 * Example 2:
 * Input: s = "abcd"
 * Output: "dcbabcd"
 *
 *
 * Constraints:
 *
 *
 * 0 <= s.length <= 5 * 10^4
 * s consists of lowercase English letters only.
 *
 *
 */
package main

import (
	"fmt"
)

func main() {
	tests := []struct {
		input    string
		expected string
	}{
		{"aacecaaa", "aaacecaaa"},
		{"abcd", "dcbabcd"},
		{"", ""},
		{"a", "a"},
		{"aa", "aa"},
		{"ab", "bab"},
		{"racecar", "racecar"},
		{"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"},
		{"aabba", "abbaabba"},
	}

	for _, test := range tests {
		result := shortestPalindrome(test.input)
		if result == test.expected {
			fmt.Printf("PASS: Input: %s, Output: %s\n", test.input, result)
		} else {
			fmt.Printf("FAIL: Input: %s, Expected: %s, Got: %s\n", test.input, test.expected, result)
		}
	}
}

// @lc code=start
func reverse(s string) string {
	runes := []rune(s)
	size := len(s)

	for i, j := 0, size - 1; i < size / 2; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}

func isPalindrome(s1 string) bool {
	s2 := reverse(s1)

	if s1 == s2 {
		return true
	}

	return false
}

func getPrefix(s string) string {
	size := len(s)
	rev := reverse(s)

	for i := 0; i < size; i++ {
		ns := s[:size - i]
		if ns == rev[i:] {
			return reverse(s[size - i:])
		}
	}

	return ""
}

func shortestPalindrome(s string) string {
	s1 := getPrefix(s)
	return s1+s
}
// @lc code=end

