package main

import "fmt"

// https://leetcode.com/explore/learn/card/fun-with-arrays/511/in-place-operations/3260/

// Given an array A of non-negative integers, return an array consisting of all the even elements of A, followed by all the odd elements of A.

// You may return any answer array that satisfies this condition.

// Example 1:

// Input: [3,1,2,4]
// Output: [2,4,3,1]
// The outputs [4,2,3,1], [2,4,1,3], and [4,2,1,3] would also be accepted.

func main() {
	A := []int{3, 1, 2, 4}

	result := solution(A)
	fmt.Println(result)
}

func solution(A []int) []int {
	if len(A) == 0 || len(A) == 1 {
		return A
	}

	evenIdx := 0
	for i := 0; i < len(A); i++ {
		if A[i]%2 == 0 {
			tmp := A[evenIdx]
			A[evenIdx] = A[i]
			A[i] = tmp
			evenIdx++
		}
	}
	return A
}
