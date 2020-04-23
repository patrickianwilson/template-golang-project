package main

import "github.com/gopherjs/gopherjs/js"

func main() {
	//you can run this app by:
	//gradlew release
	//1. In Intellij, Run the build/assets/sample.html file or,
	//2. Host the build/assets folder in a web server
	js.Global.Get("document").Set("title", "Welcome to My Application")
	js.Global.Get("window").Call("alert", "Hello World")
}
