package main

import (
	"fmt"
	"runtime"
)

func main() {
	arr := [...]int{1, 2, 3, 4}
	var s [2]int
	fmt.Println("Hello world")
	fmt.Println(arr, s)
	fmt.Printf("OS: %s\nArchitecture: %s\n",runtime.GOOS, runtime.GOARCH)

}
