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

type American struct {
}

func (a American) sayHello(name string) {
	fmt.Println("Hello,", name)
}

func greet(p Person, name string) {
	p.sayHello(name)
}

func test() {
	fmt.Println("test")
	c := Chinese{}
	a := American{}
	greet(c, "张三")
	greet(a, "John")
}

// 不能创建接口对象，必须指向具体的对象
func test1() {
	fmt.Println("test1")
	//// error
	//var p Person
	//p.sayHello("张三")
	//// 1
	//c := Chinese{}
	//var p Person = c
	// 2
	var p Person = Chinese{}

	p.sayHello("张三")
}

//自定义数据类型

type integer int

func (i integer) sayHello(name string) {
	fmt.Println("123,", name)
}

func test2() {
	fmt.Println("test2")
	var i integer = 10
	i.sayHello("456")
}

// 实现多个接口
type AInterface interface {
	a()
}

type BInterface interface {
	b()
}

type Stu struct {
}

func (s Stu) a() {
	fmt.Println("Stu.a()")
}

func (s Stu) b() {
	fmt.Println("Stu.b()")
}

func testMultiInterface() {
	fmt.Println("testMultiInterface")
	var s Stu
	s.a()
	s.b()
	var a AInterface = s
	var b BInterface = s
	a.a()
	b.b()
}

// 接口实现接口
type ABInterface interface {
	AInterface
	BInterface
	ab()
}

type ABChild struct {
}

func (abChild ABChild) a() {
	fmt.Println("abChild.a()")
}

func (abChild ABChild) b() {
	fmt.Println("abChild.b()")
}

func (abChild ABChild) ab() {
	fmt.Println("abChild.ab()")
}

func testSubInterface() {
	fmt.Println("testSubInterface")
	var abChild ABChild
	var a AInterface = abChild
	var b BInterface = abChild
	var ab ABInterface = abChild
	abChild.a()
	abChild.b()
	abChild.ab()
	a.a()
	b.b()
	ab.a()
	ab.b()
	ab.ab()

}

// 测试接口对象是指针（引用类型）
func testInterfaceObjectIsPointer() {
	fmt.Println("testInterfaceObjectIsPointer")
	var a American
	fmt.Println(a)
	fmt.Printf("%T %v %v %p", a, a, &a, &a)
}

// 任何类型都可视为实现任何空接口
type EmptyInterface interface {
}

func testEmptyInterface() {
	var s Stu
	var e EmptyInterface = s
	fmt.Println(e)
	e = 1
	fmt.Println(e)

}

func main() {
	fmt.Println("testInterface.main")
	//test()
	//test1()
	//test2()
	//testMultiInterface()
	//testSubInterface()
	//testInterfaceObjectIsPointer()
	testEmptyInterface()
}
