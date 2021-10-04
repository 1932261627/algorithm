package main

import (
	"fmt"
)

func main() {

	str := "abcabcbb"

	len := lengthOfLongestSubstring(str)

	fmt.Println(len)

}

func lengthOfLongestSubstring(s string) int {

	max := 0
	for i := 0; i < len(s); i++ {
		length := 0
		flag := false
	Loop:
		for j := i + 1; j < len(s); j++ {

			//最内侧的循环判断字符串是否重复

			for index := i; index < j; index++ {

				if s[index] == s[j] {
					length = j - i
					flag = true
					break Loop
				}

			}

		}

		if !flag {
			length = len(s) - i
		}

		if max < length {
			max = length
		}

	}

	return max
}
