package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

func test(count int) {
	for i := 1; i <= count; i++ {
		fmt.Println("inner: " + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}

// 1.顺序执行
func test1() {
	fmt.Println("test1()")
	test(5)
	for i := 1; i <= 5; i++ {
		fmt.Println("outer: " + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}

// 2.协程
func test2() {
	fmt.Println("test2()")
	go test(5)
	for i := 1; i <= 5; i++ {
		fmt.Println("outer: " + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}

// 3.主死从随
func test3() {
	fmt.Println("test3()")
	go test(7)
	for i := 1; i <= 5; i++ {
		fmt.Println("outer: " + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}

// 4.启动多个协程错误，闭包
func test4() {
	fmt.Println("test4()")
	for i := 1; i <= 5; i++ {
		go func() {
			fmt.Println("inner: " + strconv.Itoa(i))
		}()
	}
	time.Sleep(time.Second * 2)

}

// 5.启动多个协程
func test5() {
	fmt.Println("test5()")
	for i := 1; i <= 5; i++ {
		go func(n int) {
			fmt.Println("inner: " + strconv.Itoa(n))
		}(i)
	}
	time.Sleep(time.Second * 2)

}

// 6.WaitGroup解决主死从随
func test6() {
	fmt.Println("test6()")
	var wg sync.WaitGroup
	//wg.Add(5)
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func(n int) {
			fmt.Println("inner: " + strconv.Itoa(n))
			time.Sleep(time.Second)
			fmt.Println("inner: " + strconv.Itoa(n) + " end")
			defer wg.Done()
		}(i)
	}
	fmt.Println("outer wait")
	wg.Wait()
	fmt.Println("outer end")
}

var totalNum int
var wg sync.WaitGroup

// 7.测试多协程并发访问外部变量
func test7() {
	fmt.Println("test7()")
	wg.Add(2)
	go add()
	go sub()
	wg.Wait()
	fmt.Println(totalNum)
}
func add() {
	defer wg.Done()
	for i := 1; i <= 100000; i++ {
		totalNum = totalNum + 1
	}
}
func sub() {
	defer wg.Done()
	for i := 1; i <= 100000; i++ {
		totalNum = totalNum - 1
	}
}

// 8.互斥锁
var lock sync.Mutex

func testMutex() {
	totalNum = 0
	fmt.Println("testMutex()")
	wg.Add(2)
	go addWithMutex()
	go subWithMutex()
	wg.Wait()
	fmt.Println(totalNum)
}
func addWithMutex() {
	defer wg.Done()
	for i := 1; i <= 10000000; i++ {
		lock.Lock()
		totalNum = totalNum + 1
		lock.Unlock()
	}
}
func subWithMutex() {
	defer wg.Done()
	for i := 1; i <= 10000000; i++ {
		lock.Lock()
		totalNum = totalNum - 1
		lock.Unlock()
	}
}

// 9.读写锁
var rwLock sync.RWMutex

func testRWMutex() {
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go readWithRWMutex()
	}
	wg.Add(1)
	go writeWithRWMutex()
	wg.Wait()
}
func readWithRWMutex() {
	defer wg.Done()
	rwLock.RLock()
	fmt.Println("read start")
	time.Sleep(time.Second)
	fmt.Println("read end")
	rwLock.RUnlock()
}

func writeWithRWMutex() {
	defer wg.Done()
	rwLock.Lock()
	fmt.Println("write start")
	time.Sleep(time.Second * 5)
	fmt.Println("write end")
	rwLock.Unlock()
}

func main() {
	fmt.Println("main()")
	//test1()
	//test2()
	//test3()
	//test4()
	//test5()
	//test6()
	//fmt.Println(time.Now())
	//test7()
	//fmt.Println(time.Now())
	//testMutex()
	//fmt.Println(time.Now())
	testRWMutex()
}
