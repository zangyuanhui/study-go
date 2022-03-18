package main

import (
	"fmt"
	"rsc.io/quote"
)

func main() {
	//log.SetPrefix("greetings: ")
	//log.SetFlags(0)
	//
	//names := []string{"zangyh", "zyh", "zangyuanhui"}
	//
	//// Get a greeting message and print it.
	//messages, err := greetings.Hellos(names)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(messages)
	fmt.Println(quote.Go())
}
