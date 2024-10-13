/*
 * @lc app=leetcode id=632 lang=golang
 *
 * [632] Smallest Range Covering Elements from K Lists
 *
 * https://leetcode.com/problems/smallest-range-covering-elements-from-k-lists/description/
 *
 * algorithms
 * Hard (63.43%)
 * Likes:    3756
 * Dislikes: 82
 * Total Accepted:    146.1K
 * Total Submissions: 220K
 * Testcase Example:  '[[4,10,15,24,26],[0,9,12,20],[5,18,22,30]]'
 *
 * You have k lists of sorted integers in non-decreasing order. Find the
 * smallest range that includes at least one number from each of the k lists.
 * 
 * We define the range [a, b] is smaller than range [c, d] if b - a < d - c or
 * a < c if b - a == d - c.
 * 
 * 
 * Example 1:
 * 
 * 
 * Input: nums = [[4,10,15,24,26],[0,9,12,20],[5,18,22,30]]
 * Output: [20,24]
 * Explanation: 
 * List 1: [4, 10, 15, 24,26], 24 is in range [20,24].
 * List 2: [0, 9, 12, 20], 20 is in range [20,24].
 * List 3: [5, 18, 22, 30], 22 is in range [20,24].
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: nums = [[1,2,3],[1,2,3],[1,2,3]]
 * Output: [1,1]
 * 
 * 
 * 
 * Constraints:
 * 
 * 
 * nums.length == k
 * 1 <= k <= 3500
 * 1 <= nums[i].length <= 50
 * -10^5 <= nums[i][j] <= 10^5
 * nums[i] is sorted in non-decreasing order.
 * 
 * 
 */

// @lc code=start
import (
	"container/heap"
	"math"
)

type Element struct {
	value int
	row   int
	col   int
}

type MinHeap []Element

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].value < h[j].value }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(Element))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func smallestRange(nums [][]int) []int {
	minHeap := &MinHeap{}
	heap.Init(minHeap)
	maxValue := math.MinInt32

	// Initialize the heap with the first element of each list
	for i := 0; i < len(nums); i++ {
		heap.Push(minHeap, Element{value: nums[i][0], row: i, col: 0})
		if nums[i][0] > maxValue {
			maxValue = nums[i][0]
		}
	}

	start, end := 0, math.MaxInt32

	for minHeap.Len() == len(nums) {
		minElement := heap.Pop(minHeap).(Element)
		if maxValue-minElement.value < end-start {
			start, end = minElement.value, maxValue
		}

		if minElement.col+1 < len(nums[minElement.row]) {
			nextElement := Element{
				value: nums[minElement.row][minElement.col+1],
				row:   minElement.row,
				col:   minElement.col + 1,
			}
			heap.Push(minHeap, nextElement)
			if nextElement.value > maxValue {
				maxValue = nextElement.value
			}
		} else {
			break
		}
	}

	return []int{start, end}
}
// @lc code=end

