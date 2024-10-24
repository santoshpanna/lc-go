/*
 * @lc app=leetcode id=1593 lang=golang
 *
 * [1593] Split a String Into the Max Number of Unique Substrings
 *
 * https://leetcode.com/problems/split-a-string-into-the-max-number-of-unique-substrings/description/
 *
 * algorithms
 * Medium (57.07%)
 * Likes:    988
 * Dislikes: 44
 * Total Accepted:    52.5K
 * Total Submissions: 85.2K
 * Testcase Example:  '"ababccc"'
 *
 * Given a string s, return the maximum number of unique substrings that the
 * given string can be split into.
 * 
 * You can split string s into any list of non-empty substrings, where the
 * concatenation of the substrings forms the original string. However, you must
 * split the substrings such that all of them are unique.
 * 
 * A substring is a contiguous sequence of characters within a string.
 * 
 * 
 * Example 1:
 * 
 * 
 * Input: s = "ababccc"
 * Output: 5
 * Explanation: One way to split maximally is ['a', 'b', 'ab', 'c', 'cc'].
 * Splitting like ['a', 'b', 'a', 'b', 'c', 'cc'] is not valid as you have 'a'
 * and 'b' multiple times.
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: s = "aba"
 * Output: 2
 * Explanation: One way to split maximally is ['a', 'ba'].
 * 
 * 
 * Example 3:
 * 
 * 
 * Input: s = "aa"
 * Output: 1
 * Explanation: It is impossible to split the string any further.
 * 
 * 
 * 
 * Constraints:
 * 
 * 
 * 
 * 1 <= s.length <= 16
 * 
 * 
 * s contains only lower case English letters.
 * 
 * 
 * 
 */
package main

import "fmt"

func main() {
	testCases := []struct{
		s string
		want int
	}{
		{s: "ababccc", want: 5},
		{s: "aba", want: 2},
		{s: "aa", want: 1},
	}

	for _, tc := range testCases {
		got := maxUniqueSplit(tc.s)
		fmt.Println(got == tc.want)
	}
}

// @lc code=start
func maxUniqueSplit(s string) int {
	seen := make(map[string]bool)
	return backtrack(s, seen)
}

func backtrack(s string, seen map[string]bool) int {
	if len(s) == 0 {
		return 0
	}

	maxCount := 0
	for i := 1; i <= len(s); i++ {
		substr := s[:i]
		if !seen[substr] {
			seen[substr] = true
			count := 1 + backtrack(s[i:], seen)
			if count > maxCount {
				maxCount = count
			}
			delete(seen, substr)
		}
	}
	return maxCount
}
// @lc code=end

