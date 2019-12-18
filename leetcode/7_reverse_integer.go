package main

import "fmt"

func main() {
	x := 1534236469
	fmt.Println(reverse(x))
}

func reverse(x int) int {
	// int is a signed integer type that is at least 32 bits in size.
	// It is a distinct type, however, and not an alias for, say, int32.
	// uint is a variable sized type, on your 64 bit computer uint is 64 bits wide.
	const INT_MAX = int(^uint32(0) >> 1)
	fmt.Println(INT_MAX)
	const INT_MIN = ^INT_MAX
	fmt.Println(INT_MIN)
	result := 0.0
	for ; x != 0 ; {
		tmp := x % 10
		x = x / 10
		result = result * 10 + float64(tmp)
		if result > float64(INT_MAX) || result < float64(INT_MIN) {
			result = 0
			break
		}
	}
	return int(result)
}