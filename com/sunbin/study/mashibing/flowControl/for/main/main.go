package main

import "fmt"

func main() {
	fmt.Println("main()")
	testRange()
	testRune()
}

func test() {
	fmt.Println("test()")
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

func test1() {
	fmt.Println("test1()")
	var i int = 0
	for ; i < 10; i++ {
		fmt.Println(i)
	}
}

func test2() {
	fmt.Println("test2()")
	var i int = 0
	for {
		if i < 10 {
			fmt.Println(i)
			i++
		} else {
			break
		}
	}
}

func testRange() {
	fmt.Println("testRange()")
	var str string = "hello 中国"
	for i, v := range str {
		fmt.Printf("i=%d, v=%c\n", i, v)
	}
}

func testRune() {
	fmt.Println("testRune()")
	var str string = "hello 中国"
	r := []rune(str)
	for i := 0; i < len(r); i++ {
		fmt.Printf("i=%d, v=%c\n", i, r[i])
	}
	for i, v := range r {
		fmt.Printf("i=%d, v=%c\n", i, v)
	}
}
