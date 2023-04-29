package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("main")
	num1 := 10
	num2 := 0
	test(num1, num2)
	err := testNewError(num1, num2)
	if err != nil {
		fmt.Println("自定义错误:", err)
		panic(err)
	}
	fmt.Println("main end")
}

func test(num1 int, num2 int) {
	fmt.Println("test")
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("err:", err)
		}
	}()
	result := num1 / num2
	fmt.Println("result:", result)
}

func testNewError(num1 int, num2 int) (err error) {
	fmt.Println("testNewError")
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("err:", err)
		}
	}()
	if num2 == 0 {
		return errors.New("除数为0")
	} else {
		result := num1 / num2
		fmt.Println("result:", result)
		return nil
	}
}
