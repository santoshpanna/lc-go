/*
 * @lc app=leetcode id=1429 lang=golang
 *
 * [1429] First Unique Number
 *
 * https://leetcode.com/problems/first-unique-number/description/
 *
 * algorithms
 * Medium (54.04%)
 * Likes:    566
 * Dislikes: 32
 * Total Accepted:    89.3K
 * Total Submissions: 164.3K
 * Testcase Example:  '["FirstUnique","showFirstUnique","add","showFirstUnique","add","showFirstUnique","add","showFirstUnique"]\r\n' +
  '[[[2,3,5]],[],[5],[],[2],[],[3],[]]\r'
 *
 * You have a queue of integers, you need to retrieve the first unique integer
 * in the queue.
 *
 * Implement the FirstUnique class:
 *
 *
 * FirstUnique(int[] nums) Initializes the object with the numbers in the
 * queue.
 * int showFirstUnique() returns the value of the first unique integer of the
 * queue, and returns -1 if there is no such integer.
 * void add(int value) insert value to the queue.
 *
 *
 *
 * Example 1:
 *
 *
 * Input:
 *
 * ["FirstUnique","showFirstUnique","add","showFirstUnique","add","showFirstUnique","add","showFirstUnique"]
 * [[[2,3,5]],[],[5],[],[2],[],[3],[]]
 * Output:
 * [null,2,null,2,null,3,null,-1]
 * Explanation:
 * FirstUnique firstUnique = new FirstUnique([2,3,5]);
 * firstUnique.showFirstUnique(); // return 2
 * firstUnique.add(5);            // the queue is now [2,3,5,5]
 * firstUnique.showFirstUnique(); // return 2
 * firstUnique.add(2);            // the queue is now [2,3,5,5,2]
 * firstUnique.showFirstUnique(); // return 3
 * firstUnique.add(3);            // the queue is now [2,3,5,5,2,3]
 * firstUnique.showFirstUnique(); // return -1
 *
 *
 * Example 2:
 *
 *
 * Input:
 *
 * ["FirstUnique","showFirstUnique","add","add","add","add","add","showFirstUnique"]
 * [[[7,7,7,7,7,7]],[],[7],[3],[3],[7],[17],[]]
 * Output:
 * [null,-1,null,null,null,null,null,17]
 * Explanation:
 * FirstUnique firstUnique = new FirstUnique([7,7,7,7,7,7]);
 * firstUnique.showFirstUnique(); // return -1
 * firstUnique.add(7);            // the queue is now [7,7,7,7,7,7,7]
 * firstUnique.add(3);            // the queue is now [7,7,7,7,7,7,7,3]
 * firstUnique.add(3);            // the queue is now [7,7,7,7,7,7,7,3,3]
 * firstUnique.add(7);            // the queue is now [7,7,7,7,7,7,7,3,3,7]
 * firstUnique.add(17);           // the queue is now [7,7,7,7,7,7,7,3,3,7,17]
 * firstUnique.showFirstUnique(); // return 17
 *
 *
 * Example 3:
 *
 *
 * Input:
 * ["FirstUnique","showFirstUnique","add","showFirstUnique"]
 * [[[809]],[],[809],[]]
 * Output:
 * [null,809,null,-1]
 * Explanation:
 * FirstUnique firstUnique = new FirstUnique([809]);
 * firstUnique.showFirstUnique(); // return 809
 * firstUnique.add(809);          // the queue is now [809,809]
 * firstUnique.showFirstUnique(); // return -1
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= nums.length <= 10^5
 * 1 <= nums[i] <= 10^8
 * 1 <= value <= 10^8
 * At most 50000 calls will be made to showFirstUnique and add.
 *
 *
*/

package main

import (
	"fmt"
)

func main() {
  testCases := []struct{
    fnx []string
    args [][]int
    expected []int
  } {
    {[]string{"FirstUnique", "showFirstUnique", "add", "showFirstUnique", "add", "showFirstUnique", "add", "showFirstUnique"}, [][]int{{2,3,5}, {}, {5}, {}, {2}, {}, {3}, {}}, []int{0, 2, 0, 2, 0, 3, 0, -1}},
    {[]string{"FirstUnique", "showFirstUnique", "add", "add", "add", "add", "add", "showFirstUnique"}, [][]int{{7,7,7,7,7,7}, {}, {7}, {3}, {3}, {7}, {17}, {}}, []int{0, -1, 0, 0, 0, 0, 0, 17}},
    {[]string{"FirstUnique", "showFirstUnique", "add", "showFirstUnique"}, [][]int{{809}, {}, {809}, {}}, []int{0, 809, 0, -1}},
  }

  for _, testCase := range testCases {
    obj := Constructor(testCase.args[0])
    for i := 1; i < len(testCase.fnx); i++ {
      switch testCase.fnx[i] {
      case "showFirstUnique":
        result := obj.ShowFirstUnique()
        fmt.Printf("expected: %v, got: %v\n", testCase.expected[i], result)
      case "add":
        obj.Add(testCase.args[i][0])
      }
    }
  }
}

// @lc code=start
type FirstUnique struct {
  freqMap map[int]int
  unique []int
}


func Constructor(nums []int) FirstUnique {
  freqMap := make(map[int]int)
  unique := make([]int, 0)

  for _, num := range nums {
    freqMap[num]++
  }

  for _, num := range nums {
    if freqMap[num] == 1 {
      unique = append(unique, num)
    }
  }

  return FirstUnique{
    freqMap: freqMap,
    unique: unique,
  }
}


func (this *FirstUnique) ShowFirstUnique() int {
  if len(this.unique) == 0 {
    return -1
  }

  return this.unique[0]
}


func (this *FirstUnique) Add(value int)  {
  this.freqMap[value]++
  if this.freqMap[value] == 1 {
    this.unique = append(this.unique, value)
  } else {
    for i := 0; i < len(this.unique); i++ {
      if this.unique[i] == value {
        this.unique = append(this.unique[:i], this.unique[i+1:]...)
        break
      }
    }
  }
}


/**
 * Your FirstUnique object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.ShowFirstUnique();
 * obj.Add(value);
 */
// @lc code=end

