package main

import (
	"github.com/kr/pretty"
	"time"
)

type Article struct {
	Author string
}

type Video struct {
	Duration time.Duration
}

type Announce struct {
	Title string
}

type Content interface {
	Article | Video | Announce
}

func getArticleContent[C Content]() C {
	var c C
	return c
}

func main() {
	c := getArticleContent[Article]()
	pretty.Println(c)
}
