package main

import "fmt"

func main() {
	// 这是一个slice，slice传进去直接在原slice上修改，直接传slice就行，不用指针
	//var arr = []int{5,2,8,4,3,9,7,5,10}
	var arr = [9]int{5,2,8,4,3,9,7,5,10}
	quickSort(&arr, 0, len(arr)-1)
	fmt.Println(arr)
}

func quickSort(arr *[9]int, start, end int) {
	if start < end {
		index := divide(arr, start, end)
		quickSort(arr, start, index - 1)
		quickSort(arr, index + 1, end)
	}
}

func divide(arr *[9]int, start, end int) (index int) {
	target := (*arr)[start]
	for i, j := start, end; i < j; {
		for ; j >= i && target <= (*arr)[j]; j-- {}
		if j < i {
			index = i
			break
		}
		(*arr)[i] = (*arr)[j]
		for ; i <= j && target > (*arr)[i]; i++ {}
		if i > j {
			(*arr)[j] = target
			index = j
			break
		} else {
			(*arr)[j] = (*arr)[i]
		}
	}
	return
}
