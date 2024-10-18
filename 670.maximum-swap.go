/*
 * @lc app=leetcode id=670 lang=golang
 *
 * [670] Maximum Swap
 *
 * https://leetcode.com/problems/maximum-swap/description/
 *
 * algorithms
 * Medium (49.13%)
 * Likes:    3877
 * Dislikes: 247
 * Total Accepted:    352.9K
 * Total Submissions: 695.5K
 * Testcase Example:  '2736'
 *
 * You are given an integer num. You can swap two digits at most once to get
 * the maximum valued number.
 *
 * Return the maximum valued number you can get.
 *
 *
 * Example 1:
 *
 *
 * Input: num = 2736
 * Output: 7236
 * Explanation: Swap the number 2 and the number 7.
 *
 *
 * Example 2:
 *
 *
 * Input: num = 9973
 * Output: 9973
 * Explanation: No swap.
 *
 *
 *
 * Constraints:
 *
 *
 * 0 <= num <= 10^8
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
		num int
		expected int
	} {
		{98368, 98863},
		// {6687, 8667},
		// {2736, 7236},
		// {9973, 9973},
		// {1993, 9913},
		// {98368, 98863},
		// {56867898, 58667898},
		// {98004365, 98004365},
		// {98765432, 98765432},
		// {98800435, 98800435},
	}

	for _, testCase := range testCases {
		result := maximumSwap(testCase.num)
		fmt.Printf("expected = %v, result = %v\n", testCase.expected, result)
	}
}

// @lc code=start
func maximumSwap(num int) int {
	numStr := []byte(fmt.Sprintf("%d", num))
	last := make([]int, 10)
	for i := range numStr {
		last[numStr[i]-'0'] = i
	}
	for i := range numStr {
		for d := 9; d > int(numStr[i]-'0'); d-- {
			if last[d] > i {
				numStr[i], numStr[last[d]] = numStr[last[d]], numStr[i]
				result, _ := strconv.Atoi(string(numStr))
				return result
			}
		}
	}
	result, _ := strconv.Atoi(string(numStr))
	return result
}

// @lc code=end

