/*
 * @lc app=leetcode id=2707 lang=golang
 *
 * [2707] Extra Characters in a String
 *
 * https://leetcode.com/problems/extra-characters-in-a-string/description/
 *
 * algorithms
 * Medium (52.56%)
 * Likes:    2402
 * Dislikes: 125
 * Total Accepted:    144.6K
 * Total Submissions: 257.3K
 * Testcase Example:  '"leetscode"\n["leet","code","leetcode"]'
 *
 * You are given a 0-indexed string s and a dictionary of words dictionary. You
 * have to break s into one or more non-overlapping substrings such that each
 * substring is present in dictionary. There may be some extra characters in s
 * which are not present in any of the substrings.
 *
 * Return the minimum number of extra characters left over if you break up s
 * optimally.
 *
 *
 * Example 1:
 *
 *
 * Input: s = "leetscode", dictionary = ["leet","code","leetcode"]
 * Output: 1
 * Explanation: We can break s in two substrings: "leet" from index 0 to 3 and
 * "code" from index 5 to 8. There is only 1 unused character (at index 4), so
 * we return 1.
 *
 *
 *
 * Example 2:
 *
 *
 * Input: s = "sayhelloworld", dictionary = ["hello","world"]
 * Output: 3
 * Explanation: We can break s in two substrings: "hello" from index 3 to 7 and
 * "world" from index 8 to 12. The characters at indices 0, 1, 2 are not used
 * in any substring and thus are considered as extra characters. Hence, we
 * return 3.
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= s.length <= 50
 * 1 <= dictionary.length <= 50
 * 1 <= dictionary[i].length <= 50
 * dictionary[i] and s consists of only lowercase English letters
 * dictionary contains distinct words
 *
 *
 */
package main

import (
	"fmt"
)

func main() {
	testCases := []struct {
		s string
		dictionary []string
		expected int
	}{
		{s: "leetscode", dictionary: []string{"leet","code","leetcode"}, expected: 1},
		{s: "sayhelloworld", dictionary: []string{"hello","world"}, expected: 3},
	}

	for _, testCase := range testCases {
		testCase := testCase
		result := minExtraChar(testCase.s, testCase.dictionary)
		fmt.Printf("expected: %v, returned: %v\n", testCase.expected, result)
	}
}

// @lc code=start
func minExtraChar(s string, dictionary []string) int {
	n := len(s)
	dp := make([]int, n+1)
	dp[0] = 0
	for i := 1; i <= n; i++ {
		dp[i] = dp[i-1] + 1
		for _, word := range dictionary {
			if i >= len(word) && s[i-len(word):i] == word {
				dp[i] = min(dp[i], dp[i-len(word)])
			}
		}
	}

	return dp[n]
}
// @lc code=end

