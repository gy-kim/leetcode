package main

import "fmt"

// https://leetcode.com/explore/challenge/card/july-leetcoding-challenge/544/week-1-july-1st-july-7th/3377/

// You have a total of n coins that you want to form in a staircase shape, where every k-th row must have exactly k coins.

// Given n, find the total number of full staircase rows that can be formed.

// n is a non-negative integer and fits within the range of a 32-bit signed integer.

// Example 1:

// n = 5

// The coins can form the following rows:
// ¤
// ¤ ¤
// ¤ ¤

// Because the 3rd row is incomplete, we return 2.
// Example 2:

// n = 8

// The coins can form the following rows:
// ¤
// ¤ ¤
// ¤ ¤ ¤
// ¤ ¤

// Because the 4th row is incomplete, we return 3.

func main() {

	n := 8

	result := solution(n)
	fmt.Println(result)

}

func solution(n int) int {
	left, right := 0, n
	for {

		// searching binary search
		// https://en.wikipedia.org/wiki/Binary_search_algorithm
		k := (left + right) / 2
		curr := k * (k + 1) / 2
		fmt.Printf("left: %d, right: %d, k: %d, curr: %d\n", left, right, k, curr)
		if curr == n {
			return k
		}

		if n < curr {
			right = k - 1
		} else {
			left = k + 1
		}

		if left > right {
			break
		}
	}

	return right
}
