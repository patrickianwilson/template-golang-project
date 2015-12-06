package main

import "fmt"

func main() {
	fmt.Printf("hello, world\n")
	_ = "breakpoint" //this is a "godebug" breakpoint.
	fmt.Printf("After Breakpoint")
}
