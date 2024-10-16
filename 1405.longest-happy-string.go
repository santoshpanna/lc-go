/*
 * @lc app=leetcode id=1405 lang=golang
 *
 * [1405] Longest Happy String
 *
 * https://leetcode.com/problems/longest-happy-string/description/
 *
 * algorithms
 * Medium (57.52%)
 * Likes:    2441
 * Dislikes: 289
 * Total Accepted:    145.6K
 * Total Submissions: 230.9K
 * Testcase Example:  '1\n1\n7'
 *
 * A string s is called happy if it satisfies the following conditions:
 *
 *
 * s only contains the letters 'a', 'b', and 'c'.
 * s does not contain any of "aaa", "bbb", or "ccc" as a substring.
 * s contains at most a occurrences of the letter 'a'.
 * s contains at most b occurrences of the letter 'b'.
 * s contains at most c occurrences of the letter 'c'.
 *
 *
 * Given three integers a, b, and c, return the longest possible happy string.
 * If there are multiple longest happy strings, return any of them. If there is
 * no such string, return the empty string "".
 *
 * A substring is a contiguous sequence of characters within a string.
 *
 *
 * Example 1:
 *
 *
 * Input: a = 1, b = 1, c = 7
 * Output: "ccaccbcc"
 * Explanation: "ccbccacc" would also be a correct answer.
 *
 *
 * Example 2:
 *
 *
 * Input: a = 7, b = 1, c = 0
 * Output: "aabaa"
 * Explanation: It is the only correct answer in this case.
 *
 *
 *
 * Constraints:
 *
 *
 * 0 <= a, b, c <= 100
 * a + b + c > 0
 *
 *
 */

package main

import (
	"fmt"
	"sort"
)

func main() {
	testCases := []struct {
		a int
		b int
		c int
		expected string
	}{
		{1, 1, 7, "ccaccbcc"},
		{7, 1, 0, "aabaa"},
		{1, 1, 7, "ccaccbcc"},
		{2, 2, 1, "aabbc"},
		{7, 1, 0, "aabaa"},
		{0, 0, 1, "c"},
	}

	for _, testCase := range testCases {
		result := longestDiverseString(testCase.a, testCase.b, testCase.c)
		fmt.Printf("expected: %v, result: %v\n", testCase.expected, result)
	}
}

// @lc code=start
func longestDiverseString(a int, b int, c int) string {
	type charCount struct {
		count int
		char  byte
	}

	chars := []charCount{
		{a, 'a'},
		{b, 'b'},
		{c, 'c'},
	}

	result := []byte{}

	for {
		sort.Slice(chars, func(i, j int) bool {
			return chars[i].count > chars[j].count
		})

		added := false
		for i := 0; i < 3; i++ {
			if chars[i].count == 0 {
				break
			}
			n := len(result)
			if n >= 2 && result[n-1] == chars[i].char && result[n-2] == chars[i].char {
				continue
			}
			result = append(result, chars[i].char)
			chars[i].count--
			added = true
			break
		}

		if !added {
			break
		}
	}

	return string(result)
}
// @lc code=end

