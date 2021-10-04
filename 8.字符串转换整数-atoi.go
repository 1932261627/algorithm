package main

/*
 * @lc app=leetcode.cn id=8 lang=golang
 *
 * [8] 字符串转换整数 (atoi)
 */

// @lc code=start
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

// @lc code=end
