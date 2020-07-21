package main

// https://leetcode.com/explore/learn/card/fun-with-arrays/525/inserting-items-into-an-array/3245/

// Given a fixed length array arr of integers, duplicate each occurrence of zero, shifting the remaining elements to the right.

// Note that elements beyond the length of the original array are not written.

// Do the above modifications to the input array in place, do not return anything from your function.

// Example 1:

// Input: [1,0,2,3,0,4,5,0]
// Output: null
// Explanation: After calling your function, the input array is modified to: [1,0,0,2,3,0,0,4]
// Example 2:

// Input: [1,2,3]
// Output: null
// Explanation: After calling your function, the input array is modified to: [1,2,3]

// Note:

// 1 <= arr.length <= 10000
// 0 <= arr[i] <= 9

func main() {

	arr := []int{1, 0, 2, 3, 0, 4, 5, 0}
	solution(arr)
}

func solution(arr []int) {
	arrLen := len(arr)

	if arrLen == 0 {
		return
	}

	newIdx := 0
	result := make([]int, arrLen, arrLen)

	for i := 0; i < arrLen; i++ {
		if newIdx > arrLen-1 {
			break
		}
		if arr[i] != 0 {
			result[newIdx] = arr[i]
		} else {
			newIdx++
		}
		newIdx++
	}

	copy(arr, result)
}
