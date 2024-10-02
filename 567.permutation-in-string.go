/*
 * @lc app=leetcode id=567 lang=golang
 *
 * [567] Permutation in String
 *
 * https://leetcode.com/problems/permutation-in-string/description/
 *
 * algorithms
 * Medium (44.78%)
 * Likes:    11408
 * Dislikes: 431
 * Total Accepted:    939.3K
 * Total Submissions: 2.1M
 * Testcase Example:  '"ab"\n"eidbaooo"'
 *
 * Given two strings s1 and s2, return true if s2 contains a permutation of s1,
 * or false otherwise.
 *
 * In other words, return true if one of s1's permutations is the substring of
 * s2.
 *
 *
 * Example 1:
 *
 *
 * Input: s1 = "ab", s2 = "eidbaooo"
 * Output: true
 * Explanation: s2 contains one permutation of s1 ("ba").
 *
 *
 * Example 2:
 *
 *
 * Input: s1 = "ab", s2 = "eidboaoo"
 * Output: false
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= s1.length, s2.length <= 10^4
 * s1 and s2 consist of lowercase English letters.
 *
 *
 */
package main

import "fmt"

func main() {
	testCases := []struct {
		s1 string
		s2 string
		expected bool
	}{
		{"ab", "eidbaooo", true},
		{"ab", "eidboaoo", false},
		{"hello", "ooolleoooleh", false},
		{"abcd", "abcaabdbcdeafgbcdiabcbbcaabclcabcababccabd", true},
		{"adc", "dcda", true},
	}

	for _, testCase := range testCases {
		result := checkInclusion(testCase.s1, testCase.s2)
		fmt.Printf("expected: %v, result: %v\n", testCase.expected, result)
	}
}

// @lc code=start
func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	s1arr := make([]int, 26)

	for i := 0; i < len(s1); i++ {
		s1arr[s1[i] - 'a']++
	}

	for j := 0; j <= len(s2) - len(s1); j++ {
		s2arr := make([]int, 26)

		for i := 0; i < len(s1); i++ {
			s2arr[s2[i+j] - 'a']++
		}

		res := true

		for i := 0; i < len(s1); i++ {
			if s1arr[s1[i] - 'a'] != s2arr[s1[i] - 'a'] {
				res = false
				break
			}
		}

		if res == true {
			return true
		}
	}

	return false
}
// @lc code=end

