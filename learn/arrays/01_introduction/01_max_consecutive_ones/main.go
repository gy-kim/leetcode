package main

import "fmt"

// https://leetcode.com/explore/learn/card/fun-with-arrays/521/introduction/3238/

// Given a binary array, find the maximum number of consecutive 1s in this array.

// Example 1:
// Input: [1,1,0,1,1,1]
// Output: 3
// Explanation: The first two digits or the last three digits are consecutive 1s.
//     The maximum number of consecutive 1s is 3.
// Note:

// The input array will only contain 0 and 1.
// The length of input array is a positive integer and will not exceed 10,000

func main() {

	input := []int{1, 1, 0, 1, 1, 1}

	result := solution(input)
	fmt.Println(result)

}

func solution(nums []int) int {
	max := 0

	if len(nums) == 0 {
		return max
	}

	cnt := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			cnt++
		} else {
			if cnt > max {
				max = cnt
			}
			cnt = 0
		}
	}

	if cnt > max {
		max = cnt
	}

	return max
}
