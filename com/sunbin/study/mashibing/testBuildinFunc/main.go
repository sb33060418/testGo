package main

import "fmt"

func main() {
	fmt.Println("main()")
	test()
}

func test() {
	// len
	s := "hello world"
	fmt.Println(len(s))
	// new
	p := new(int)
	fmt.Printf("类型：%T 地址：%v 值：%v 指向地址的值：%v\n", p, &p, p, *p)

}
