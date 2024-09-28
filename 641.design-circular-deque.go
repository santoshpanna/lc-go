/*
 * @lc app=leetcode id=641 lang=golang
 *
 * [641] Design Circular Deque
 *
 * https://leetcode.com/problems/design-circular-deque/description/
 *
 * algorithms
 * Medium (57.31%)
 * Likes:    1365
 * Dislikes: 89
 * Total Accepted:    106.4K
 * Total Submissions: 174.8K
 * Testcase Example:  '["MyCircularDeque","insertLast","insertLast","insertFront","insertFront","getRear","isFull","deleteLast","insertFront","getFront"]\n' +
  '[[3],[1],[2],[3],[4],[],[],[],[4],[]]'
 *
 * Design your implementation of the circular double-ended queue (deque).
 *
 * Implement the MyCircularDeque class:
 *
 *
 * MyCircularDeque(int k) Initializes the deque with a maximum size of k.
 * boolean insertFront() Adds an item at the front of Deque. Returns true if
 * the operation is successful, or false otherwise.
 * boolean insertLast() Adds an item at the rear of Deque. Returns true if the
 * operation is successful, or false otherwise.
 * boolean deleteFront() Deletes an item from the front of Deque. Returns true
 * if the operation is successful, or false otherwise.
 * boolean deleteLast() Deletes an item from the rear of Deque. Returns true if
 * the operation is successful, or false otherwise.
 * int getFront() Returns the front item from the Deque. Returns -1 if the
 * deque is empty.
 * int getRear() Returns the last item from Deque. Returns -1 if the deque is
 * empty.
 * boolean isEmpty() Returns true if the deque is empty, or false
 * otherwise.
 * boolean isFull() Returns true if the deque is full, or false otherwise.
 *
 *
 *
 * Example 1:
 *
 *
 * Input
 * ["MyCircularDeque", "insertLast", "insertLast", "insertFront",
 * "insertFront", "getRear", "isFull", "deleteLast", "insertFront", "getFront"]
 * [[3], [1], [2], [3], [4], [], [], [], [4], []]
 * Output
 * [null, true, true, true, false, 2, true, true, true, 4]
 *
 * Explanation
 * MyCircularDeque myCircularDeque = new MyCircularDeque(3);
 * myCircularDeque.insertLast(1);  // return True
 * myCircularDeque.insertLast(2);  // return True
 * myCircularDeque.insertFront(3); // return True
 * myCircularDeque.insertFront(4); // return False, the queue is full.
 * myCircularDeque.getRear();      // return 2
 * myCircularDeque.isFull();       // return True
 * myCircularDeque.deleteLast();   // return True
 * myCircularDeque.insertFront(4); // return True
 * myCircularDeque.getFront();     // return 4
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= k <= 1000
 * 0 <= value <= 1000
 * At most 2000 calls will be made to insertFront, insertLast, deleteFront,
 * deleteLast, getFront, getRear, isEmpty, isFull.
 *
 *
*/

package main

import "fmt"

