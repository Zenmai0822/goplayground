package goplayground

import (
	"sort"
)

// ThreeSum is a function that returns a slice of integer slices of 3 nums that adds to 0.
// This is a direct translation of my python code so that I can get my hands dirty.
// Lec 0015
func ThreeSum(nums []int) [][]int {
	var res [][]int
	sort.Ints(nums)
	for i, n := range nums {
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}
		li := i + 1
		ri := len(nums) - 1
		for li < ri {
			sum := n + nums[li] + nums[ri]
			if sum == 0 {
				res = append(res, []int{n, nums[li], nums[ri]})
				li++
				ri--
				// we have incremented already, thus we are comparing li with li-1
				for li < ri && nums[li] == nums[li-1] {
					li++
				}
				for li < ri && nums[ri] == nums[ri+1] {
					ri--
				}
			} else if sum > 0 {
				ri--
			} else {
				li++
			}
		}
	}
	return res
}
