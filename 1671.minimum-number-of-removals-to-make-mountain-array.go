/*
 * @lc app=leetcode id=1671 lang=golang
 *
 * [1671] Minimum Number of Removals to Make Mountain Array
 *
 * https://leetcode.com/problems/minimum-number-of-removals-to-make-mountain-array/description/
 *
 * algorithms
 * Hard (42.22%)
 * Likes:    2019
 * Dislikes: 36
 * Total Accepted:    87.4K
 * Total Submissions: 163.9K
 * Testcase Example:  '[1,3,1]'
 *
 * You may recall that an array arr is a mountain array if and only if:
 *
 *
 * arr.length >= 3
 * There exists some index i (0-indexed) with 0 < i < arr.length - 1 such
 * that:
 *
 * arr[0] < arr[1] < ... < arr[i - 1] < arr[i]
 * arr[i] > arr[i + 1] > ... > arr[arr.length - 1]
 *
 *
 *
 *
 * Given an integer array nums​​​, return the minimum number of elements to
 * remove to make nums​​​ a mountain array.
 *
 *
 * Example 1:
 *
 *
 * Input: nums = [1,3,1]
 * Output: 0
 * Explanation: The array itself is a mountain array so we do not need to
 * remove any elements.
 *
 *
 * Example 2:
 *
 *
 * Input: nums = [2,1,1,5,6,2,3,1]
 * Output: 3
 * Explanation: One solution is to remove the elements at indices 0, 1, and 5,
 * making the array nums = [1,5,6,3,1].
 *
 *
 *
 * Constraints:
 *
 *
 * 3 <= nums.length <= 1000
 * 1 <= nums[i] <= 10^9
 * It is guaranteed that you can make a mountain array out of nums.
 *
 *
 */

package main

import (
	"fmt"
	"math"
)

func main() {
	nums := []int{2,1,1,5,6,2,3,1}
	fmt.Println(minimumMountainRemovals(nums))
}

// @lc code=start
func minimumMountainRemovals(nums []int) int {
	l := len(nums)

	lis := make([]int, l)
	lds := make([]int, l)

	for i := 0; i < l; i++ {
		lis[i] = 1
		lds[i] = 1
	}

	for i := 0; i < l; i++ {	
		for j := i-1; j >= 0; j-- {
			if nums[i] > nums[j] {
				lis[i] = max(lis[i], lis[j] + 1) 
			}
		}
	}

	for i := l - 1; i >= 0; i-- {
		for j := i+1; j < l; j++ {
			if nums[i] > nums[j]{
				lds[i] = max(lds[i], lds[j] + 1)
			}
		}
	}

	res := math.MaxInt32

	for i := 0; i < l; i++ {
		if lis[i] > 1 && lds[i] > 1 {
			res = min(res, l - lis[i] - lds[i] + 1)
		}
	}

	return res
}
// @lc code=end

