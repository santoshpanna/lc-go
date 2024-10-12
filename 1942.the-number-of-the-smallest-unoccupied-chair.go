/*
 * @lc app=leetcode id=1942 lang=golang
 *
 * [1942] The Number of the Smallest Unoccupied Chair
 *
 * https://leetcode.com/problems/the-number-of-the-smallest-unoccupied-chair/description/
 *
 * algorithms
 * Medium (42.63%)
 * Likes:    1086
 * Dislikes: 55
 * Total Accepted:    68.3K
 * Total Submissions: 119.4K
 * Testcase Example:  '[[1,4],[2,3],[4,6]]\n1'
 *
 * There is a party where n friends numbered from 0 to n - 1 are attending.
 * There is an infinite number of chairs in this party that are numbered from 0
 * to infinity. When a friend arrives at the party, they sit on the unoccupied
 * chair with the smallest number.
 *
 *
 * For example, if chairs 0, 1, and 5 are occupied when a friend comes, they
 * will sit on chair number 2.
 *
 *
 * When a friend leaves the party, their chair becomes unoccupied at the moment
 * they leave. If another friend arrives at that same moment, they can sit in
 * that chair.
 *
 * You are given a 0-indexed 2D integer array times where times[i] = [arrivali,
 * leavingi], indicating the arrival and leaving times of the i^th friend
 * respectively, and an integer targetFriend. All arrival times are distinct.
 *
 * Return the chair number that the friend numbered targetFriend will sit
 * on.
 *
 *
 * Example 1:
 *
 *
 * Input: times = [[1,4],[2,3],[4,6]], targetFriend = 1
 * Output: 1
 * Explanation:
 * - Friend 0 arrives at time 1 and sits on chair 0.
 * - Friend 1 arrives at time 2 and sits on chair 1.
 * - Friend 1 leaves at time 3 and chair 1 becomes empty.
 * - Friend 0 leaves at time 4 and chair 0 becomes empty.
 * - Friend 2 arrives at time 4 and sits on chair 0.
 * Since friend 1 sat on chair 1, we return 1.
 *
 *
 * Example 2:
 *
 *
 * Input: times = [[3,10],[1,5],[2,6]], targetFriend = 0
 * Output: 2
 * Explanation:
 * - Friend 1 arrives at time 1 and sits on chair 0.
 * - Friend 2 arrives at time 2 and sits on chair 1.
 * - Friend 0 arrives at time 3 and sits on chair 2.
 * - Friend 1 leaves at time 5 and chair 0 becomes empty.
 * - Friend 2 leaves at time 6 and chair 1 becomes empty.
 * - Friend 0 leaves at time 10 and chair 2 becomes empty.
 * Since friend 0 sat on chair 2, we return 2.
 *
 *
 *
 * Constraints:
 *
 *
 * n == times.length
 * 2 <= n <= 10^4
 * times[i].length == 2
 * 1 <= arrivali < leavingi <= 10^5
 * 0 <= targetFriend <= n - 1
 * Each arrivali time is distinct.
 *
 *
 */

package main

import (
	"container/heap"
	"fmt"
	"sort"
)

func main() {
	testCases := []struct {
		times [][]int
		targetFriend int
		expected int
	}{
		{times: [][]int{{1,4},{2,3},{4,6}}, targetFriend: 1, expected: 1},
		{times: [][]int{{3,10},{1,5},{2,6}}, targetFriend: 0, expected: 2},
	}
	
	for _, testCase := range testCases {
		result := smallestChair(testCase.times, testCase.targetFriend)
		fmt.Printf("expected: %d, got: %d\n", testCase.expected, result)
	}
}

// @lc code=start
type Event struct {
	time  int
	index int
	isArr bool
}

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

func smallestChair(times [][]int, targetFriend int) int {
	n := len(times)
	events := make([]Event, 0, 2*n)
	for i, time := range times {
		events = append(events, Event{time[0], i, true})
		events = append(events, Event{time[1], i, false})
	}

	sort.Slice(events, func(i, j int) bool {
		if events[i].time == events[j].time {
			return !events[i].isArr && events[j].isArr
		}
		return events[i].time < events[j].time
	})

	freeChairs := &MinHeap{}
	heap.Init(freeChairs)
	for i := 0; i < n; i++ {
		heap.Push(freeChairs, i)
	}

	occupiedChairs := make([]int, n)
	for _, event := range events {
		if event.isArr {
			chair := heap.Pop(freeChairs).(int)
			occupiedChairs[event.index] = chair
			if event.index == targetFriend {
				return chair
			}
		} else {
			heap.Push(freeChairs, occupiedChairs[event.index])
		}
	}
	return -1
}
// @lc code=end

