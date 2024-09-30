/*
 * @lc app=leetcode id=3 lang=golang
 *
 * [3] Longest Substring Without Repeating Characters
 *
 * https://leetcode.com/problems/longest-substring-without-repeating-characters/description/
 *
 * algorithms
 * Medium (35.42%)
 * Likes:    40278
 * Dislikes: 1934
 * Total Accepted:    6.4M
 * Total Submissions: 17.9M
 * Testcase Example:  '"abcabcbb"'
 *
 * Given a string s, find the length of the longest substring without repeating
 * characters.
 * 
 * 
 * Example 1:
 * 
 * 
 * Input: s = "abcabcbb"
 * Output: 3
 * Explanation: The answer is "abc", with the length of 3.
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: s = "bbbbb"
 * Output: 1
 * Explanation: The answer is "b", with the length of 1.
 * 
 * 
 * Example 3:
 * 
 * 
 * Input: s = "pwwkew"
 * Output: 3
 * Explanation: The answer is "wke", with the length of 3.
 * Notice that the answer must be a substring, "pwke" is a subsequence and not
 * a substring.
 * 
 * 
 * 
 * Constraints:
 * 
 * 
 * 0 <= s.length <= 5 * 10^4
 * s consists of English letters, digits, symbols and spaces.
 * 
 * 
 */

package main

import "fmt"

func main() {
	testCases := []struct {
		input  string
		output int
	}{
		{
			input:  "abcabcbb",
			output: 3,
		},
		{
			input:  "bbbbb",
			output: 1,
		},
		{
			input:  "pwwkew",
			output: 3,
		},
	}

	for _, testCase := range testCases {
		res := lengthOfLongestSubstring(testCase.input)
		fmt.Println("Input: ", testCase.input, " expected output: ", testCase.output, " output: ", res)
	}
}

// @lc code=start
func lengthOfLongestSubstring(s string) int {
	lookup := make(map [rune]int)
	j, res := 0, 0

	for i, char := range s {
		if _, ok := lookup[char]; ok {
			j = max(j, lookup[char] + 1)
		}

		lookup[char] = i

		res = max(res, i - j + 1)
	}

	return res
}
// @lc code=end

