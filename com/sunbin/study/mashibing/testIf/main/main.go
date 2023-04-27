package main

import "fmt"

func main() {
	fmt.Println("main()")
	testIf(59)
	testIf(61)
	testIfElse(59)
	testIfElse(61)
	testElseIf(59)
	testElseIf(61)
	testElseIf(92)
}

func testIf(score int) {
	fmt.Println("testIf()")
	var result bool
	if score >= 60 {
		result = true
	}
	fmt.Printf("score: %d\t result: %v\n", score, result)
}

func testIfElse(score int) {
	fmt.Println("testIfElse()")
	var result bool
	if score < 60 {
		result = false
	} else {
		result = true
	}
	fmt.Printf("score: %d\t result: %v\n", score, result)
}

func testElseIf(score int) {
	fmt.Println("testElseIf()")
	var result string
	if score >= 90 {
		result = "A"
	} else if score >= 80 {
		result = "B"
	} else if score >= 70 {
		result = "C"
	} else if score >= 60 {
		result = "D"
	} else {
		result = "E"
	}
	fmt.Printf("score: %d\t result: %s\n", score, result)
}
