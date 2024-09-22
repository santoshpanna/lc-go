/*
 * @lc app=leetcode id=125 lang=golang
 *
 * [125] Valid Palindrome
 *
 * https://leetcode.com/problems/valid-palindrome/description/
 *
 * algorithms
 * Easy (48.69%)
 * Likes:    9559
 * Dislikes: 8422
 * Total Accepted:    3.4M
 * Total Submissions: 7M
 * Testcase Example:  '"A man, a plan, a canal: Panama"'
 *
 * A phrase is a palindrome if, after converting all uppercase letters into
 * lowercase letters and removing all non-alphanumeric characters, it reads the
 * same forward and backward. Alphanumeric characters include letters and
 * numbers.
 *
 * Given a string s, return true if it is a palindrome, or false otherwise.
 *
 *
 * Example 1:
 *
 *
 * Input: s = "A man, a plan, a canal: Panama"
 * Output: true
 * Explanation: "amanaplanacanalpanama" is a palindrome.
 *
 *
 * Example 2:
 *
 *
 * Input: s = "race a car"
 * Output: false
 * Explanation: "raceacar" is not a palindrome.
 *
 *
 * Example 3:
 *
 *
 * Input: s = " "
 * Output: true
 * Explanation: s is an empty string "" after removing non-alphanumeric
 * characters.
 * Since an empty string reads the same forward and backward, it is a
 * palindrome.
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= s.length <= 2 * 10^5
 * s consists only of printable ASCII characters.
 *
 *
 */

package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	testCases := []struct {
		input string
		expected bool
	} {
		{input: "A man, a plan, a canal: Panama", expected: true},
		{input: "race a car", expected: false},
		{input: " ", expected: true},
	}

	for _, testCase := range testCases {
		result := isPalindrome(testCase.input)
		fmt.Printf("input = %v, expected = %v, result = %v\n", testCase.input, testCase.expected, result)
	}
}

// @lc code=start
func checkPalindrome(s string) bool {
	runes := []rune(s)

	size := len(runes)

	for i, j := 0, size - 1; i < size / 2; i, j = i+1, j-1 {
		if runes[i] != runes[j] {
			return false
		}
	}

	return true
}

func getParsedString(s string) string {
	var nonAlphanumericRegex = regexp.MustCompile(`[^a-zA-Z0-9]+`)
	alphanumericString := nonAlphanumericRegex.ReplaceAllString(s, "")
	return strings.ToLower(alphanumericString)
}

func isPalindrome(s string) bool {
	parsedString := getParsedString(s)
	return checkPalindrome(parsedString)
}
// @lc code=end

