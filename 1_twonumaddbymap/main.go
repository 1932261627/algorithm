package main

import(
	"fmt"
)

func main(){

	var a=[5]int{2,7,11,15}

	//使用map处理
	var tool=make(map[int]int)

	var target int=17

	var i int
	for i=0;i<len(a);i++{
		key:=target-a[i]
		res,exist:=tool[key]
		if exist{
			fmt.Println(res,i)
			break
		}
		tool[a[i]]=i
	}

	if i==len(a){
		fmt.Println("not found")
	}


}