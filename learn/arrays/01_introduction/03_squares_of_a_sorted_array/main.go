package main

import (
	"fmt"
	"sort"
)

// https://leetcode.com/explore/learn/card/fun-with-arrays/521/introduction/3240/

// Given an array of integers A sorted in non-decreasing order, return an array of the squares of each number, also in sorted non-decreasing order.

// Example 1:

// Input: [-4,-1,0,3,10]
// Output: [0,1,9,16,100]
// Example 2:

// Input: [-7,-3,2,3,11]
// Output: [4,9,9,49,121]

// Note:

// 1 <= A.length <= 10000
// -10000 <= A[i] <= 10000
// A is sorted in non-decreasing order.

func main() {
	// input := []int{-4, -1, 0, 3, 10}
	input := []int{-7, -3, 2, 3, 11}

	// result := solution1(input)
	result := solution2(input)
	fmt.Println(result)
}

func solution1(A []int) []int {
	arr := make([]int, len(A))

	for i := 0; i < len(A); i++ {
		arr[i] = A[i] * A[i]
	}

	sort.Ints(arr)

	return arr

}

func solution2(A []int) []int {
	n := len(A)
	j := 0

	for j < n && A[j] < 0 {
		j++
	}
	i := j - 1

	ans := make([]int, len(A))
	t := 0

	for i >= 0 && j < n {
		if A[i]*A[i] < A[j]*A[j] {
			ans[t] = A[i] * A[i]
			i--
		} else {
			ans[t] = A[j] * A[j]
			j++
		}
		t++
	}

	for i >= 0 {
		ans[t] = A[i] * A[i]
		t++
		i--
	}

	for j < n {
		ans[t] = A[j] * A[j]
		t++
		j++
	}

	return ans
}
