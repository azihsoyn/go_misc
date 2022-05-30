package main

import "fmt"

func main(){
	ms := map[int][]int{}
	for i := 0; i < 10; i++ {
		ms[i] = append(ms[i], i)
	}
	fmt.Println(ms)
}
