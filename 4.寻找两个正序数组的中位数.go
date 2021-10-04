package main

/*
 * @lc app=leetcode.cn id=4 lang=golang
 *
 * [4] 寻找两个正序数组的中位数
 */

// @lc code=start
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

// @lc code=end
