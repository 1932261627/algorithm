package main

/*
 * @lc app=leetcode.cn id=6 lang=golang
 *
 * [6] Z 字形变换
 */

// @lc code=start
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

// @lc code=end
