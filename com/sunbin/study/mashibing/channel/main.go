package main

import "fmt"

var intChan chan int

func test() {
	fmt.Println("test()")
	intChan = make(chan int, 3)
	fmt.Printf("intChan本身的地址=%p, intChan 的值=%v\n", intChan, &intChan)

	intChan <- 10
	num := 20
	intChan <- num
	intChan <- 30
	fmt.Printf("len:%v, cap:%v\n", len(intChan), cap(intChan))

	num1 := <-intChan
	num2 := <-intChan
	num3 := <-intChan
	fmt.Println("num1=", num1, "num2=", num2, "num3=", num3)
	defer deferRecover()
	num4 := <-intChan
	fmt.Println("num4=", num4)
}

func deferRecover() {
	error := recover()
	if error != nil {
		fmt.Println("error=", error)
	}
}
func testCloseChan() {
	intChan = make(chan int, 3)
	intChan <- 10
	intChan <- 20
	close(intChan)

	num1 := <-intChan
	fmt.Println("num1=", num1)

	defer deferRecover()
	intChan <- 30 // panic: send on closed channel
}

func testLoopChan() {
	len := 5
	intChan = make(chan int, len)

	for i := 0; i < len; i++ {
		intChan <- i
	}
	// 不关闭会报错：fatal error: all goroutines are asleep - deadlock!
	close(intChan)
	for v := range intChan {
		fmt.Println("v=", v)
	}
}

func main() {
	fmt.Println("main()")
	//test()
	//testCloseChan()
	//testLoopChan()
	testChanRead()
	testChanWrite()
}

func testChanRead() {
	fmt.Println("")
	var intChanRead <-chan int = make(<-chan int, 3)
	if intChanRead != nil {
		//fatal error: all goroutines are asleep - deadlock!
		num1 := <-intChanRead
		fmt.Println("num1:", num1)
	}

}

func testChanWrite() {
	fmt.Println("testReadWriteOnlyChan()")
	var intChanWrite chan<- int = make(chan<- int, 3)
	intChanWrite <- 10
	//num1 := <- intChanWrite

}
