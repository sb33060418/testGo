package main

import "fmt"

type Person interface {
	sayHello(name string)
}

type Chinese struct {
}

func (c Chinese) sayHello(name string) {
	fmt.Println("你好，", name)
}

func (c Chinese) kongfu(name string) {
	fmt.Println(name, " 练功夫")
}

type American struct {
}

func (a American) sayHello(name string) {
	fmt.Println("Hello,", name)
}

func (a American) disco(name string) {
	fmt.Println(name, " 跳迪斯科")
}

func greet(p Person, name string) {
	p.sayHello(name)
	//// 报错
	//p.kongfu()
	//写法1
	//var c Chinese = p.(Chinese)
	//c.kongfu(name)
	////写法2
	//if c, flag := p.(Chinese); flag {
	//	c.kongfu(name)
	//} else {
	//	fmt.Println(name, "不是中国人，不会练功服")
	//}
	//wwitch type
	switch p.(type) {
	case Chinese:
		c := p.(Chinese)
		c.kongfu(name)
	case American:
		a := p.(American)
		a.disco(name)
	default:
		fmt.Println(name, "不是中国人，也不是美国人")

	}

}

func test() {
	fmt.Println("test")
	c := Chinese{}
	greet(c, "张三")
	a := American{}
	greet(a, "John")
	fmt.Println("test over")
}

func main() {
	fmt.Println("main")
	test()
}
