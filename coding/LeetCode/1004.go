package main

/*

https://leetcode.com/problems/max-consecutive-ones-iii/

Medium, Sliding Window

*/

func longestOnes(nums []int, k int) int {
	var left int
	flips := k
	var ones, ans int
	for right := 0; right < len(nums); right++ {
		if nums[right] == 0 {
			if flips > 0 {
				ones++
				flips--
			} else {
				if ones > ans {
					ans = ones
				}

				// get one flip back to continue
				for nums[left] == 1 {
					ones--
					left++
				}
				// now nums[left] == 0, move over it, and gets 1 flip back
				// but still consumes by nums[right],
				// ! ones unchanged, lost one, get one back
				left++
			}
		} else {
			ones++
		}
	}

	// ! [0, 0, 0, 1], k = 4
	if ones > ans {
		ans = ones
	}

	return ans
}
