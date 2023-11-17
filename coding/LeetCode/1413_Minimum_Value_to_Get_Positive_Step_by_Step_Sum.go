package main

/*

https://leetcode.com/problems/minimum-value-to-get-positive-step-by-step-sum/

Easy, Prefix Sum

*/

func minStartValue(nums []int) int {
	sum := nums[0]
	minimun := sum

	for i := 1; i < len(nums); i++ {
		sum += +nums[i]
		if minimun > sum {
			minimun = sum
		}
	}

	if minimun >= 0 {
		return 1
	} else {
		return -minimun + 1
	}

}
