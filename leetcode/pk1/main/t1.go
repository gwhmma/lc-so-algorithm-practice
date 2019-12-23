package main

import "fmt"

/*
两数之和
给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。
 */
func twoSum1(nums []int, target int) []int {
	for k1, v1 := range nums {
		for i := k1 + 1; i < len(nums); i++ {
			if v1+nums[i] == target {
				return []int{k1, i}
			}
		}
	}
	return nil
}

func twoSum2(nums []int, target int) []int {
	var m = make(map[int]int)
	for k, v := range nums {
		m[v] = k
	}

	for i := 0; i < len(nums); i++ {
		c := target - nums[i]
		if k, ok := m[c]; ok && k != i {
			return []int{i, k}
		}
	}
	return nil
}

func main() {
	nums := []int{2, 7, 11, 15}
	fmt.Println(twoSum1(nums, 18))
	fmt.Println(twoSum2(nums, 18))
}
