package main

import "fmt"

func main() {
	greeting := greet()
	fmt.Println(greeting)
}

func greet() string {
	//return a greeting message
	return "Hello world"
}
