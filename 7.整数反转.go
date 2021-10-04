package main

/*
 * @lc app=leetcode.cn id=7 lang=golang
 *
 * [7] 整数反转
 */

// @lc code=start
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

	if res > 2147483647 || res < (-2147483647) {
		return 0
	}

	return res
}

// @lc code=end
