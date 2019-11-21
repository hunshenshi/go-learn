package main

import (
	"fmt"
)

func main() {
//输入: "abcabcbb"
//输出: 3
//解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
	s := "pwwkew"
	fmt.Println(lengthOfLongestSubstring(s))
}

func lengthOfLongestSubstring(s string) int {
	length := 0
	if len(s) == 0 {
		return length
	}
	start := 0
	read := make(map[rune]int)
	for index, value := range s {
		//length = maxInt(index - start, length)
		if i, ok := read[value]; ok { // i < index
			if i >= start {
				length = maxInt(index - start, length)
				start = i + 1
			}
		} else {
			length = maxInt(index - start + 1, length)
		}
		//end = index - start
		read[value] = index
	}
	return length
}

func maxInt(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}