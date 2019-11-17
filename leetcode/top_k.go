package main

import (
	"fmt"
)

func main() {
	arr := [16]int{1, 17, 3, 4, 5, 6, 7, 16, 9, 10, 11, 12, 13, 14, 15, 8}
	k := 5
	topK := createHeap(arr, k)
	fmt.Println(topK)
	for i := k; i < len(arr); i++ {
		sort(topK, arr[i])
	}
	fmt.Println(topK)
	fmt.Println(printHeap(topK))
}

// 当前节点之前已经是小顶堆了，新来一个元素要逐层向上调整小顶堆
// 也就是新来的元素放在最后
func createHeap(array [16]int, k int) []int {
	// slice 要先创建
	top := make([]int, k)
	for i := 0; i < k; i++ {
		top[i] = array[i]
		makeHeap(top, i)
	}
	return top
}

// i的父节点为(i-1)/2，i的子节点为2i+1和2i+2
func makeHeap(top []int, index int) {
	for parent := (index - 1) / 2; parent >= 0; {
		if top[parent] > top[index] {
			tmp := top[parent]
			top[parent] = top[index]
			top[index] = tmp
			index = parent
			parent = (index - 1) / 2
		} else {
			break
		}
	}
}

func sort(topK []int, target int) {
	if target > topK[0] {
		topK[0] = target
		heap(topK)
	}
}

func heap(topK []int)  {
	for i := 0; i < len(topK); i++ {
		makeHeap(topK, i)
	}
}

func printHeap(topK []int) []int {
	sort := make([]int, len(topK))
	for i := 0; len(topK) > 0; i++ {
		sort[i] = topK[0]
		topK = topK[1:]
		heap(topK)
	}
	return sort
}