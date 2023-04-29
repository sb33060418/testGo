package main

import "fmt"

//多态1：多态参数，定义参数类型为接口，可以传入各种实现了接口的类的对象

// 多态2，多态数组
type Person interface {
	sayHello(name string)
}

type NamedPerson struct {
	name string
}

type Chinese struct {
	NamedPerson
}

func (c Chinese) sayHello(name string) {
	fmt.Println("你好，", name)
}

type American struct {
	NamedPerson
}

func (a American) sayHello(name string) {
	fmt.Println("Hello,", name)
}

func greet(p Person, name string) {
	p.sayHello(name)
}
func testPolymorphic() {
	fmt.Println("testPolymorphic")
	var array [3]Person
	a := American{}
	a.NamedPerson.name = "Mary"
	array[0] = a
	var c Chinese
	c.NamedPerson = NamedPerson{"李四"}
	array[1] = c
	fmt.Println(array)
}
func main() {
	fmt.Println("main")
	testPolymorphic()
}
