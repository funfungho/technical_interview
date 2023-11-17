package main

/*

https://leetcode.com/problems/subarray-product-less-than-k/

Medium, Two Pointer

*/

func numSubarrayProductLessThanK(nums []int, k int) int {
	// nums[i] >= 1
	// ! handle special cases
	if k <= 1 {
		// otherwise prod will always be >= k, infinite loop,
		// array will go out of bounds
		return 0
	}

	var left int
	var ans int
	prod := 1
	for right := 0; right <= len(nums)-1; right++ {
		prod *= nums[right]

		for prod >= k {
			prod /= nums[left]
			left++
		}

		// valid boundary: [left, right]
		// subarrays that ends at index right
		ans += right - left + 1
	}

	return ans
}
