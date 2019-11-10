package main

import (
	"fmt"
	"math"
)

func main() {
	nums1 := []int{}
	nums2 := []int{3}
	//k := 5
	//fmt.Printf("第%d小的数是：%d\n", k, findK(nums1, nums2, k))
	fmt.Println(findMedianSortedArrays(nums1, nums2))
	fmt.Println(findMedianSorted(nums1, nums2))
}

// 	nums1 := []int{}
//	nums2 := []int{3}
// case失败，长度为1，num为0
func findMedianSorted(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		tmp := nums1
		nums1 = nums2
		nums2 = tmp
	}
	m, n := len(nums1), len(nums2)
	num := (m + n)/2
	// 0-m为m+1种切分的方法，二分查找在哪里切分，
	for low, high := 0, m; low <= high; {
		i := (low + high)/2
		j := num - i
		if j > 0 && i < m && nums2[j-1] > nums1[i] {
			low = i + 1
		} else if i > 0 && j < n && nums1[i-1] > nums2[j] {
			high = i - 1
		} else {
			maxLeft := 0.0
			if i == 0 {
				maxLeft = func(x int) float64 {
					if x == 0 {
						return float64(nums2[x])
					} else {
						return float64(nums2[x - 1])
					}
				}(j)
			} else if j == 0 {
				maxLeft = float64(nums1[i - 1])
			} else {
				maxLeft = math.Max(float64(nums1[i - 1]), float64(nums2[j - 1]))
			}

			minRight := 0.0
			if i == m {
				minRight = float64(nums2[j])
			} else if j == n {
				minRight = float64(nums1[i])
			} else {
				minRight = math.Min(float64(nums1[i]), float64(nums2[j]))
			}
			if (m + n)%2 == 1 {
				return minRight
			} else {
				return (maxLeft + minRight) / 2.0
			}
		}
	}
	return 0.0
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		tmp := nums1
		nums1 = nums2
		nums2 = tmp
	}
	m, n := len(nums1), len(nums2)
	// 为什么 m+n+1，而不是m+n, 防止两个数组长度为1时，num为0
	num := (m + n + 1)/2
	// 0-m为m+1种切分的方法，二分查找在哪里切分，
	for low, high := 0, m; low <= high; {
		i := (low + high)/2
		j := num - i
		if j > 0 && i < m && nums2[j-1] > nums1[i] {
			low = i + 1
		} else if i > 0 && j < n && nums1[i-1] > nums2[j] {
			high = i - 1
		} else {
			maxLeft := 0.0
			if i == 0 {
				maxLeft = float64(nums2[j - 1])
			} else if j == 0 {
				maxLeft = float64(nums1[i - 1])
			} else {
				maxLeft = math.Max(float64(nums1[i - 1]), float64(nums2[j - 1]))
			}
			if (m + n)%2 == 1 {
				return maxLeft
			}
			minRight := 0.0
			if i == m {
				minRight = float64(nums2[j])
			} else if j == n {
				minRight = float64(nums1[i])
			} else {
				minRight = math.Min(float64(nums1[i]), float64(nums2[j]))
			}
			return (maxLeft + minRight)/2.0
		}
	}
	return 0.0
}

func findK(nums1, nums2 []int, k int) int {
	result := 0
	if k <= 0 {
		return result
	}
	for index,i,j := 0,0,0; index < k; index++ {
		if nums1[i] <= nums2[j] {
			result = nums1[i]
			i++
		} else {
			result = nums2[j]
			j++
		}
	}
	return result
}
