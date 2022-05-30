package main

import (
	"errors"
	"fmt"
)

type MyError struct {
	Message string
}

func (e MyError) Error() string {
	return e.Message
}

func hoge() error {
	return MyError{Message: "fuga"}
}

func main() {
	err := hoge()
	if errors.Is(err, MyError{Message: "fuga"}) {
		fmt.Println("ok")
	} else {
		fmt.Println("ng")
	}

	if errors.As(err, &MyError{}) {
		fmt.Println("ok")
	} else {
		fmt.Println("ng")
	}

}
