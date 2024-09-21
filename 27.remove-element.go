/*
 * @lc app=leetcode id=27 lang=golang
 *
 * [27] Remove Element
 *
 * https://leetcode.com/problems/remove-element/description/
 *
 * algorithms
 * Easy (58.14%)
 * Likes:    2858
 * Dislikes: 3939
 * Total Accepted:    3.4M
 * Total Submissions: 5.8M
 * Testcase Example:  '[3,2,2,3]\n3'
 *
 * Given an integer array nums and an integer val, remove all occurrences of
 * val in nums in-place. The order of the elements may be changed. Then return
 * the number of elements in nums which are not equal to val.
 *
 * Consider the number of elements in nums which are not equal to val be k, to
 * get accepted, you need to do the following things:
 *
 *
 * Change the array nums such that the first k elements of nums contain the
 * elements which are not equal to val. The remaining elements of nums are not
 * important as well as the size of nums.
 * Return k.
 *
 *
 * Custom Judge:
 *
 * The judge will test your solution with the following code:
 *
 *
 * int[] nums = [...]; // Input array
 * int val = ...; // Value to remove
 * int[] expectedNums = [...]; // The expected answer with correct length.
 * ⁠                           // It is sorted with no values equaling val.
 *
 * int k = removeElement(nums, val); // Calls your implementation
 *
 * assert k == expectedNums.length;
 * sort(nums, 0, k); // Sort the first k elements of nums
 * for (int i = 0; i < actualLength; i++) {
 * ⁠   assert nums[i] == expectedNums[i];
 * }
 *
 *
 * If all assertions pass, then your solution will be accepted.
 *
 *
 * Example 1:
 *
 *
 * Input: nums = [3,2,2,3], val = 3
 * Output: 2, nums = [2,2,_,_]
 * Explanation: Your function should return k = 2, with the first two elements
 * of nums being 2.
 * It does not matter what you leave beyond the returned k (hence they are
 * underscores).
 *
 *
 * Example 2:
 *
 *
 * Input: nums = [0,1,2,2,3,0,4,2], val = 2
 * Output: 5, nums = [0,1,4,0,3,_,_,_]
 * Explanation: Your function should return k = 5, with the first five elements
 * of nums containing 0, 0, 1, 3, and 4.
 * Note that the five elements can be returned in any order.
 * It does not matter what you leave beyond the returned k (hence they are
 * underscores).
 *
 *
 *
 * Constraints:
 *
 *
 * 0 <= nums.length <= 100
 * 0 <= nums[i] <= 50
 * 0 <= val <= 100
 *
 *
 */

package main

import "fmt"

func main() {
	testCases := []struct {
		nums				[]int
		val 				int
		expected 		int
	} {
		// {nums: []int{3,2,2,3}, val: 3, expected: 2},
		// {nums: []int{0,1,2,2,3,0,4,2}, val: 2, expected: 5},
		// {nums: []int{0,1,2,2,3,0,4,2}, val: 0, expected: 6},
		{nums: []int{1}, val: 1, expected: 0},
		{nums: []int{2,2,3}, val: 2, expected: 1},
	}

	for _, tc := range testCases {
		fmt.Println("Input =", tc.nums, "val =", tc.val, "expected =", tc.expected, "result =", removeElement(tc.nums, tc.val))
	}
}

// @lc code=start
func removeElement(nums []int, val int) int {
	size := len(nums)
	ptr := size - 1

	if size == 0 {
		return 0
	}

	for nums[ptr] == val {
		if ptr == 0 {
			break
		}
		ptr--
		
	}

	for i := 0; i < size; i++ {
		if nums[i] == val {
			if i < ptr {
				nums[i], nums[ptr] = nums[ptr], nums[i]

				for ptr > 0 && nums[ptr] == val {
					ptr--
				}
			}
		}
	}

	for ptr = 0; ptr < size; ptr++ {
		if nums[ptr] == val {
			break
		}
	}

	return ptr
}
// @lc code=end

