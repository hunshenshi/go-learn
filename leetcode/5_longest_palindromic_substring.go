package main

import "fmt"

func main() {
	s := "abba"
	//fmt.Println(s[1:2+1])
	fmt.Println(longestPalindrome(s))
}

func longestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}

	start, end, max := 0, 0, 0
	for i := 0; i < len(s); i++ {
		// 这里以i为中心向外扩展，只考虑了奇数的情况，
		// 当偶数时，中心点不是一个数，而是以两个数为中心
		for j := i; j < len(s); j++ {
			left := i - (j - i)
			//if left < 0 || s[left] != s[j] {
			if left < 0 {
				break
			} else if s[left] != s[j] {
				break
			} else {
				if (j - left + 1) > max {
					max = j - left + 1
					start, end = left, j
				}
			}
		}
        // 判断偶数情况下的中心扩散
		for j := i+1; j < len(s); j++ {
			left := i - (j - (i + 1))
			//if left < 0 || s[left] != s[j] {
			if left < 0 {
				break
			} else if s[left] != s[j] {
				break
			} else {
				if (j - left + 1) > max {
					max = j - left + 1
					start, end = left, j
				}
			}
		}
	}
	return s[start:end+1]
}