/*
 * @lc app=leetcode id=347 lang=golang
 *
 * [347] Top K Frequent Elements
 *
 * https://leetcode.com/problems/top-k-frequent-elements/description/
 *
 * algorithms
 * Medium (63.28%)
 * Likes:    17531
 * Dislikes: 675
 * Total Accepted:    2.3M
 * Total Submissions: 3.7M
 * Testcase Example:  '[1,1,1,2,2,3]\n2'
 *
 * Given an integer array nums and an integer k, return the k most frequent
 * elements. You may return the answer in any order.
 *
 *
 * Example 1:
 * Input: nums = [1,1,1,2,2,3], k = 2
 * Output: [1,2]
 * Example 2:
 * Input: nums = [1], k = 1
 * Output: [1]
 *
 *
 * Constraints:
 *
 *
 * 1 <= nums.length <= 10^5
 * -10^4 <= nums[i] <= 10^4
 * k is in the range [1, the number of unique elements in the array].
 * It is guaranteed that the answer is unique.
 *
 *
 *
 * Follow up: Your algorithm's time complexity must be better than O(n log n),
 * where n is the array's size.
 *
 */
package main

import (
	"fmt"
	"sort"
)

func main() {
	testCases := []struct {
		nums     []int
		k        int
		expected []int
	}{
		{nums: []int{1, 1, 1, 2, 2, 3}, k: 2, expected: []int{1, 2}},
		{nums: []int{1}, k: 1, expected: []int{1}},
		{nums: []int{1,2}, k: 2, expected: []int{1,2}},
		{nums: []int{4, 4, 4, 5, 5, 5, 5, 6, 6, 7}, k: 3, expected: []int{5, 4, 6}},
	}

	for _, tc := range testCases {
		result := topKFrequent(tc.nums, tc.k)
		fmt.Printf("nums: %v, k: %d, expected: %v, got: %v\n", tc.nums, tc.k, tc.expected, result)
	}
}

// @lc code=start
func topKFrequent(nums []int, k int) []int {
	freqMap := make(map [int]int)
	
	for i := 0; i < len(nums); i++ {
		freqMap[nums[i]]++
	}

	res := []int{}

	for k := range freqMap {
		res = append(res, k)
	}

	sort.Slice(res, func(i int, j int) bool {
		return freqMap[res[i]] > freqMap[res[j]]
	})

	return res[:k]
}
// @lc code=end

