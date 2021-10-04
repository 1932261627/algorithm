package main

import("fmt")

func main(){

	var num = []int{2,7,11,15}
	var target int=17

	for i:=0;i<len(num)-1;i++{
		for j:=i+1;j<len(num);j++{
			if num[i]+num[j]==target{
				fmt.Println(i,j)
				break
			}
		}
	}

	fmt.Println("--")
}