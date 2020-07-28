package main

import "fmt"

// https://leetcode.com/explore/learn/card/fun-with-arrays/527/searching-for-items-in-an-array/3250/

// Given an array arr of integers, check if there exists two integers N and M such that N is the double of M ( i.e. N = 2 * M).

// More formally check if there exists two indices i and j such that :

// i != j
// 0 <= i, j < arr.length
// arr[i] == 2 * arr[j]

// Example 1:

// Input: arr = [10,2,5,3]
// Output: true
// Explanation: N = 10 is the double of M = 5,that is, 10 = 2 * 5.
// Example 2:

// Input: arr = [7,1,14,11]
// Output: true
// Explanation: N = 14 is the double of M = 7,that is, 14 = 2 * 7.
// Example 3:

// Input: arr = [3,1,7,11]
// Output: false
// Explanation: In this case does not exist N and M, such that N = 2 * M.

func main() {
	// arr := []int{10, 2, 5, 3}
	arr := []int{-2, 0, 0, 10, 19, 4, 6, -8}
	result := solution(arr)
	fmt.Println(result)
}

func solution(arr []int) bool {
	if len(arr) < 2 {
		return false
	}

	m := make(map[int]int)

	for _, a := range arr {
		if _, ok := m[a]; ok {
			cnt := m[a]
			m[a] = cnt + 1
		} else {
			m[a] = 1
		}

	}

	// fmt.Println(m)

	for _, a := range arr {
		if a == 0 {
			if m[a] > 1 {
				return true
			}
			continue
		}
		if _, ok := m[a*2]; ok {
			fmt.Println(a)
			return true
		}
	}

	return false
}
