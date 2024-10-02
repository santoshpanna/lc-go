/*
 * @lc app=leetcode id=1497 lang=golang
 *
 * [1497] Check If Array Pairs Are Divisible by k
 *
 * https://leetcode.com/problems/check-if-array-pairs-are-divisible-by-k/description/
 *
 * algorithms
 * Medium (38.48%)
 * Likes:    2270
 * Dislikes: 137
 * Total Accepted:    136.8K
 * Total Submissions: 304.2K
 * Testcase Example:  '[1,2,3,4,5,10,6,7,8,9]\n5'
 *
 * Given an array of integers arr of even length n and an integer k.
 *
 * We want to divide the array into exactly n / 2 pairs such that the sum of
 * each pair is divisible by k.
 *
 * Return true If you can find a way to do that or false otherwise.
 *
 *
 * Example 1:
 *
 *
 * Input: arr = [1,2,3,4,5,10,6,7,8,9], k = 5
 * Output: true
 * Explanation: Pairs are (1,9),(2,8),(3,7),(4,6) and (5,10).
 *
 *
 * Example 2:
 *
 *
 * Input: arr = [1,2,3,4,5,6], k = 7
 * Output: true
 * Explanation: Pairs are (1,6),(2,5) and(3,4).
 *
 *
 * Example 3:
 *
 *
 * Input: arr = [1,2,3,4,5,6], k = 10
 * Output: false
 * Explanation: You can try all possible pairs to see that there is no way to
 * divide arr into 3 pairs each with sum divisible by 10.
 *
 *
 *
 * Constraints:
 *
 *
 * arr.length == n
 * 1 <= n <= 10^5
 * n is even.
 * -10^9 <= arr[i] <= 10^9
 * 1 <= k <= 10^5
 *
 *
 */

package main

import (
	"fmt"
)

func main() {
	testCases := []struct {
		input []int
		k int
		expected bool
	}{
		{input: []int{1,2,3,4,5,10,6,7,8,9}, k: 5, expected: true},
		{input: []int{1,2,3,4,5,6}, k: 7, expected: true},
		{input: []int{1,2,3,4,5,6}, k: 10, expected: false},
	}

	for _, testCase := range testCases {
		testCase := testCase
		result := canArrange(testCase.input, testCase.k)
		fmt.Printf("expected: %v, result: %v\n", testCase.expected, result)
	}
}

// @lc code=start
func canArrange(arr []int, k int) bool {
	freqMap := make([]int, k)

	for i := 0; i < len(arr); i++ {
		hint1 := ((arr[i] % k) + k) % k
		freqMap[hint1]++
	}

	if (freqMap[0] % 2) != 0 {
		return false
	}

	for i := 1; i <= (k / 2); i++ {
		if freqMap[i] != freqMap[k-i] {
			return false
		}
	}
	
	return true
}
// @lc code=end

