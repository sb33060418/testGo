package main

import "fmt"

func testSwitch() {
	fmt.Println("testSwitch()")
	var score int
	for {
		fmt.Println("Please input your score:")
		fmt.Scanln(&score)
		fmt.Printf("Your score is %d\n", score)
		switch score / 10 {
		case 6:
			fmt.Println("D")
		case 7:
			fmt.Println("C")
		case 8:
			fmt.Println("B")
		case 9, 10:
			fmt.Println("A")
		default:
			fmt.Println("E")
		}
	}
}

func main() {
	fmt.Println("main()")
	testSwitch()
}
