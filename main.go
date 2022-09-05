// print the text hello world to the terminal

package main

import "fmt"

func main() {
	// var whatToSay string
	// whatToSay = "Hello, World!"
	whatToSay := "Hello, World!"
	sayHelloWorld(whatToSay)
}

func sayHelloWorld(whatToSay string) {
	fmt.Println(whatToSay)
}
