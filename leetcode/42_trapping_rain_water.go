package main

import "fmt"

func main() {
	height := []int{0,1,0,2,1,0,1,3,2,1,2,1}
	//height := []int{2,0,2}
	fmt.Println(trap(height)) // 6
}

// 指针
func trap(height []int) int {
	result := 0
	length := len(height)
	if length == 0 {
		return result
	}
	leftMax, rightMax := 0, 0
	for left, right := 0, length-1; left < right; {
		if height[left] < height[right] {
			if height[left] >= leftMax {
				leftMax = height[left]
			} else {
				result += leftMax - height[left]
			}
			left++
		} else {
			if height[right] >= rightMax {
				rightMax = height[right]
			} else {
				result += rightMax - height[right]
			}
			right--
		}
	}
	return result
}

//func trap(height []int) int {
//	result := 0
//	if len(height) == 0 {
//		return result
//	}
//	// 存储i的左右最大值中的最小值
//	left := make([]int, len(height))
//	left[0] = height[0]
//	for i := 1; i < len(height); i++ {
//		left[i] = maxInt(height[i], left[i-1])
//	}
//	right := make([]int ,len(height))
//	right[len(height)-1] = height[len(height)-1]
//	for i := len(height)-2; i > 0; i-- {
//		right[i] = maxInt(right[i+1], height[i])
//	}
//	for i := 0; i < len(height) ; i++ {
//		min := minInt(findHighest(i, 0, height,left), findHighest(i, 1, height, right))
//		if min > 0 {
//			result = result + min - height[i]
//		}
//	}
//	return result
//}

// 重复计算的内容为每次向左向右查找最大值，则关注下i的向左的最大值与i-1向左最大值的关系
// +cache，减少计算
func findHighest(index, path int, height, cache []int) int {
	highest := 0
	if path == 0 {
		if index == 0 {
			cache[index] = height[0]
		} else {
			cache[index] = maxInt(height[index], cache[index-1])
		}
		//highest = cache[index]
	} else {
		if index == len(height)-1 {
			cache[index] = height[len(height)-1]
		} else {
			cache[index] = maxInt(height[index], cache[index+1])
		}
		//highest = cache[index]
	}
	highest = cache[index]
	return highest
}

// 重复计算的内容为每次向左向右查找最大值，则关注下i的向左的最大值与i-1向左最大值的关系
//func cap(index int, height, cache []int) int {
//	highest := 0
//	if cache[index] > 0 {
//		highest = cache[index]
//		return highest
//	}
//	if height {
//
//	}
//		if cache[index-1] > 0 {
//			if cache[index-1] >= height[index] {
//				left := cache[index-1]
//			}
//		} else {
//			for i := index-1; i > 0; i-- {
//				if height[i] >= height[index] {
//					highest = maxInt(height[i], highest)
//				}
//			}
//
//		}
//
//		if cache[index+1] >= height[index] {
//			highest = cache[index+1]
//		}
//	}
//	return highest
//}

//func trap(height []int) int {
//	result := 0
//	for i := 0; i < len(height) ; i++ {
//		left := findHighest(i, 0, height)
//		right := findHighest(i, 1, height)
//		min := minInt(left, right)
//		if min > 0 {
//			result = result + min - height[i]
//		}
//	}
//	return result
//}
//
///*
//target 目标值
//path 0代表向左，1代表向右
//height 高度数组
//*/
//// 重复计算的内容为每次向左向右查找最大值，则关注下i的向左的最大值与i-1向左最大值的关系
//func findHighest(index, path int, height []int) int {
//	highest := 0
//	if path == 0 {
//		for i := index-1; i >= 0; i-- {
//			if height[i] >= height[index] {
//				highest = maxInt(height[i], highest)
//			}
//		}
//	} else {
//		for i := index+1; i < len(height); i++ {
//			if height[i] >= height[index] {
//				highest = maxInt(height[i], highest)
//			}
//		}
//	}
//	return highest
//}

func maxInt(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func minInt(x, y int) int {
	if x < y {
return x
}
return y
}