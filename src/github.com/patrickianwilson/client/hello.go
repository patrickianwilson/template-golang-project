package main

import "github.com/gopherjs/gopherjs/js"

func main() {
	js.Global.Get("document").Set("title", "Welcome to My Application")
	js.Global.Get("window").Call("alert", "Hello World")
}
