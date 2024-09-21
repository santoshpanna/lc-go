/*
 * @lc app=leetcode id=205 lang=golang
 *
 * [205] Isomorphic Strings
 *
 * https://leetcode.com/problems/isomorphic-strings/description/
 *
 * algorithms
 * Easy (45.70%)
 * Likes:    9169
 * Dislikes: 2126
 * Total Accepted:    1.5M
 * Total Submissions: 3.3M
 * Testcase Example:  '"egg"\n"add"'
 *
 * Given two strings s and t, determine if they are isomorphic.
 *
 * Two strings s and t are isomorphic if the characters in s can be replaced to
 * get t.
 *
 * All occurrences of a character must be replaced with another character while
 * preserving the order of characters. No two characters may map to the same
 * character, but a character may map to itself.
 *
 *
 * Example 1:
 *
 *
 * Input: s = "egg", t = "add"
 *
 * Output: true
 *
 * Explanation:
 *
 * The strings s and t can be made identical by:
 *
 *
 * Mapping 'e' to 'a'.
 * Mapping 'g' to 'd'.
 *
 *
 *
 * Example 2:
 *
 *
 * Input: s = "foo", t = "bar"
 *
 * Output: false
 *
 * Explanation:
 *
 * The strings s and t can not be made identical as 'o' needs to be mapped to
 * both 'a' and 'r'.
 *
 *
 * Example 3:
 *
 *
 * Input: s = "paper", t = "title"
 *
 * Output: true
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= s.length <= 5 * 10^4
 * t.length == s.length
 * s and t consist of any valid ascii character.
 *
 *
 */
package main

import "fmt"

func main() {
	testCases := []struct {
		s						string
		t						string
		output			bool
	} {
		{s: "egg", t: "add", output: true},
		{s: "foo", t: "bar", output: false},
		{s: "paper", t: "title", output: true},
		{s: "badc", t: "baba", output: false},
		{s: "bbbaaaba", t: "aaabbbba", output: false},
	}

	for _, tc := range testCases {
		fmt.Println("inputs =", tc.s, tc.t, "result =", isIsomorphic(tc.s, tc.t), "expected output =", tc.output, "\n\n\n\n")
	}
}

// @lc code=start
func isIsomorphic(s string, t string) bool {
	// if len(s) != len(t) {
	// 	return false
	// }

	srunes := []rune(s)
	trunes := []rune(t)

	lookup := make(map [rune]rune)
	seenMap := make(map [rune]bool)

	for i := 0; i < len(srunes); i++ {
		val, exist := lookup[srunes[i]]
		_, seen := seenMap[trunes[i]]
		
		// fmt.Println("exists and seen", exist, seen)
		if exist && seen == false {
			// fmt.Println("found in lookup", srunes[i], lookup, i, trunes[i], val)
			if val != trunes[i] {
				return false
			}
		} else if seen == false {
			lookup[srunes[i]] = trunes[i]
			seenMap[trunes[i]] = true
			// fmt.Println("putting into lookup", srunes[i], trunes[i], lookup)
		} else if exist == false && seen {
			return false
		} else if exist && seen && val != trunes[i] {
			return false
		}
	}

	return true
}
// @lc code=end

