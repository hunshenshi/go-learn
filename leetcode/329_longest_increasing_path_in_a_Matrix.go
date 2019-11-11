package main

import (
	"fmt"
)

func main() {
	//[
	//[9,9,4],
	//[6,6,8],
	//[2,1,1]
	//]
	matrix := [][]int{{9,9,4},
		{6,6,8},
		{2,1,1}}
	//matrix := [][]int{{3,4,5},
	//	{3,2,6},
	//	{2,2,1}}
	//matrix := [][]int{{0,1,2,3,4,5,6,7,8,9},
	//	{19,18,17,16,15,14,13,12,11,10},
	//	{20,21,22,23,24,25,26,27,28,29},
	//	{39,38,37,36,35,34,33,32,31,30},
	//	{40,41,42,43,44,45,46,47,48,49},
	//	{59,58,57,56,55,54,53,52,51,50},
	//	{60,61,62,63,64,65,66,67,68,69},
	//	{79,78,77,76,75,74,73,72,71,70},
	//	{80,81,82,83,84,85,86,87,88,89},
	//	{99,98,97,96,95,94,93,92,91,90},
	//	{100,101,102,103,104,105,106,107,108,109},
	//	{119,118,117,116,115,114,113,112,111,110},
	//	{120,121,122,123,124,125,126,127,128,129},
	//	{139,138,137,136,135,134,133,132,131,130},
	//	{0,0,0,0,0,0,0,0,0,0}}
	fmt.Println(longestIncreasingPath(matrix))
}

func longestIncreasingPath(matrix [][]int) int {
	if len(matrix) == 0 {
		return 0
	}

	m, n := len(matrix), len(matrix[0])
	longest := 0
	// 二维slice声明
	cache := make([][]int, m)
	for i := range cache {
		cache[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			//length := findIncrease(matrix, cache, i, j)
			//longest = maxInt(length, longest)
			//length := 0
			right := findIncrease(matrix, cache, i, j, i, j+1) + 1
			left := findIncrease(matrix, cache, i, j, i, j-1) + 1
			up := findIncrease(matrix, cache, i, j, i-1, j) + 1
			down := findIncrease(matrix, cache, i, j, i+1, j) + 1
			cache[i][j] = maxInt(maxInt(right, left), maxInt(up, down))
			longest = maxInt(longest, cache[i][j])
		}
	}
	return longest
}

// 回溯(dfs) + cache
func findIncrease(matrix, cache [][]int, i, j, targetI, targetJ int) int {
	if cache[i][j] > 0 {
		return cache[i][j]
	}
	if targetI < 0 || targetJ < 0 || targetI >= len(matrix) || targetJ >= len(matrix[0]) || matrix[i][j] >= matrix[targetI][targetJ] {
		return 0
	}
	if matrix[i][j] < matrix[targetI][targetJ] {
		right := findIncrease(matrix, cache, targetI, targetJ, targetI, targetJ+1) + 1
		left := findIncrease(matrix, cache, targetI, targetJ, targetI, targetJ-1) + 1
		up := findIncrease(matrix, cache, targetI, targetJ, targetI-1, targetJ) + 1
		down := findIncrease(matrix, cache, targetI, targetJ, targetI+1, targetJ) + 1
		cache[i][j] = maxInt(maxInt(right, left), maxInt(up, down))
		return cache[i][j]
	}
	return 0
}

// 回溯
//func findIncrease(matrix [][]int, i, j, targetI, targetJ, length int) int {
//	if targetI < 0 || targetJ < 0 || targetI >= len(matrix) || targetJ >= len(matrix[0]) {
//		return length
//	}
//	if matrix[i][j] < matrix[targetI][targetJ] {
//		right := findIncrease(matrix, targetI, targetJ, targetI, targetJ+1, length+1)
//		left := findIncrease(matrix, targetI, targetJ, targetI, targetJ-1, length+1)
//		up := findIncrease(matrix, targetI, targetJ, targetI-1, targetJ, length+1)
//		down := findIncrease(matrix, targetI, targetJ, targetI+1, targetJ, length+1)
//		return maxInt(maxInt(right, left), maxInt(up, down))
//	}
//	return length
//}

func maxInt(x, y int) int {
	if x < y {
		return y
	}
	return x
}