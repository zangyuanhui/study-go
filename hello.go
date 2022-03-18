package main

import (
	"fmt"
	"github.com/google/go-cmp/cmp"
	"zangyuanhui/study-go/hello"
)

func main() {
	fmt.Println(hello.ReverseRunes("!oG ,olleH"))
	fmt.Println(cmp.Diff("Hello World", "Hello Go"))
}
