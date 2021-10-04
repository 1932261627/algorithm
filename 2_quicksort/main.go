package main

import(
	"fmt"
)

//选择左边第一个为pivot,空出一个区域
//开始判断时从右边进行判断

//快排的递归函数
func quickSort(a *[7]int, L int, R int) int {

	pivot := (*a)[L]
	// L++
Loop:
	for {
		//从右边进行判断
		if (*a)[R] >= pivot {

			for {
				R--
				if (*a)[R] < pivot {
					break
				}

				if L >= R {
					break Loop
				}

			}
		}

		if (*a)[R] < pivot {
			(*a)[L], (*a)[R] = (*a)[R], (*a)[L]
			L++
			if L >= R {
				break
			}
		}

		if (*a)[L] <= pivot {
			for {
				L++
				if (*a)[L] > pivot {
					break
				}
				if L >= R {
					break Loop
				}
			}
		}

		if (*a)[L] > pivot {
			//(*a)[R], (*a)[L] = (*a)[L], (*a)[R]
			(*a)[L], (*a)[R] = (*a)[R], (*a)[L]
			R--
			if L >= R {
				break
			}
		}

	}

	(*a)[L] = pivot
	return L

}

//快速排序
func sort(a *[7]int, L int, R int) {

	index := quickSort(a, L, R)

	if L < index-1 {
		sort(a, L, index-1)
	}

	if R > index+1 {
		sort(a, index+1, R)
	}

}

//选择排序
func selectSort(arr []int, first int, last int) {

	for i := 0; i < len(arr); i++ {
		for j := i; j < len(arr); j++ {
			if arr[j] < arr[i] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
}

//冒泡排序
func bubbleSort(arr []int, first int, last int) {

	for i := len(arr) - 1; i >= 0; i-- {
		for j := 0; j < i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

//注意点
//==2
// fmt.Println(5/2)

//希尔排序
func shellSort(arr []int, first int, last int) {

	//确定增量
	for i := (last - first) / 2; i >= 1; i = i / 2 {
		//选择排序
		for m := 0; m <= len(arr)-i; m = m + i {
			for n := m + i; n <= len(arr)-i; n = n + i {
				if arr[m] > arr[n] {
					arr[m], arr[n] = arr[n], arr[m]
				}

			}

		}

		//冒泡排序-有问题
		// for m := 0; m < (last - first)/i; m++ {
		// 	for n := m; n <= last-i; n = n + i {
		// 		if arr[n] > arr[n+i] {
		// 			arr[n], arr[n+i] = arr[n+i], arr[n]
		// 		}
		// 	}
		// }

	}

}

//堆排序
//编号从0开始
//2i+1   2i+2
func heapSort(arr []int, first int, last int) {

	//多次循环获取最大值
	for j := last; j >= 1; j-- {
		//最大移动到数组头部
		for i := j; i>=0; i-- {
			//判断左子节点
			//(2*i+1 <= last)
			//i/2
			// if (2*i+1 <= last) && arr[2*i+1] > arr[i] {
			// 	arr[i], arr[2*i+1] = arr[2*i+1], arr[i]
			// }

			// //判断右子节点
			// //(2*i+2) <= last
			// //i/2-1
			// if (2*i+2) <= last && arr[2*i+2] > arr[i]  {
			// 	arr[i], arr[2*i+2] = arr[2*i+2], arr[i]
			// }

			// if (i/2)>=0 && arr[i/2]<arr[i]{
			// 	arr[i/2], arr[i] = arr[i], arr[i/2]
			// }

			//判断右子节点
			if i%2==0 && (i/2-1)>=0 && arr[i/2-1]<arr[i]{
				arr[i/2-1], arr[i] = arr[i], arr[i/2-1]
			}

			//判断左子节点
			if i%2!=0 && (i/2)>=0 && arr[i/2]<arr[i]{
				arr[i/2], arr[i] = arr[i], arr[i/2]
			}


		}
		//最大值与尾部替换
		arr[first], arr[j] = arr[j], arr[first]
	}

}

func main() {

	//测试选择,冒泡排序
	// a := []int{9, 7, 8, 5, 3, 1, 1}
	// fmt.Println(a)
	// // selectSort(a, 0, 6)
	// bubbleSort(a, 0, 6)
	// fmt.Println(a)

	//测试希尔排序
	// a := []int{9, 7, 8, 5, 3, 2, 5, 5}
	// fmt.Println(a)
	// shellSort(a, 1, 8)
	// fmt.Println(a)


	//堆排序
	a := []int{9,8,7,8,3, 7}
	fmt.Println(a)
	heapSort(a, 0, 5)
	fmt.Println(a)



	//测试快速排序
	// a := [7]int{9, 7, 8, 4, 3, 1, 1}
	// fmt.Println(a)
	// sort(&a, 0, 6)
	// fmt.Println(a)

}
