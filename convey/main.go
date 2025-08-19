package main

import (
	"fmt"

	"convey/mylib"
)

func main() {
	fmt.Printf("split test:%v\n", mylib.Split("a, b", ","))
}