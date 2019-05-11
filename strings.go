package main

import (
	"unicode/utf8"
	"fmt"
)

func main() {
	s := "Hello盛钧白!"
	for _, b := range []byte(s) {
		fmt.Printf("%X ", b)
	}
	fmt.Println()
	for i, ch := range s {
		fmt.Printf("(%d, %x) ", i, ch)
    }
	fmt.Println()
    fmt.Println("Rune count:",utf8.RuneCountInString(s))    
}
