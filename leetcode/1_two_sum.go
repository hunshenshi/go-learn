package main

import "fmt"

func main() {
	nums := []int{2, 7, 11, 15}
	fmt.Println(twoSum(nums, 9))
}

func twoSum(nums []int, target int) []int {
	//var result []int
	result := []int{0,0}
	hashMap := make(map[int]int)

	// 快速查找使用hashmap O(1)
	for i := 0; i < len(nums); i++ {
		if _, ok := hashMap[target - nums[i]]; ok {
			result[0] = i
			result[1] = hashMap[target - nums[i]]
		} else {
			hashMap[nums[i]] = i
		}
	}

	//for i := 0; i < len(nums); i++ {
	//	hashMap[nums[i]] = i
	//}
	//
	//for i := 0; i< len(nums); i++ {
	//	//_, ok := hashMap[20]; fmt.Println(ok)
	//	if _, ok := hashMap[target - nums[i]]; ok {
	//		result[0] = i
	//		result[1] = hashMap[target - nums[i]]
	//		break
	//	}
	//}
	return  result
}
