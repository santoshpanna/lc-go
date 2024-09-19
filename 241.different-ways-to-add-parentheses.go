/*
 * @lc app=leetcode id=241 lang=golang
 *
 * [241] Different Ways to Add Parentheses
 *
 * https://leetcode.com/problems/different-ways-to-add-parentheses/description/
 *
 * algorithms
 * Medium (66.02%)
 * Likes:    5903
 * Dislikes: 345
 * Total Accepted:    297.9K
 * Total Submissions: 423K
 * Testcase Example:  '"2-1-1"'
 *
 * Given a string expression of numbers and operators, return all possible
 * results from computing all the different possible ways to group numbers and
 * operators. You may return the answer in any order.
 *
 * The test cases are generated such that the output values fit in a 32-bit
 * integer and the number of different results does not exceed 10^4.
 *
 *
 * Example 1:
 *
 *
 * Input: expression = "2-1-1"
 * Output: [0,2]
 * Explanation:
 * ((2-1)-1) = 0
 * (2-(1-1)) = 2
 *
 *
 * Example 2:
 *
 *
 * Input: expression = "2*3-4*5"
 * Output: [-34,-14,-10,-10,10]
 * Explanation:
 * (2*(3-(4*5))) = -34
 * ((2*3)-(4*5)) = -14
 * ((2*(3-4))*5) = -10
 * (2*((3-4)*5)) = -10
 * (((2*3)-4)*5) = 10
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= expression.length <= 20
 * expression consists of digits and the operator '+', '-', and '*'.
 * All the integer values in the input expression are in the range [0, 99].
 * The integer values in the input expression do not have a leading '-' or '+'
 * denoting the sign.
 *
 *
 */
// package main

// import (
// 	"fmt"
// 	"strconv"
// )

// func main() {
// 	expression := "2*3-4*5"
// 	results := diffWaysToCompute(expression)
// 	fmt.Println(results)
// }

// @lc code=start
func diffWaysToCompute(expression string) []int {
	if len(expression) == 0 {
		return []int{}
	}
	if len(expression) == 1 {
		num, _ := strconv.Atoi(expression)
		return []int{num}
	}
	res := []int{}
	for i := 0; i < len(expression); i++ {
		if expression[i] == '+' || expression[i] == '-' || expression[i] == '*' {
			left := diffWaysToCompute(expression[:i])
			right := diffWaysToCompute(expression[i+1:])
			for _, l := range left {
				for _, r := range right {
					if expression[i] == '+' {
						res = append(res, l+r)
					} else if expression[i] == '-' {
						res = append(res, l-r)
					} else if expression[i] == '*' {
						res = append(res, l*r)
					}
				}
			}
		}
	}

	if len(res) == 0 {
		num, _ := strconv.Atoi(expression)
		res = append(res, num)
	}

	return res
}
// @lc code=end

