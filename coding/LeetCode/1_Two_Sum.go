package main

/*

https://leetcode.com/problems/two-sum/

Easy, Hash Table

*/

func twoSum(nums []int, target int) []int {
	valueMap := make(map[int]int)
	for i, num := range nums {
		if v, has := valueMap[target-num]; has && i != v {
			return []int{i, v}
		}
		valueMap[num] = i
	}
	return nil
}
