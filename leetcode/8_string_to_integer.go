package main

import "fmt"

func main() {
	str := " "
	fmt.Println(myAtoi(str))
}

func myAtoi(str string) int {
	const INT_MAX = int(^uint32(0) >> 1)
	const INT_MIN = ^INT_MAX
	result := 0.0
	flag := 1.0
	if len(str) == 0 {
		return int(result)
	}
	i := 0
	for ; i < len(str); i++ {
		if str[i] == ' ' {
			continue
		}
		break
	}
	if i == len(str) {
		return int(result)
	}
	if str[i] == '+' {
		i++
	} else if str[i] == '-' {
		i++
		flag = -1.0
	}
	for ; i < len(str); i++ {
		if str[i] <= '9' && str[i] >= '0' {
			result = result * 10 + float64(str[i]) - float64('0')
			if flag * result > float64(INT_MAX) {
				result = float64(INT_MAX)
				break
			} else if flag * result < float64(INT_MIN) {
				result = float64(INT_MIN)
				break
			}
			continue
		}
		break
	}
	if result < 0 {
		return int(result)
	} else {
		return int(flag * result)
	}
}