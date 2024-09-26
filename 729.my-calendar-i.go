/*
 * @lc app=leetcode id=729 lang=golang
 *
 * [729] My Calendar I
 *
 * https://leetcode.com/problems/my-calendar-i/description/
 *
 * algorithms
 * Medium (56.58%)
 * Likes:    4508
 * Dislikes: 123
 * Total Accepted:    354.6K
 * Total Submissions: 615.1K
 * Testcase Example:  '["MyCalendar","book","book","book"]\n[[],[10,20],[15,25],[20,30]]'
 *
 * You are implementing a program to use as your calendar. We can add a new
 * event if adding the event will not cause a double booking.
 *
 * A double booking happens when two events have some non-empty intersection
 * (i.e., some moment is common to both events.).
 *
 * The event can be represented as a pair of integers start and end that
 * represents a booking on the half-open interval [start, end), the range of
 * real numbers x such that start <= x < end.
 *
 * Implement the MyCalendar class:
 *
 *
 * MyCalendar() Initializes the calendar object.
 * boolean book(int start, int end) Returns true if the event can be added to
 * the calendar successfully without causing a double booking. Otherwise,
 * return false and do not add the event to the calendar.
 *
 *
 *
 * Example 1:
 *
 *
 * Input
 * ["MyCalendar", "book", "book", "book"]
 * [[], [10, 20], [15, 25], [20, 30]]
 * Output
 * [null, true, false, true]
 *
 * Explanation
 * MyCalendar myCalendar = new MyCalendar();
 * myCalendar.book(10, 20); // return True
 * myCalendar.book(15, 25); // return False, It can not be booked because time
 * 15 is already booked by another event.
 * myCalendar.book(20, 30); // return True, The event can be booked, as the
 * first event takes every time less than 20, but not including 20.
 *
 *
 * Constraints:
 *
 *
 * 0 <= start < end <= 10^9
 * At most 1000 calls will be made to book.
 *
 *
 */

package main

import "fmt"

func main() {
	testCases := []struct {
		input [][]int
		output []bool
	}{
		{ input: [][]int{{10, 20}, {15, 25}, {20, 30}}, output: []bool{true, false, true} },
		{ input: [][]int{{0,0}, {10, 20}, {15, 25}, {20, 30}}, output: []bool{true, false, true, true} },
	}

	for _, testCase := range testCases {
		myCalendar := Constructor()
		res := []bool{}
		for _, input := range testCase.input {
			tRes := myCalendar.Book(input[0], input[1])
			res = append(res, tRes)
		}
		fmt.Println("Test case failed:", testCase.input, res, testCase.output)
	}
}

// @lc code=start
type MyCalendar struct {
	start 	int
	end   	int
	left 		*MyCalendar
	right		*MyCalendar
}


func Constructor() MyCalendar {
	return MyCalendar{}
}


func (this *MyCalendar) Book(start int, end int) bool {
	if this.start == 0 && this.end == 0 {
		this.start = start
		this.end = end
		return true
	}

	if start >= this.end {
		if this.right == nil {
			this.right = &MyCalendar{start: start, end: end}
			return true
		}
		return this.right.Book(start, end)
	}

	if end <= this.start {
		if this.left == nil {
			this.left = &MyCalendar{start: start, end: end}
			return true
		}
		return this.left.Book(start, end)
	}

	return false
}


/**
 * Your MyCalendar object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(start,end);
 */
// @lc code=end

