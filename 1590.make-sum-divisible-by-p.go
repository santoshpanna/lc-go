/*
 * @lc app=leetcode id=1590 lang=golang
 *
 * [1590] Make Sum Divisible by P
 *
 * https://leetcode.com/problems/make-sum-divisible-by-p/description/
 *
 * algorithms
 * Medium (29.21%)
 * Likes:    2045
 * Dislikes: 115
 * Total Accepted:    83.4K
 * Total Submissions: 229.8K
 * Testcase Example:  '[3,1,4,2]\n6'
 *
 * Given an array of positive integers nums, remove the smallest subarray
 * (possibly empty) such that the sum of the remaining elements is divisible by
 * p. It is not allowed to remove the whole array.
 *
 * Return the length of the smallest subarray that you need to remove, or -1 if
 * it's impossible.
 *
 * A subarray is defined as a contiguous block of elements in the array.
 *
 *
 * Example 1:
 *
 *
 * Input: nums = [3,1,4,2], p = 6
 * Output: 1
 * Explanation: The sum of the elements in nums is 10, which is not divisible
 * by 6. We can remove the subarray [4], and the sum of the remaining elements
 * is 6, which is divisible by 6.
 *
 *
 * Example 2:
 *
 *
 * Input: nums = [6,3,5,2], p = 9
 * Output: 2
 * Explanation: We cannot remove a single element to get a sum divisible by 9.
 * The best way is to remove the subarray [5,2], leaving us with [6,3] with sum
 * 9.
 *
 *
 * Example 3:
 *
 *
 * Input: nums = [1,2,3], p = 3
 * Output: 0
 * Explanation: Here the sum is 6. which is already divisible by 3. Thus we do
 * not need to remove anything.
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= nums.length <= 10^5
 * 1 <= nums[i] <= 10^9
 * 1 <= p <= 10^9
 *
 *
 */
package main

import "fmt"

func main() {
	testCases := []struct {
		nums []int
		p int
		output int
	}{
		{
			nums: []int{3,1,4,2},
			p: 6,
			output: 1,
		},
		{
			nums: []int{6,3,5,2},
			p: 9,
			output: 2,
		},
		{
			nums: []int{1,2,3},
			p: 3,
			output: 0,
		},
	}

	for _, testCase := range testCases {
		output := minSubarray(testCase.nums, testCase.p)
		fmt.Println(output == testCase.output)
	}
}

// @lc code=start
func minSubarray(nums []int, p int) int {
	sum := 0

	for _, num := range nums {
		sum = (sum + num) % p
	}

	rem := sum % p

	if rem == 0 {
			return 0
	}

	modMap := make(map[int] int)
	modMap[0] = -1
	currentSum, minLen := 0, len(nums)

	for i := 0; i < len(nums); i++ {
		currentSum = (currentSum + nums[i]) % p
		previousSum := (currentSum - rem + p) % p;

		if val, ok := modMap[previousSum]; ok {
			minLen = min(minLen, i - val)
		}

		modMap[currentSum] = i
	}

	if len(nums) == minLen {
		return -1
	}

	return minLen
}
// @lc code=end

