/*
 * @lc app=leetcode id=49 lang=golang
 *
 * [49] Group Anagrams
 *
 * https://leetcode.com/problems/group-anagrams/description/
 *
 * algorithms
 * Medium (69.37%)
 * Likes:    19495
 * Dislikes: 637
 * Total Accepted:    3.2M
 * Total Submissions: 4.6M
 * Testcase Example:  '["eat","tea","tan","ate","nat","bat"]'
 *
 * Given an array of strings strs, group the anagrams together. You can return
 * the answer in any order.
 *
 *
 * Example 1:
 *
 *
 * Input: strs = ["eat","tea","tan","ate","nat","bat"]
 *
 * Output: [["bat"],["nat","tan"],["ate","eat","tea"]]
 *
 * Explanation:
 *
 *
 * There is no string in strs that can be rearranged to form "bat".
 * The strings "nat" and "tan" are anagrams as they can be rearranged to form
 * each other.
 * The strings "ate", "eat", and "tea" are anagrams as they can be rearranged
 * to form each other.
 *
 *
 *
 * Example 2:
 *
 *
 * Input: strs = [""]
 *
 * Output: [[""]]
 *
 *
 * Example 3:
 *
 *
 * Input: strs = ["a"]
 *
 * Output: [["a"]]
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= strs.length <= 10^4
 * 0 <= strs[i].length <= 100
 * strs[i] consists of lowercase English letters.
 *
 *
 */
package main

import (
	"fmt"
	"sort"
)

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	result := groupAnagrams(strs)
	fmt.Println(result)
}

// @lc code=start
func groupAnagrams(strs []string) [][]string {
	groups := make(map [string][]string)

	for _, str := range strs {
		sorted := []rune(str)
		sort.Slice(sorted, func(i int, j int) bool {
			return sorted[i] < sorted[j]
		})
		groups[string(sorted)] = append(groups[string(sorted)], str)
	}

	res := [][]string{}

	for _, value := range groups {
		res = append(res, value)
	}

	return res
}
// @lc code=end

