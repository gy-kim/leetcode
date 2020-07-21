package main

import "fmt"

// https://leetcode.com/explore/learn/card/fun-with-arrays/525/inserting-items-into-an-array/3253/

// Given two sorted integer arrays nums1 and nums2, merge nums2 into nums1 as one sorted array.

// Note:

// The number of elements initialized in nums1 and nums2 are m and n respectively.
// You may assume that nums1 has enough space (size that is equal to m + n) to hold additional elements from nums2.
// Example:

// Input:
// nums1 = [1,2,3,0,0,0], m = 3
// nums2 = [2,5,6],       n = 3

// Output: [1,2,2,3,5,6]

// Constraints:

// -10^9 <= nums1[i], nums2[i] <= 10^9
// nums1.length == m + n
// nums2.length == n

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	m := 3
	nums2 := []int{2, 5, 6}
	n := 3

	solution(nums1, m, nums2, n)
}

func solution(nums1 []int, m int, nums2 []int, n int) {

	n1 := 0
	n2 := 0
	idx := 0
	newArr := make([]int, m+n, m+n)

	for n1 <= m-1 && n2 <= n-1 {
		if nums1[n1] < nums2[n2] {
			newArr[idx] = nums1[n1]
			n1++
		} else if nums1[n1] > nums2[n2] {
			newArr[idx] = nums2[n2]
			n2++
		} else {
			newArr[idx] = nums1[n1]
			idx++
			newArr[idx] = nums2[n2]
			n1++
			n2++
		}
		idx++
	}

	for n1 <= m-1 {
		newArr[idx] = nums1[n1]
		n1++
		idx++
	}

	for n2 <= n-1 {
		newArr[idx] = nums2[n2]
		n2++
		idx++
	}

	fmt.Println(newArr)
	copy(nums1, newArr)
}
