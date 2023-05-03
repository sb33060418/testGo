package main

import (
	"bufio"
	"fmt"
	"os"
)

func testScanln() {
	fmt.Println("testScan()")
	var age int
	fmt.Println("Please input your age:")
	fmt.Scanln(&age)
	fmt.Printf("Your age is %d\n", age)
}

func testScanf() {
	fmt.Println("testScanf()")
	var name string
	var age int
	var score float64
	var isVip bool
	fmt.Println("Please input your name, age, score, isVip:")
	fmt.Scanf("%s %d %f %t", &name, &age, &score, &isVip)
	fmt.Printf("Your name is %s, age is %d, score is %f, isVip is %t\n", name, age, score, isVip)

}
func testStdin() {

	for {
		reader := bufio.NewReader(os.Stdin)
		sendStr, errReader := reader.ReadString('\n')
		if errReader != nil {
			fmt.Println("read input error:", errReader)
			return
		} else {
			fmt.Println("read input success:", sendStr)
		}
		if sendStr == "exit\r\n" {
			fmt.Println("read exit")
			return
		}
	}
}
func main() {
	fmt.Println("main()")
	//testScanln()
	//testScanf()
	testStdin()
}
