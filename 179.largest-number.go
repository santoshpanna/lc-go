/*
 * @lc app=leetcode id=179 lang=golang
 *
 * [179] Largest Number
 *
 * https://leetcode.com/problems/largest-number/description/
 *
 * algorithms
 * Medium (36.93%)
 * Likes:    8355
 * Dislikes: 699
 * Total Accepted:    540.9K
 * Total Submissions: 1.4M
 * Testcase Example:  '[10,2]'
 *
 * Given a list of non-negative integers nums, arrange them such that they form
 * the largest number and return it.
 *
 * Since the result may be very large, so you need to return a string instead
 * of an integer.
 *
 *
 * Example 1:
 *
 *
 * Input: nums = [10,2]
 * Output: "210"
 *
 *
 * Example 2:
 *
 *
 * Input: nums = [3,30,34,5,9]
 * Output: "9534330"
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= nums.length <= 100
 * 0 <= nums[i] <= 10^9
 *
 *
 */

// package main

// import (
// 	"fmt"
// 	"sort"
// 	"strconv"
// )

// func main() {
// 	tests := []struct {
// 		nums     []int
// 		expected string
// 	}{
// 		{[]int{2, 10}, "210"},
// 		{[]int{3, 30, 34, 5, 9}, "9534330"},
// 		{[]int{1}, "1"},
// 		{[]int{10}, "10"},
// 		{[]int{0, 0}, "0"},
// 		{[]int{111311, 1113}, "1113111311"},
// 		{[]int{0,9,8,7,6,5,4,3,2,1}, "9876543210"},
// 	}

// 	for _, test := range tests {
// 		result := largestNumber(test.nums)
// 		fmt.Println("Result for ", test, "is", result)
// 		if result != test.expected {
// 			fmt.Println("Error", result)
// 		}
// 	}
// }

// func factorial(num int) int {
// 	if num == 1 {
// 		return 1
// 	}

// 	fac := 1

// 	for i := 1; i <= num; i++ {
// 		fac *= i
// 	}

// 	return fac
// }

// @lc code=start
func largestNumber(nums []int) string {
	sort.Slice(nums, func(i, j int) bool {
		str1 := strconv.Itoa(nums[i]) + strconv.Itoa(nums[j])
		str2 := strconv.Itoa(nums[j]) + strconv.Itoa(nums[i])
		return str1 > str2
	})

	if nums[0] == 0 {
		return "0"
	}

	result := ""

	for i := range nums {
		result += strconv.Itoa(nums[i])
	}

	return result
}
// @lc code=end

