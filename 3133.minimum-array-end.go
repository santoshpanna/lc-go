/*
 * @lc app=leetcode id=3133 lang=golang
 *
 * [3133] Minimum Array End
 *
 * https://leetcode.com/problems/minimum-array-end/description/
 *
 * algorithms
 * Medium (38.22%)
 * Likes:    537
 * Dislikes: 55
 * Total Accepted:    54.1K
 * Total Submissions: 105.1K
 * Testcase Example:  '3\n4'
 *
 * You are given two integers n and x. You have to construct an array of
 * positive integers nums of size n where for every 0 <= i < n - 1, nums[i + 1]
 * is greater than nums[i], and the result of the bitwise AND operation between
 * all elements of nums is x.
 *
 * Return the minimum possible value of nums[n - 1].
 *
 *
 * Example 1:
 *
 *
 * Input: n = 3, x = 4
 *
 * Output: 6
 *
 * Explanation:
 *
 * nums can be [4,5,6] and its last element is 6.
 *
 *
 * Example 2:
 *
 *
 * Input: n = 2, x = 7
 *
 * Output: 15
 *
 * Explanation:
 *
 * nums can be [7,15] and its last element is 15.
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= n, x <= 10^8
 *
 *
 */
package main

import (
	"fmt"
)

func main() {
	testCases := []struct{
		n		int
		x		int
		res int
	} {
		{
			n: 3, x: 4, res: 6,
		},{
			n: 2, x: 7, res: 15,
		},{
			n: 21, x: 16, res: 52,
		},
	}

	for _, tc := range testCases {
		fmt.Println("res of n x", tc.n, tc.x, "=", minEnd(tc.n, tc.x), "and expected res =", tc.res)
	}
}

// @lc code=start
// func minEnd(n int, x int) int64 {
// 	if n == 0 {
// 		return 0
// 	}
	
// 	lastNum := strconv.FormatInt(int64(x), 2)

// 	for i := 1; i < n; i++ {
// 		numToProcess := lastNum

// 		l := len(numToProcess)-1

// 		last0bit := l

// 		for last0bit > 0 && numToProcess[last0bit] != '0' {
// 			last0bit--
// 		}
		
// 		if last0bit == 0 && numToProcess[0] != '0' {
// 			if num, err := strconv.ParseInt("1"+numToProcess, 2, 64); err == nil {
// 				if num & int64(x) == int64(x) {
// 					lastNum = strconv.FormatInt(num, 2)
// 				}
// 			}
// 		} else if num, err := strconv.ParseInt(numToProcess[:last0bit]+"1"+numToProcess[last0bit:l], 2, 64); err == nil {	
// 			if num & int64(x) == int64(x) {
// 				lastNum = strconv.FormatInt(num, 2)
// 			}
// 		}
// 	}

// 	if num, err := strconv.ParseInt(lastNum, 2, 64); err == nil {
// 		return num
// 	}

// 	return int64(x)
// }

func minEnd(n int, x int) int64 {
	result := int64(x)
	mask := int64(1)
	n--

	for n > 0 {
		if (mask & int64(x)) == 0 {
			result |= (int64(n) & 1) * mask
			n >>= 1
		}
		mask <<= 1
	}

	return result
}
// @lc code=end

