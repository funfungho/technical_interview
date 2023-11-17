package main

/*

https://leetcode.com/problems/maximum-average-subarray-i/

Easy, Sliding window

*/

func findMaxAverage(nums []int, k int) float64 {
	var sum int
	var left int
	for i := 0; i < k; i++ {
		sum += nums[i]
	}

	tmpSum := sum
	for right := k; right <= len(nums)-1; right++ {
		tmpSum = tmpSum + nums[right] - nums[left]
		if tmpSum > sum {
			sum = tmpSum
		}
		left++
	}
	return float64(sum) / float64(k)
}
