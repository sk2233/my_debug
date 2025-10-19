package main

import "fmt"

type Dog struct {
	Name string
}

func main() {
	fmt.Println("Hello World")
	d := new(Dog)
	fmt.Println(d)
}
