/*
 * @lc app=leetcode id=1229 lang=golang
 *
 * [1229] Meeting Scheduler
 *
 * https://leetcode.com/problems/meeting-scheduler/description/
 *
 * algorithms
 * Medium (55.31%)
 * Likes:    917
 * Dislikes: 36
 * Total Accepted:    84.7K
 * Total Submissions: 153K
 * Testcase Example:  '[[10,50],[60,120],[140,210]]\n[[0,15],[60,70]]\n8'
 *
 * Given the availability time slots arrays slots1 and slots2 of two people and
 * a meeting duration duration, return the earliest time slot that works for
 * both of them and is of duration duration.
 *
 * If there is no common time slot that satisfies the requirements, return an
 * empty array.
 *
 * The format of a time slot is an array of two elements [start, end]
 * representing an inclusive time range from start to end.
 *
 * It is guaranteed that no two availability slots of the same person intersect
 * with each other. That is, for any two time slots [start1, end1] and [start2,
 * end2] of the same person, either start1 > end2 or start2 > end1.
 *
 *
 * Example 1:
 *
 *
 * Input: slots1 = [[10,50],[60,120],[140,210]], slots2 = [[0,15],[60,70]],
 * duration = 8
 * Output: [60,68]
 *
 *
 * Example 2:
 *
 *
 * Input: slots1 = [[10,50],[60,120],[140,210]], slots2 = [[0,15],[60,70]],
 * duration = 12
 * Output: []
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= slots1.length, slots2.length <= 10^4
 * slots1[i].length, slots2[i].length == 2
 * slots1[i][0] < slots1[i][1]
 * slots2[i][0] < slots2[i][1]
 * 0 <= slots1[i][j], slots2[i][j] <= 10^9
 * 1 <= duration <= 10^6
 *
 *
 */
package main

import (
	"fmt"
	"sort"
)

func main() {
	testCases := []struct {
		slots1   [][]int
		slots2   [][]int
		duration int
		output   []int
	}{
		{
			slots1:   [][]int{{10, 50}, {60, 120}, {140, 210}},
			slots2:   [][]int{{0, 15}, {60, 70}},
			duration: 8,
			output:   []int{60, 68},
		},
		{
			slots1:   [][]int{{10, 50}, {60, 120}, {140, 210}},
			slots2:   [][]int{{0, 15}, {60, 70}},
			duration: 12,
			output:   []int{},
		},
		{
			slots1:   [][]int{{10, 60}},
			slots2:   [][]int{{12, 17}, {21, 50}},
			duration: 8,
			output:   []int{21, 29},
		},
		{
			slots1:   [][]int{{216397070,363167701},{98730764,158208909},{441003187,466254040},{558239978,678368334},{683942980,717766451}},
			slots2:   [][]int{{50490609,222653186},{512711631,670791418},{730229023,802410205},{812553104,891266775},{230032010,399152578}},
			duration: 456085,
			output:   []int{98730764, 99186849},
		},
	}

	for _, testCase := range testCases {
		res := minAvailableDuration(testCase.slots1, testCase.slots2, testCase.duration)
		fmt.Println("Test case:", testCase.slots1, testCase.slots2, testCase.duration, res, testCase.output)
	}
}

// @lc code=start
func minAvailableDuration(slots1 [][]int, slots2 [][]int, duration int) []int {
	sort.Slice(slots1, func(i, j int) bool {
		return slots1[i][0] < slots1[j][0]
	})

	sort.Slice(slots2, func(i, j int) bool {
		return slots2[i][0] < slots2[j][0]
	})

	i, j := 0, 0

	for i < len(slots1) && j < len(slots2) {
		start := max(slots1[i][0], slots2[j][0])
		end := min(slots1[i][1], slots2[j][1])

		if end-start >= duration {
			return []int{start, start + duration}
		}

		if slots1[i][1] < slots2[j][1] {
			i++
		} else {
			j++
		}
	}

	return []int{}
}
// @lc code=end

