package main

func main() {
	nums := []int{1, 1, 2}
	removeDuplicates(nums)
}

func removeDuplicates(nums []int) int {
	length := 0
	for index := 0; index < len(nums); index++ {
		if index == 0 || nums[index] != nums[index-1] {
			nums[length] = nums[index]
			length++
		}
	}
	return length
}
