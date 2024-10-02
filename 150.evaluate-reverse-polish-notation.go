/*
 * @lc app=leetcode id=150 lang=golang
 *
 * [150] Evaluate Reverse Polish Notation
 *
 * https://leetcode.com/problems/evaluate-reverse-polish-notation/description/
 *
 * algorithms
 * Medium (52.67%)
 * Likes:    7785
 * Dislikes: 1107
 * Total Accepted:    1.2M
 * Total Submissions: 2.2M
 * Testcase Example:  '["2","1","+","3","*"]'
 *
 * You are given an array of strings tokens that represents an arithmetic
 * expression in a Reverse Polish Notation.
 *
 * Evaluate the expression. Return an integer that represents the value of the
 * expression.
 *
 * Note that:
 *
 *
 * The valid operators are '+', '-', '*', and '/'.
 * Each operand may be an integer or another expression.
 * The division between two integers always truncates toward zero.
 * There will not be any division by zero.
 * The input represents a valid arithmetic expression in a reverse polish
 * notation.
 * The answer and all the intermediate calculations can be represented in a
 * 32-bit integer.
 *
 *
 *
 * Example 1:
 *
 *
 * Input: tokens = ["2","1","+","3","*"]
 * Output: 9
 * Explanation: ((2 + 1) * 3) = 9
 *
 *
 * Example 2:
 *
 *
 * Input: tokens = ["4","13","5","/","+"]
 * Output: 6
 * Explanation: (4 + (13 / 5)) = 6
 *
 *
 * Example 3:
 *
 *
 * Input: tokens = ["10","6","9","3","+","-11","*","/","*","17","+","5","+"]
 * Output: 22
 * Explanation: ((10 * (6 / ((9 + 3) * -11))) + 17) + 5
 * = ((10 * (6 / (12 * -11))) + 17) + 5
 * = ((10 * (6 / -132)) + 17) + 5
 * = ((10 * 0) + 17) + 5
 * = (0 + 17) + 5
 * = 17 + 5
 * = 22
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= tokens.length <= 10^4
 * tokens[i] is either an operator: "+", "-", "*", or "/", or an integer in the
 * range [-200, 200].
 *
 *
 */

package main

import (
	"fmt"
	"strconv"
)

func main() {
	testCases := []struct {
		tokens []string
		expected int
	}{
		{[]string{"2","1","+","3","*"}, 9},
		{[]string{"4","13","5","/","+"}, 6},
		{[]string{"10","6","9","3","+","-11","*","/","*","17","+","5","+"}, 22},
	}

	for _, testCase := range testCases {
		result := evalRPN(testCase.tokens)
		fmt.Printf("expected: %v, result: %v\n", testCase.expected, result)
	}
}

// @lc code=start
type Stack struct {
	stack			[]int
	top				int
}

func Constructor() Stack {
	return Stack{
		stack: make([]int, 0),
		top: -1,
	}
}

func (this *Stack) pop() int {
	this.top--
	val := this.stack[this.top]
	this.stack = this.stack[:this.top]

	return val
}

func (this *Stack) push(val int) {
	this.stack = append(this.stack, val)
	this.top++
} 

func evalRPN(tokens []string) int {
	stack := Stack{}

	for i := 0; i < len(tokens); i++ {
		if tokens[i] == "+" || tokens[i] == "-" || tokens[i] == "*" || tokens[i] == "/" {
			n1 := stack.pop()
			n2 := stack.pop()

			res := 0

			switch tokens[i] {
				case "+":
					res = n1 + n2
					break
				case "-":
					res = n2 - n1
					break
				case "/":
					res = n2 / n1
					break
				case "*":
					res = n1 * n2
					break
			}

			stack.push(res)
		} else {
			num,_ := strconv.Atoi(tokens[i])
			stack.push(num)
		}
	}

	return stack.pop()
}
// @lc code=end

