package main

import (
	"fmt"
)

func main() {

	//四题
	// nums1 := []int{1, 2}
	// nums2 := []int{-1, 3}
	// res := findMedianSortedArrays(nums1, nums2)
	// fmt.Println(res)

	//六题
	// s := "AB"
	// numRows := 1
	// res := convert(s, numRows)
	// fmt.Println(res)

	//七题
	// res := reverse(-1273)
	// fmt.Println(res)

	//八题
	res := myAtoi("9223372036854775808")
	fmt.Println(res)

}

//八题
func myAtoi(s string) int {

	res := 0

	//索引
	index := 0

	//正负的标志
	flag := true

	if len(s) == 0 {
		return 0
	}

	//去掉空格
	for {
		if s[index] == ' ' {
			index++
			if index >= len(s) {
				break
			}
		} else {

			//只能是正负号或者数字
			if s[index] != '-' && s[index] != '+' && (s[index] < '0' || s[index] > '9') {
				return 0
			}

			//如果是正负号，下一个数必须是数字
			if s[index] == '-' || s[index] == '+' {

				//判断下一个数
				if index+1 < len(s) && (s[index+1] < '0' || s[index+1] > '9') {
					return 0

				}

				//负号进行标志
				if s[index] == '-' {
					flag = false

				}
				index++

			}
			break

		}
	}

	//开始进行数字转化
	if index < len(s) {

		for {
			res = res*10 + (int(s[index]) - 48)
			if res >= 2147483647 || res < -2147483648 {

				break
			}

			index++

			//判断结束条件
			if index >= len(s) || s[index] < '0' || s[index] > '9' {
				break
			}

		}
	}

	//判断正负
	if !flag {
		res = -res
	}

	fmt.Println(res)

	// 	//考虑超过范围
	//2,147,483,648-1
	if res >= 2147483647 {

		return 2147483647
	}

	if res < -2147483648 {
		return -2147483648
	}

	return res
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	//中位数-分奇偶
	len1 := len(nums1)
	len2 := len(nums2)

	i := 0
	j := 0
	res := 0.0

	if (len1+len2)%2 == 0 {

		//偶数情况
		for {
			//移动指针
			if len1 == 0 && len2 == 0 {
				return 0
			} else if i == len1 && j < len2 {

				if (i + j + 1) == (len1+len2)/2 {
					res = float64(nums2[j]+nums2[j+1]) / 2
					break
				}
				j++

			} else if i < len1 && j == len2 {
				if (i + j + 1) == (len1+len2)/2 {
					res = float64(nums1[i]+nums1[i+1]) / 2
					break

				}
				i++

			} else {
				if nums1[i] > nums2[j] {

					//nums[j]为第一个
					if (i + j + 1) == (len1+len2)/2 {

						if j+1 < len2 && nums1[i] > nums2[j+1] {
							res = float64(nums2[j]+nums2[j+1]) / 2

						} else {
							res = float64(nums2[j]+nums1[i]) / 2

						}
						break

					}
					j++

				} else {
					if (i + j + 1) == (len1+len2)/2 {

						//nums[i]为第一个
						if i+1 < len1 && nums2[j] > nums1[i+1] {
							res = float64(nums1[i]+nums1[i+1]) / 2

						} else {
							res = float64(nums1[i]+nums2[j]) / 2

						}

						break

					}
					i++

				}
			}

		}
		return res

	} else {

		//奇数情况-至少有一个数
		for {
			//移动指针
			if i == len1 && j < len2 {

				if (i + j) == (len1+len2)/2 {
					res = float64(nums2[j])
					break
				}
				j++

			} else if i < len1 && j == len2 {
				if (i + j) == (len1+len2)/2 {
					res = float64(nums1[i])
					break

				}
				i++

			} else {
				if nums1[i] > nums2[j] {
					if (i + j) == (len1+len2)/2 {
						res = float64(nums2[j])
						break

					}
					j++

				} else {
					if (i + j) == (len1+len2)/2 {
						res = float64(nums1[i])
						break

					}
					i++

				}
			}

		}
		return res

	}

}

func convert(s string, numRows int) string {

	index := []int{}
	flag := 1

	headOrBack := true

	for i := 0; i < len(s); i++ {
		index = append(index, flag)

		if headOrBack && flag < numRows {
			flag++

		} else if !headOrBack && flag > 1 {
			flag--

		}

		if flag == numRows || flag == 1 {
			headOrBack = !headOrBack
		}
	}

	//遍历输出结果
	var res string
	for i := 1; i <= numRows; i++ {
		for j := 0; j < len(index); j++ {
			if index[j] == i {
				res += string(s[j])
			}
		}
	}

	return res

}

func reverse(x int) int {

	res := 0
	for {
		a := x % 10
		x = x / 10

		res = res*10 + a

		//结束条件
		if x == 0 {
			break
		}

	}

	return res
}
