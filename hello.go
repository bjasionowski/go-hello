package main

import (
	"fmt"
	"github.com/bjasionowski/stringutil"
)

func main() {
	msg := "Hello, world!"
	fmt.Println(msg)
	fmt.Println(stringutil.Reverse(msg))
}
