package main

import "fmt"

func main() {
	helloPrefix := prefixer("Hello")
	fmt.Println(helloPrefix("Bob"))   // should print Hello Bob
	fmt.Println(helloPrefix("Maria")) // should print Hello Maria
}

func prefixer(in1 string) func(string) string {
	return func(in2 string) string {
		return in1 + in2
	}
}
