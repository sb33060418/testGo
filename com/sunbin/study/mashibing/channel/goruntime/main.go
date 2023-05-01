package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("main()")
	//testReadWrite()
	//testJustWrite()
	//testWriteFastReadSlow()
	//testSelect()
	testRuntineDefer()
}

func testReadWrite() {
	fmt.Println("testReadWriteData()")
	maxLength := 5
	intChan := make(chan int, maxLength)
	wg.Add(2)
	go writeData(intChan, maxLength)
	go readData(intChan)
	wg.Wait()
}
func testJustWrite() {
	fmt.Println("testJustWrite()")
	maxLength := 5
	intChan := make(chan int, maxLength)
	wg.Add(2)
	go writeData(intChan, maxLength)
	// 只写不读，阻塞：fatal error: all goroutines are asleep - deadlock!
	//go readData(intChan)
	wg.Wait()
}
func testWriteFastReadSlow() {
	fmt.Println("testWriteFastReadSlow()")
	maxLength := 5
	intChan := make(chan int, maxLength)
	wg.Add(2)
	// 写快读慢，只会等待不会报错
	go writeDataNoWait(intChan, maxLength*2)
	go readData(intChan)
	wg.Wait()
}

func writeData(intChan chan int, maxLength int) {
	fmt.Println("writeData()")
	defer wg.Done()
	for i := 0; i < maxLength; i++ {
		intChan <- i
		fmt.Println("writeData:", i)
		time.Sleep(time.Second)
	}
	close(intChan)
}

func writeDataNoWait(intChan chan int, maxLength int) {
	fmt.Println("writeData()")
	defer wg.Done()
	for i := 0; i < maxLength; i++ {
		intChan <- i
		fmt.Println("writeData:", i)
		//time.Sleep(time.Second)
	}
	close(intChan)
}

func readData(intChan chan int) {
	fmt.Println("readData()")
	defer wg.Done()
	for v := range intChan {
		fmt.Println("readData:", v)
		time.Sleep(time.Second)
	}
}

func testSelect() {
	fmt.Println("testSelect()")
	intChan := make(chan int, 1)
	go func() {
		for i := 0; i < 5; i++ {
			time.Sleep(time.Second * 5)
			intChan <- i
			fmt.Println("intChan write:", i)
		}
		close(intChan)
		fmt.Println("intChan close")
	}()
	stringChan := make(chan string, 1)
	go func() {
		for i := 0; i < 5; i++ {
			time.Sleep(time.Second * 2)
			value := "hello " + strconv.Itoa(i)
			stringChan <- value
			fmt.Println("stringChan write:", value)
		}
		close(stringChan)
		fmt.Println("stringChan close")
	}()

	select {
	case v := <-intChan:
		fmt.Println("intChan select:", v)
	case v := <-stringChan:
		fmt.Println("stringChan select:", v)
	default:
		time.Sleep(time.Second)
	}
}

func testRuntineDefer() {
	fmt.Println("testRuntimeDefer()")
	go printNum(10)
	go device(10, 2)
	go device(10, 0)
	time.Sleep(time.Second * 10)
}

func printNum(maxLength int) {
	fmt.Println("printNum()")
	for i := 0; i < maxLength; i++ {
		fmt.Println("i:", i)
		time.Sleep(time.Second)
	}
}

func device(num1 int, num2 int) {
	defer deferRecover()
	fmt.Println("device()")
	result := num1 / num2
	fmt.Println("result:", result)
}

func deferRecover() {
	error := recover()
	if error != nil {
		fmt.Println("error=", error)
	}
}
