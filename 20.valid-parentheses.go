/*
 * @lc app=leetcode id=20 lang=golang
 *
 * [20] Valid Parentheses
 *
 * https://leetcode.com/problems/valid-parentheses/description/
 *
 * algorithms
 * Easy (41.02%)
 * Likes:    24459
 * Dislikes: 1808
 * Total Accepted:    5.1M
 * Total Submissions: 12.5M
 * Testcase Example:  '"()"'
 *
 * Given a string s containing just the characters '(', ')', '{', '}', '[' and
 * ']', determine if the input string is valid.
 * 
 * An input string is valid if:
 * 
 * 
 * Open brackets must be closed by the same type of brackets.
 * Open brackets must be closed in the correct order.
 * Every close bracket has a corresponding open bracket of the same type.
 * 
 * 
 * 
 * Example 1:
 * 
 * 
 * Input: s = "()"
 * 
 * Output: true
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: s = "()[]{}"
 * 
 * Output: true
 * 
 * 
 * Example 3:
 * 
 * 
 * Input: s = "(]"
 * 
 * Output: false
 * 
 * 
 * Example 4:
 * 
 * 
 * Input: s = "([])"
 * 
 * Output: true
 * 
 * 
 * 
 * Constraints:
 * 
 * 
 * 1 <= s.length <= 10^4
 * s consists of parentheses only '()[]{}'.
 * 
 * 
 */

package main

import "fmt"

func main() {
	testCases := []struct {
		s string
		expected bool
	}{
		{"()", true},
		{"()[]{}", true},
		{"(]", false},
		{"([])", true},
		{"([)]", false},
		{"{[]}", true},
		{"{", false},
		{"}", false},
		{"{[}", false},
		{"{[]", false},
		{"{[]}", true},
		{"{[()]}", true},
		{"{[()]}{[()]}", true},
	}

	for _, testCase := range testCases {
		result := isValid(testCase.s)
		fmt.Printf("expected: %v, result: %v\n", testCase.expected, result)
	}
}

// @lc code=start
func isValid(s string) bool {
	stack := make([]rune, 0)

	for _, c := range s {
		if c == '(' || c == '[' || c == '{' {
			stack = append(stack, c)
		} else {
			if len(stack) == 0 {
				return false
			}

			if c == ')' && stack[len(stack)-1] != '(' {
				return false
			}

			if c == ']' && stack[len(stack)-1] != '[' {
				return false
			}

			if c == '}' && stack[len(stack)-1] != '{' {
				return false
			}

			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}
// @lc code=end

