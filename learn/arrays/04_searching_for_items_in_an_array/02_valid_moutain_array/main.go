package main

import "fmt"

// https://leetcode.com/explore/learn/card/fun-with-arrays/527/searching-for-items-in-an-array/3251/

// Given an array A of integers, return true if and only if it is a valid mountain array.

// Recall that A is a mountain array if and only if:

// A.length >= 3
// There exists some i with 0 < i < A.length - 1 such that:
// A[0] < A[1] < ... A[i-1] < A[i]
// A[i] > A[i+1] > ... > A[A.length - 1]

// Example 1:

// Input: [2,1]
// Output: false
// Example 2:

// Input: [3,5,5]
// Output: false
// Example 3:

// Input: [0,3,2,1]
// Output: true

// Note:

// 0 <= A.length <= 10000
// 0 <= A[i] <= 10000

func main() {
	A := []int{0, 3, 2, 1}
	// A := []int{0, 3, 4, 3, 2, 3, 1}

	result := solution(A)
	fmt.Println(result)
}

func solution(A []int) bool {

	if len(A) <= 2 {
		return false
	}

	if A[0] >= A[1] || A[len(A)-2] <= A[len(A)-1] {
		return false
	}

	up := true

	for i := 1; i < len(A)-1; i++ {
		if A[i] == A[i+1] {
			return false
		}

		if A[i] > A[i+1] {
			if up {
				up = false
			}
		} else {
			if !up {
				return false
			}
		}
	}

	if up {
		return false
	}
	return true
}