func main() {
  testCases := []struct {
    fn []string
    input []int
    output []interface{}
  }{
    // {
    //   fn: []string{"MyCircularDeque", "insertLast", "insertLast", "insertFront", "insertFront", "getRear", "isFull", "deleteLast", "insertFront", "getFront"},
    //   input: []int{3, 1, 2, 3, 4, 0, 0, 0, 4, 0},
    //   output: []interface{}{nil,true,true,true,false,2,true,true,true,4},
    // },
    // {
    //   fn: []string{"MyCircularDeque","insertFront","insertFront","insertFront","deleteLast","deleteLast","deleteLast"},
    //   input: []int{3, 1, 2, 3, 0, 0, 0},
    //   output: []interface{}{nil,true,true,true,true,true,true},
    // },
    // {
    //   fn: []string{"MyCircularDeque","insertLast","insertLast","insertLast","deleteLast","deleteLast","deleteLast", "deleteLast"},
    //   input: []int{3, 1, 2, 3, 0, 0, 0, 0},
    //   output: []interface{}{nil,true,true,true,true,true,true,false},
    // },
    // {
    //   fn: []string{"MyCircularDeque","insertFront","insertFront","insertFront","deleteLast","deleteLast","deleteLast", "deleteLast"},
    //   input: []int{3, 1, 2, 3, 0, 0, 0, 0},
    //   output: []interface{}{nil,true,true,true,true,true,true,false},
    // },
    // {
    //   fn: []string{"MyCircularDeque","insertFront","insertLast","getRear","getFront","getFront","deleteLast","deleteFront","getRear"},
    //   input: []int{41, 70, 11, 0, 0, 0, 0, 0, 0},
    //   output: []interface{}{nil, true,true,11,70,70,true,true,-1},
    // },
    // {
    //   fn: []string{"MyCircularDeque","insertFront","insertFront","getRear","getFront","getFront","deleteLast","deleteFront","getRear"},
    //   input: []int{41, 70, 11, 0, 0, 0, 0, 0, 0},
    //   output: []interface{}{nil,true,true,70,11,11,true,true,-1},
    // },
    // {
    //   fn: []string{"MyCircularDeque","insertFront","insertLast","deleteFront","getFront","deleteLast","insertLast","isEmpty","deleteLast","insertFront","getRear","deleteFront","insertFront","insertLast","deleteLast","getFront","getRear","insertFront","getRear","getFront"},
    //   input: []int{999, 93, 578, 0, 0, 0, 533, 0, 0, 913, 0, 0, 100, 57, 0, 0, 0, 900, 0, 0},
    //   output: []interface{}{nil,true,true,true,578,true,true,false,true,true,913,true,true,true,true,100,100,true,100,900},
    // },
    {
      fn: []string{"MyCircularDeque","insertFront","deleteLast","getRear","getFront","getFront","deleteFront","insertFront","insertLast","insertFront","getFront","insertFront"},
      input: []int{4,9,0,0,0,0,0,6,5,9,0,6},
      output: []interface{}{nil,true,true,-1,-1,-1,false,true,true,true,9,true},
    },
    // {
    //   fn: []string{"MyCircularDeque","insertFront","getRear","deleteLast","getRear","insertFront","insertFront","insertFront","insertFront","isFull","insertFront","isFull","getRear","deleteLast","getFront","getFront","insertLast","deleteFront","getFront","insertLast","getRear","insertLast","getRear","getFront","getFront","getFront","getRear","getRear","insertFront","getFront","getFront","getFront","getFront","deleteFront","insertFront","getFront","deleteLast","insertLast","insertLast","getRear","getRear","getRear","isEmpty","insertFront","deleteLast","getFront","deleteLast","getRear","getFront","isFull","isFull","deleteFront","getFront","deleteLast","getRear","insertFront","getFront","insertFront","insertFront","getRear","isFull","getFront","getFront","insertFront","insertLast","getRear","getRear","deleteLast","insertFront","getRear","insertLast","getFront","getFront","getFront","getRear","insertFront","isEmpty","getFront","getFront","insertFront","deleteFront","insertFront","deleteLast","getFront","getRear","getFront","insertFront","getFront","deleteFront","insertFront","isEmpty","getRear","getRear","getRear","getRear","deleteFront","getRear","isEmpty","deleteFront","insertFront","insertLast","deleteLast"},
    //   input: []int{77,89,0,0,0,19,23,23,82,0,45,0,0,0,0,0,74,0,0,98,0,99,0,0,0,0,0,0,8,0,0,0,0,0,75,0,0,35,59,0,0,0,0,22,0,0,0,0,0,0,0,0,0,0,0,21,0,26,63,0,0,0,0,87,76,0,0,0,26,0,67,0,0,0,0,36,0,0,0,72,0,87,0,0,0,0,85,0,0,91,0,0,0,0,0,0,0,0,0,34,44,0},
    //   output: []interface{}{nil,true,89,true,-1,true,true,true,true,false,true,false,19,true,45,45,true,true,82,true,98,true,99,82,82,82,99,99,true,8,8,8,8,true,true,75,true,true,true,59,59,59,false,true,true,22,true,98,22,false,false,true,75,true,74,true,21,true,true,74,false,63,63,true,true,76,76,true,true,74,true,26,26,26,67,true,false,36,36,true,true,true,true,87,74,87,true,85,true,true,false,74,74,74,74,true,74,false,true,true,true,true},
    // },
  }

  for _, testCase := range testCases {
    var res []interface{}
    var obj MyCircularDeque
    for i := 0; i < len(testCase.fn); i++ {
      switch testCase.fn[i] {
      case "MyCircularDeque":
        obj = Constructor(testCase.input[i])
        res = append(res, nil)
      case "insertFront":
        res = append(res, obj.InsertFront(testCase.input[i]))
      case "insertLast":
        res = append(res, obj.InsertLast(testCase.input[i]))
      case "deleteFront":
        res = append(res, obj.DeleteFront())
      case "deleteLast":
        res = append(res, obj.DeleteLast())
      case "getFront":
        res = append(res, obj.GetFront())
      case "getRear":
        res = append(res, obj.GetRear())
      case "isEmpty":
        res = append(res, obj.IsEmpty())
      case "isFull":
        res = append(res, obj.IsFull())
      }
    }
    fmt.Println("Test case failed:", testCase.input, res, testCase.output)
  }
}

// @lc code=start
type MyCircularDeque struct {
  queue     []int
  max       int
  head      int
}


func Constructor(k int) MyCircularDeque {
  return MyCircularDeque{max : k, queue: make([]int, k), head: 0}
}


func (this *MyCircularDeque) InsertFront(value int) bool {
  if this.head == this.max {
    return false
  }

  this.queue = append([]int{value}, this.queue...)
  this.head = this.head + 1

  return true
}


func (this *MyCircularDeque) InsertLast(value int) bool {
  if this.head == this.max {
    return false
  }

  this.queue[this.head] = value
  this.head = this.head + 1

  return true
}


func (this *MyCircularDeque) DeleteFront() bool {
  if this.head > 0 {
    this.queue = append(this.queue[1:])
    this.head = this.head - 1
    
    return true
  }

  return false
}


func (this *MyCircularDeque) DeleteLast() bool {
  if this.head > 0 {
    this.head = this.head - 1

    return true
  }

  return false
}


func (this *MyCircularDeque) GetFront() int {
  if this.head > 0 {
    return this.queue[0]
  }    

  return -1
}


func (this *MyCircularDeque) GetRear() int {
  if this.head > 0 {
    return this.queue[this.head - 1]
  }

  return -1
}


func (this *MyCircularDeque) IsEmpty() bool {
  if this.head == 0 {
    return true
  }

  return false
}


func (this *MyCircularDeque) IsFull() bool {
  if this.head == this.max {
    return true
  }

  return false
}


/**
 * Your MyCircularDeque object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.InsertFront(value);
 * param_2 := obj.InsertLast(value);
 * param_3 := obj.DeleteFront();
 * param_4 := obj.DeleteLast();
 * param_5 := obj.GetFront();
 * param_6 := obj.GetRear();
 * param_7 := obj.IsEmpty();
 * param_8 := obj.IsFull();
 */
// @lc code=end

