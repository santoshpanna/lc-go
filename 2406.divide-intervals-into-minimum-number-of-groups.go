/*
 * @lc app=leetcode id=2406 lang=golang
 *
 * [2406] Divide Intervals Into Minimum Number of Groups
 *
 * https://leetcode.com/problems/divide-intervals-into-minimum-number-of-groups/description/
 *
 * algorithms
 * Medium (47.53%)
 * Likes:    1099
 * Dislikes: 24
 * Total Accepted:    72K
 * Total Submissions: 122.8K
 * Testcase Example:  '[[5,10],[6,8],[1,5],[2,3],[1,10]]'
 *
 * You are given a 2D integer array intervals where intervals[i] = [lefti,
 * righti] represents the inclusive interval [lefti, righti].
 * 
 * You have to divide the intervals into one or more groups such that each
 * interval is in exactly one group, and no two intervals that are in the same
 * group intersect each other.
 * 
 * Return the minimum number of groups you need to make.
 * 
 * Two intervals intersect if there is at least one common number between them.
 * For example, the intervals [1, 5] and [5, 8] intersect.
 * 
 * 
 * Example 1:
 * 
 * 
 * Input: intervals = [[5,10],[6,8],[1,5],[2,3],[1,10]]
 * Output: 3
 * Explanation: We can divide the intervals into the following groups:
 * - Group 1: [1, 5], [6, 8].
 * - Group 2: [2, 3], [5, 10].
 * - Group 3: [1, 10].
 * It can be proven that it is not possible to divide the intervals into fewer
 * than 3 groups.
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: intervals = [[1,3],[5,6],[8,10],[11,13]]
 * Output: 1
 * Explanation: None of the intervals overlap, so we can put all of them in one
 * group.
 * 
 * 
 * 
 * Constraints:
 * 
 * 
 * 1 <= intervals.length <= 10^5
 * intervals[i].length == 2
 * 1 <= lefti <= righti <= 10^6
 * 
 * 
 */

// @lc code=start
// import (
// 	"sort"
// )

// func minGroups(intervals [][]int) int {
// 	sort.Slice(intervals, func(i, j int) bool {
// 		return intervals[i][0] < intervals[j][0]
// 	})

// 	endTimes := []int{}

// 	for _, interval := range intervals {
// 		start, end := interval[0], interval[1]

// 		if len(endTimes) > 0 && endTimes[0] < start {
// 			endTimes = endTimes[1:]
// 		}

// 		endTimes = append(endTimes, end)

// 		sort.Ints(endTimes)
// 	}

// 	return len(endTimes)
// }
import (
	"container/heap"
)

type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func minGroups(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	h := &MinHeap{}
	heap.Init(h)

	for _, interval := range intervals {
		start, end := interval[0], interval[1]

		if h.Len() > 0 && (*h)[0] < start {
			heap.Pop(h)
		}

		heap.Push(h, end)
	}

	return h.Len()
}
// @lc code=end

