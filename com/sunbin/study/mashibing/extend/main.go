package main

import "fmt"

type Animal struct {
	Age    int64
	Weight float64
}

func (a *Animal) Shout() {
	fmt.Println("Animal.show()")
}

func (a *Animal) ShowInfo() {
	fmt.Printf("animal info: Age: %v, Weight: %v\n", a.Age, a.Weight)
}

type Cat struct {
	Animal
	Name string
}

func (c Cat) ShowInfo() {
	fmt.Printf("animal info: Name: %v, Age: %v, Weight: %v\n", c.Name, c.Age, c.Weight)
}
func (c *Cat) scratch() {
	fmt.Println("Cat.scratch()")
}

func test() {
	fmt.Println("test()")
	cat := Cat{}
	fmt.Println(cat)
	cat.Animal.Age = 3
	cat.Animal.Weight = 33.3
	cat.Name = "咪咪"
	fmt.Println(cat)
	cat.Shout()
	cat.ShowInfo()
	cat.scratch()
	cat.Animal.Shout()
	cat.Animal.ShowInfo()
}

func main() {
	fmt.Println("main()")
	test()
}
