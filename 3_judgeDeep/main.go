package main

import(
	"fmt"
	"sort"
)

func maxDepth(n int) int {
	var depth int
	for i := n; i > 0; i >>= 1 {
		depth++
	}
	return depth * 2
}

func main(){

	res:=maxDepth(-5)
	fmt.Println(res)

	a:=[]int{1,2,3}

	sort.Ints(a)

}