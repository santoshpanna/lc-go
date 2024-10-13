/*
 * @lc app=leetcode id=22 lang=golang
 *
 * [22] Generate Parentheses
 *
 * https://leetcode.com/problems/generate-parentheses/description/
 *
 * algorithms
 * Medium (75.47%)
 * Likes:    21382
 * Dislikes: 976
 * Total Accepted:    2M
 * Total Submissions: 2.7M
 * Testcase Example:  '3'
 *
 * Given n pairs of parentheses, write a function to generate all combinations
 * of well-formed parentheses.
 * 
 * 
 * Example 1:
 * Input: n = 3
 * Output: ["((()))","(()())","(())()","()(())","()()()"]
 * Example 2:
 * Input: n = 1
 * Output: ["()"]
 * 
 * 
 * Constraints:
 * 
 * 
 * 1 <= n <= 8
 * 
 * 
 */
package main

import "fmt"

func main() {
	testCases := []struct {
		input int
		output []string
	}{
		{
			input: 3,
			output: []string{"((()))","(()())","(())()","()(())","()()()"},
		},
		{
			input: 1,
			output: []string{"()"},
		},
	}

	for _, testCase := range testCases {
		result := generateParenthesis(testCase.input)
		fmt.Println("Expected:", testCase.output)
		fmt.Println("Output:", result)
	}
}

// @lc code=start
func generateParenthesis(n int) []string {
	var result []string
	var backtrack func(open, close int, current string)
	backtrack = func(open, close int, current string) {
		if len(current) == n*2 {
			result = append(result, current)
			return
		}
		if open < n {
			backtrack(open+1, close, current+"(")
		}
		if close < open {
			backtrack(open, close+1, current+")")
		}
	}
	backtrack(0, 0, "")
	return result
}
// @lc code=end

