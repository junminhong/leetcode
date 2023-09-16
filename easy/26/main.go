package main

func main() {
	nums := []int{1, 2}
	removeDuplicates(nums)
}

func removeDuplicates(nums []int) int {
	// create new pointer
	// default index 1
	pointer := 1
	for index := 1; index < len(nums); index++ {
		if nums[index] != nums[index-1] {
			// 如有有不一樣的話就取代位置
			nums[pointer] = nums[index]
			pointer++
		}
	}
	return pointer
}
