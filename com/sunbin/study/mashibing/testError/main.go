package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("main() start")

	num1 := 10
	num2 := 0

	fmt.Println("testError() start")
	err := testError(num1, num2)
	if err != nil {
		fmt.Println("err=", err)
		panic(err)
	}
	fmt.Println("testError() over")

	fmt.Println("testDefer() start")
	testDefer(num1, num2)
	fmt.Println("testDefer() over")

	fmt.Println("test() start")
	test(num1, num2)
	fmt.Println("test() over")

	fmt.Println("main() over")
}

func test(num1 int, num2 int) {
	fmt.Println("test()")
	res := num1 / num2
	fmt.Println("res=", res)
}

func testDefer(num1 int, num2 int) {
	fmt.Println("testDefer()")
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("err=", err)
		}
	}()
	res := num1 / num2
	fmt.Println("res=", res)
}

func testError(num1 int, num2 int) (err error) {
	fmt.Println("testError()")
	if num2 == 0 {
		return errors.New("num2 is 0")
	} else {
		res := num1 / num2
		fmt.Println("res=", res)
		return nil
	}
}
