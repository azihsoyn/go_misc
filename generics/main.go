package main

import (
	"fmt"
	"github.com/azihsoyn/go_misc/generics/pkg"
	"github.com/kr/pretty"
)

func Myprint[T any](s ...T) string {
	for _, v := range s {
		fmt.Print(v)
	}
	return "hoge"
}

func main() {
	fmt.Println(Myprint[string]("hoge"))
	p1 := pkg.NewGenericsStruct[int, int]()
	fmt.Println(p1)
	p1.Set(10, 100)
	pretty.Println(p1.Get(1))
	pretty.Println(p1.Get(10))
	/*
		fmt.Println(SumIntsOrFloats[int, int64](map[int]int64{1: 1, 2: 2, 3: 3}))
	*/
}

/*
func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}
*/
