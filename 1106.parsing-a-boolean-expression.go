/*
 * @lc app=leetcode id=1106 lang=golang
 *
 * [1106] Parsing A Boolean Expression
 *
 * https://leetcode.com/problems/parsing-a-boolean-expression/description/
 *
 * algorithms
 * Hard (60.03%)
 * Likes:    1380
 * Dislikes: 69
 * Total Accepted:    56.7K
 * Total Submissions: 87K
 * Testcase Example:  '"&(|(f))"'
 *
 * A boolean expression is an expression that evaluates to either true or
 * false. It can be in one of the following shapes:
 *
 *
 * 't' that evaluates to true.
 * 'f' that evaluates to false.
 * '!(subExpr)' that evaluates to the logical NOT of the inner expression
 * subExpr.
 * '&(subExpr1, subExpr2, ..., subExprn)' that evaluates to the logical AND of
 * the inner expressions subExpr1, subExpr2, ..., subExprn where n >= 1.
 * '|(subExpr1, subExpr2, ..., subExprn)' that evaluates to the logical OR of
 * the inner expressions subExpr1, subExpr2, ..., subExprn where n >= 1.
 *
 *
 * Given a string expression that represents a boolean expression, return the
 * evaluation of that expression.
 *
 * It is guaranteed that the given expression is valid and follows the given
 * rules.
 *
 *
 * Example 1:
 *
 *
 * Input: expression = "&(|(f))"
 * Output: false
 * Explanation:
 * First, evaluate |(f) --> f. The expression is now "&(f)".
 * Then, evaluate &(f) --> f. The expression is now "f".
 * Finally, return false.
 *
 *
 * Example 2:
 *
 *
 * Input: expression = "|(f,f,f,t)"
 * Output: true
 * Explanation: The evaluation of (false OR false OR false OR true) is true.
 *
 *
 * Example 3:
 *
 *
 * Input: expression = "!(&(f,t))"
 * Output: true
 * Explanation:
 * First, evaluate &(f,t) --> (false AND true) --> false --> f. The expression
 * is now "!(f)".
 * Then, evaluate !(f) --> NOT false --> true. We return true.
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= expression.length <= 2 * 10^4
 * expression[i] is one following characters: '(', ')', '&', '|', '!', 't',
 * 'f', and ','.
 *
 *
 */

package main

import "fmt"

func main() {
	testCases := []struct {
		input string
		expected bool
	} {
		{"&(t,t,t)", true},
		{"&(t,t,f)", false},
		{"&(|(f))", false},
		{"|(f,f,f,t)", true},
		{"!(&(f,t))", true},
		{"!(&(!(&(f)),&(t),|(f,f,t)))", false},
	}

	for _, tc := range testCases {
		fmt.Printf("Expected: %v, got: %v\n", tc.expected, parseBoolExpr(tc.input))
	}
}

// @lc code=start
type Stack struct {
	stack		[]rune
	top			int
}

func (this *Stack) push(val rune) {
	this.stack = append(this.stack, val)
	this.top++
}

func (this *Stack) pop() rune {
	if this.top != -1 {
		val := this.stack[this.top]
		this.stack = this.stack[:this.top]
		this.top--
		return val
	}

	return 0
}

func (this *Stack) peek() rune {
	if this.top != -1 {
		return this.stack[this.top]
	}

	return 0
}

func (this *Stack) isEmpty() bool {
	if this.top == -1 {
		return true
	}

	return false
}

func parseBoolByte(b rune) bool {
	return b == 't'
}

func convertBoolToByte(b bool) rune {
	if b == true {
		return 't'
	}

	return 'f'
}

func parseBoolExpr(expression string) bool {
	modeStack := Stack{
		stack: make([]rune, 0),
		top: -1,
	}

	valStack := Stack{
		stack: make([]rune, 0),
		top: -1,
	}

	for _, ch := range expression {
		if ch == '&' || ch == '|' || ch == '!' {
			modeStack.push(ch)
		} else if ch == '(' {
			valStack.push('(')
		} else if ch == 't' || ch == 'f' {
			valStack.push(ch)
		} else if ch == ')' {
			var vals []bool
			for valStack.peek() != '(' {
				vals = append([]bool{parseBoolByte(valStack.pop())}, vals...)
			}
			valStack.pop()

			mode := modeStack.pop()
			var result bool
			if mode == '&' {
				result = true
				for _, val := range vals {
					result = result && val
				}
			} else if mode == '|' {
				result = false
				for _, val := range vals {
					result = result || val
				}
			} else if mode == '!' {
				result = !vals[0]
			}
			valStack.push(convertBoolToByte(result))
		}
	}

	return parseBoolByte(valStack.pop())
}
// @lc code=end

