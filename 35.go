package main

import "fmt"

//二分查找
func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1
	ans := len(nums)
	for left <= right {
		mid := (right-left)/2 + left
		if target <= nums[mid] {
			right = mid - 1
			ans = mid
		} else {
			left = mid + 1
		}
	}
	return ans
}

func main() {
	nums := []int{1, 2, 5, 9}
	target := 4
	fmt.Println(searchInsert(nums, target))
}
