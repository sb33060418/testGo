package main

import "fmt"

func testSlice() {
	fmt.Println("testSlice()")
	var intArr [6]int = [6]int{1, 2, 3, 4, 5, 6}
	fmt.Println("intArr:", intArr)
	fmt.Printf("%v %T %p\n", intArr, intArr, intArr)
	slice := intArr[1:3]
	fmt.Println("slice:", slice)
	fmt.Printf("%v %T %p\n", slice, slice, slice)
	fmt.Println("len:", len(slice))
	fmt.Println("cap:", cap(slice))
}

func main() {
	fmt.Println("main()")
	testSlice()
}
