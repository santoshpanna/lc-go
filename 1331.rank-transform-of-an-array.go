/*
 * @lc app=leetcode id=1331 lang=golang
 *
 * [1331] Rank Transform of an Array
 *
 * https://leetcode.com/problems/rank-transform-of-an-array/description/
 *
 * algorithms
 * Easy (62.64%)
 * Likes:    1951
 * Dislikes: 97
 * Total Accepted:    180.1K
 * Total Submissions: 267.6K
 * Testcase Example:  '[40,10,20,30]'
 *
 * Given an array of integers arr, replace each element with its rank.
 * 
 * The rank represents how large the element is. The rank has the following
 * rules:
 * 
 * 
 * Rank is an integer starting from 1.
 * The larger the element, the larger the rank. If two elements are equal,
 * their rank must be the same.
 * Rank should be as small as possible.
 * 
 * 
 * 
 * Example 1:
 * 
 * 
 * Input: arr = [40,10,20,30]
 * Output: [4,1,2,3]
 * Explanation: 40 is the largest element. 10 is the smallest. 20 is the second
 * smallest. 30 is the third smallest.
 * 
 * Example 2:
 * 
 * 
 * Input: arr = [100,100,100]
 * Output: [1,1,1]
 * Explanation: Same elements share the same rank.
 * 
 * 
 * Example 3:
 * 
 * 
 * Input: arr = [37,12,28,9,100,56,80,5,12]
 * Output: [5,3,4,2,8,6,7,1,3]
 * 
 * 
 * 
 * Constraints:
 * 
 * 
 * 0 <= arr.length <= 10^5
 * -10^9 <= arr[i] <= 10^9
 * 
 * 
 */

// @lc code=start
import (
	"sort"
)

func arrayRankTransform(arr []int) []int {
	size := len(arr)

	rankMap := make(map[int]int)

	temp := make([]int, size)
	copy(temp, arr)

	sort.Slice(temp, func(i, j int) bool {
			return temp[i] < temp[j]
	})

	dup := 0
	for i := 0; i < size; i++ {
			if i != 0 && temp[i] == temp[i-1] {
					dup++
			}
			if _, ok := rankMap[temp[i]]; !ok {
					rankMap[temp[i]] = i - dup
			}
	}

	res := make([]int, size)

	for i := 0; i < size; i++ {
			res[i] = rankMap[arr[i]] + 1
	}

	return res
}
// @lc code=end

