package main

import "fmt"

func main() {
	arr := []int{0} // 空だとpanic
	ret := make(map[int][]int)
	ret[0] = append(ret[0], arr[0])
	fmt.Println(ret)
	fmt.Println(arr[1:])
	ret[5] = append(ret[5], arr[1:]...)
	fmt.Println(ret)
}
