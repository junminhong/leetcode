package main

import "sort"

func main() {
	nums1 := []int{1, 3}
	nums2 := []int{2}
	findMedianSortedArrays(nums1, nums2)
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums := append(nums1, nums2...)
	sort.Ints(nums)
	middle := len(nums) / 2
	if flag := len(nums) % 2; flag != 0 {
		return float64(nums[middle])
	} else {
		return (float64(nums[middle-1]) + float64(nums[middle])) / 2
	}
}
