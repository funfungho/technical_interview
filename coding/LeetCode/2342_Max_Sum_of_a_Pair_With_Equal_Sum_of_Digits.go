package main

/*

https://leetcode.com/problems/max-sum-of-a-pair-with-equal-sum-of-digits/

Medium, Hash Table

*/

func maximumSum(nums []int) int {
	max := -1
	// key: sum of digits
	// value: number
	ht := make(map[int]int)
	for _, num := range nums {
		sum := digitsSum(num)
		if v, has := ht[sum]; has {
			if max < v+num {
				max = v + num
			}
		}
		if ht[sum] < num {
			ht[sum] = num
		}
	}
	return max
}

func digitsSum(num int) int {
	var ans int
	for num > 0 {
		ans += num % 10
		num /= 10
	}
	return ans
}
