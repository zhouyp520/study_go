package main

import "fmt"

var (
	name string
	age  int
	isOK bool
)

const (
	pi = 3.1415
	pi1
	pi2
	n1 = iota
	n2
	n3
)

func main() {
	name = "JAM"
	age = 30
	isOK = true
	fmt.Printf(name)
	fmt.Println(age)
	fmt.Printf("name is :%s\n", name)
	fmt.Println(pi)
	fmt.Println(pi1)
	fmt.Println(pi2)
	fmt.Println(n1)
	fmt.Println(n2)
	fmt.Println(n3)
}
