package main

import "fmt"

// https://leetcode.com/explore/learn/card/fun-with-arrays/511/in-place-operations/3259/

// Given an array arr, replace every element in that array with the greatest element among the elements to its right, and replace the last element with -1.

// After doing so, return the array.

// Example 1:

// Input: arr = [17,18,5,4,6,1]
// Output: [18,6,6,6,1,-1]

// Constraints:

// 1 <= arr.length <= 10^4
// 1 <= arr[i] <= 10^5

func main() {
	arr := []int{17, 18, 5, 4, 6, 1}
	result := solution(arr)
	fmt.Println(result)
}

func solution(arr []int) []int {
	if len(arr) == 0 {
		return []int{}
	}
	if len(arr) == 1 {
		return []int{-1}
	}

	largest := arr[len(arr)-1]
	result := make([]int, len(arr), len(arr))

	for i := len(arr) - 2; i >= 0; i-- {
		if arr[i+1] > largest {
			largest = arr[i+1]
		}
		result[i] = largest
	}

	result[len(result)-1] = -1

	return result
}
