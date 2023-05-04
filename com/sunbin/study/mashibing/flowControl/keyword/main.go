package main

import "fmt"

func testBreak() {
	fmt.Println("testBreak()")
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
		if i > 20 {
			fmt.Println("i > 20")
			break
		}
		fmt.Printf("i: %d, sum: %d\n", i, sum)
	}
}

func testBreakDoubleLoop() {
	fmt.Println("testBreakDoubleLoop()")
	for i := 1; i <= 5; i++ {
		for j := 1; j <= 4; j++ {
			if i == 2 && j == 2 {
				//只跳出最内层循环
				break
			}
			fmt.Printf("i: %d, j: %d\n", i, j)
		}
	}
}
func testBreakLabel() {
	fmt.Println("testBreakLabel()")
	for k := 1; k <= 3; k++ {
	labelTest:
		for i := 1; i <= 5; i++ {
			for j := 1; j <= 4; j++ {
				if i == 2 && j == 2 {
					fmt.Println("break labelTest")
					// 跳出至指定循环
					break labelTest
				}
				fmt.Printf("k: %d, i: %d, j: %d\n", k, i, j)
			}
		}
	}
}

func testContinue() {
	fmt.Println("testContinue()")
	for i := 1; i <= 10; i++ {
		if i%3 == 0 {
			continue
		}
		fmt.Printf("i: %d\n", i)
	}

}

func testContinueDoubleLoop() {
	fmt.Println("testContinueDoubleLoop()")
	for i := 1; i <= 5; i++ {
		for j := 1; j <= 4; j++ {
			if i == 2 && j == 2 {
				//只跳出最内层循环
				continue
			}
			fmt.Printf("i: %d, j: %d\n", i, j)
		}
	}
}
func testContinueLabel() {
	fmt.Println("testContinueLabel()")
	for k := 1; k <= 3; k++ {
	labelTest:
		for i := 1; i <= 5; i++ {
			for j := 1; j <= 4; j++ {
				if i == 2 && j == 2 {
					fmt.Println("continue labelTest")
					// 跳出至指定循环
					continue labelTest
				}
				fmt.Printf("k: %d, i: %d, j: %d\n", k, i, j)
			}
		}
	}
}

func testGoto() {
	fmt.Println("testGoto()")
	for i := 1; i <= 5; i++ {
		if i == 3 {
			//只跳出最内层循环
			goto labelTest
		}
		fmt.Printf("i: %d\n", i)
	}
labelTest:
	fmt.Println("goto labelTest")
	fmt.Println("testGoto() end")
}
func testReturn() {
	fmt.Println("testReturn()")
	for i := 1; i <= 5; i++ {
		if i == 3 {
			fmt.Println("return")
			//只跳出最内层循环
			return
		}
		fmt.Printf("i: %d\n", i)
	}
	fmt.Println("testReturn() end")
}
func main() {
	fmt.Println("main()")
	//testBreak()
	//testBreakDoubleLoop()
	//testBreakLabel()
	//testContinue()
	//testContinueDoubleLoop()
	//testContinueLabel()
	//testGoto()
	testReturn()
}
