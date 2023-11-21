package main

/*

https://leetcode.com/problems/minimum-consecutive-cards-to-pick-up/

Medium, Hash Table

*/

func minimumCardPickup(cards []int) int {
	ans := 100001
	// key: card number
	// value: latest index of some card number
	ht := make(map[int]int)
	for i, v := range cards {
		if vv, has := ht[v]; has {
			if ans > i-vv+1 {
				ans = i - vv + 1
			}
		}
		ht[v] = i
	}

	if ans == 100001 {
		ans = -1
	}

	return ans
}
