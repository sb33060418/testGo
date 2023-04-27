package main

import "fmt"

func main() {
	fmt.Println("main()")
	testPointer()
}

func testPointer() {
	fmt.Println("testPointer()")
	var a int = 10
	var p *int = &a
	fmt.Printf("%T %p %v %v\n", a, a, a, &a)
	fmt.Printf("%T %p %v %v\n", p, p, p, *p)
	*p = 20
	fmt.Println(a)
	a = 30
	fmt.Println(a)
}
