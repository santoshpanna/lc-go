/*
 * @lc app=leetcode id=155 lang=golang
 *
 * [155] Min Stack
 *
 * https://leetcode.com/problems/min-stack/description/
 *
 * algorithms
 * Medium (54.96%)
 * Likes:    14377
 * Dislikes: 891
 * Total Accepted:    1.9M
 * Total Submissions: 3.4M
 * Testcase Example:  '["MinStack","push","push","push","getMin","pop","top","getMin"]\n' +
  '[[],[-2],[0],[-3],[],[],[],[]]'
 *
 * Design a stack that supports push, pop, top, and retrieving the minimum
 * element in constant time.
 *
 * Implement the MinStack class:
 *
 *
 * MinStack() initializes the stack object.
 * void push(int val) pushes the element val onto the stack.
 * void pop() removes the element on the top of the stack.
 * int top() gets the top element of the stack.
 * int getMin() retrieves the minimum element in the stack.
 *
 *
 * You must implement a solution with O(1) time complexity for each
 * function.
 *
 *
 * Example 1:
 *
 *
 * Input
 * ["MinStack","push","push","push","getMin","pop","top","getMin"]
 * [[],[-2],[0],[-3],[],[],[],[]]
 *
 * Output
 * [null,null,null,null,-3,null,0,-2]
 *
 * Explanation
 * MinStack minStack = new MinStack();
 * minStack.push(-2);
 * minStack.push(0);
 * minStack.push(-3);
 * minStack.getMin(); // return -3
 * minStack.pop();
 * minStack.top();    // return 0
 * minStack.getMin(); // return -2
 *
 *
 *
 * Constraints:
 *
 *
 * -2^31 <= val <= 2^31 - 1
 * Methods pop, top and getMin operations will always be called on non-empty
 * stacks.
 * At most 3 * 10^4 calls will be made to push, pop, top, and getMin.
 *
 *
*/

package main

import (
	"fmt"
)

func main() {
  testCases := []struct {
    fn        []string
    input     [][]int
    expected  []interface{}
  } {
    {[]string{"MinStack","push","push","push","getMin","pop","top","getMin"}, [][]int{{}, {-2}, {0}, {-3}, {}, {}, {}, {}}, []interface{}{nil, nil, nil, nil, -3, nil, 0, -2}},
  }

  for _, testCase := range testCases {
    var minStack MinStack
    for i, fn := range testCase.fn {
      switch fn {
        case "MinStack":
          minStack = Constructor()
          break
        case "push":
          minStack.Push(testCase.input[i][0])
          break
        case "pop":
          minStack.Pop()
          break
        case "top":
          result := minStack.Top()
          fmt.Printf("expected: %v, result: %v\n", testCase.expected[i], result)
          break
        case "getMin":
          result := minStack.GetMin()
          fmt.Printf("expected: %v, result: %v\n", testCase.expected[i], result)
          break
      }
    }
  }
}

// @lc code=start
type MinStack struct {
  stack   []int
  min     []int
  top     int
}


func Constructor() MinStack {
  return MinStack{
    stack: make([]int, 0),
    min: make([]int, 0),
    top: -1,
  }
}


func (this *MinStack) Push(val int)  {
  this.stack = append(this.stack, val)

  if this.top == -1 {
    this.min = append(this.min, val)
  } else {
    this.min = append(this.min, min(val, this.min[this.top]))
  }
  this.top++
}


func (this *MinStack) Pop()  {
  this.stack = this.stack[:this.top]
  this.min = this.min[:this.top]
  this.top--
  // fmt.Println("pop new stack =", this.stack, "min =", this.min, "top =", this.top)
}


func (this *MinStack) Top() int {
  // fmt.Println("top stack =", this.stack, "min =", this.min, "top =", this.top)
  return this.stack[this.top]    
}


func (this *MinStack) GetMin() int {
  // fmt.Println("getMin stack =", this.stack, "min =", this.min, "top =", this.top)
  return this.min[this.top]
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
// @lc code=end

