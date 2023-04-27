package main

import (
	"TestGo/com/sunbin/study/mashibing/testArray/array2"
	"fmt"
)

func main() {
	fmt.Println("main()")
	array2.TestArray2()
	//testCallArr()
	//testArrType()
	//testDefineArr()
	//testScan()
	//testArray()
	//testMemory()
}

func testArray() {
	fmt.Println("testArray()")
	//定义数组
	var scores [5]int
	scores[0] = 100
	scores[1] = 90
	scores[2] = 80
	scores[3] = 70
	scores[4] = 60
	fmt.Println(scores)
	fmt.Printf("len(scores)=%v\n", len(scores))

	sum := 0
	//遍历数组1
	for i := 0; i < len(scores); i++ {
		fmt.Printf("i=%v v=%v\n", i, scores[i])
		sum += scores[i]
	}
	//遍历数组2
	sum = 0
	for i, v := range scores {
		fmt.Printf("i=%v v=%v\n", i, v)
		sum += v
	}

	//求平均值
	avg := sum / len(scores)
	fmt.Printf("sum=%v avg=%v\n", sum, avg)
}

func testMemory() {
	fmt.Println("testMemory()")
	var arr [3]int16 = [3]int16{1, 2, 3}
	fmt.Printf("arr=%v\n", arr)
	fmt.Printf("arr地址：%p\n", &arr)
	fmt.Printf("arr[0]=%v\n", arr[0])
	fmt.Printf("arr[0]地址：%p\n", &arr[0])
	fmt.Printf("arr[1]=%v\n", arr[1])
	fmt.Printf("arr[1]地址：%p\n", &arr[1])
	fmt.Printf("arr[2]=%v\n", arr[2])
	fmt.Printf("arr[2]地址：%p\n", &arr[2])
	p := &arr[0]
	fmt.Printf("p=%v\n", p)
	fmt.Printf("p地址：%p\n", &p)
	fmt.Printf("p指向的值：%v\n", *p)
	fmt.Printf("p指向的值地址：%p\n", &(*p))
	*p = 100
	fmt.Printf("p指向的值：%v\n", *p)
	fmt.Printf("arr=%v\n", arr)

}

func testScan() {
	fmt.Println("testScan()")
	var arr [5]int
	for i := 0; i < len(arr); i++ {
		fmt.Printf("请输入第%v个数：", i+1)
		fmt.Scanln(&arr[i])
	}
	fmt.Printf("arr=%v\n", arr)
}

func testDefineArr() {
	fmt.Println("testDefineArr()")
	//定义数组1
	var scores1 [5]int = [5]int{90, 80, 70, 60, 50}
	fmt.Printf("scores1=%v\n", scores1)
	//定义数组2
	var scores2 = [5]int{90, 80, 70, 60, 50}
	fmt.Printf("scores2=%v\n", scores2)
	//定义数组3
	var score3 = [...]int{80, 70, 60, 50, 40}
	fmt.Printf("score3=%v\n", score3)
	//定义数组4
	var score4 = [...]int{2: 22, 4: 44, 6: 66}
	fmt.Printf("score4=%v\n", score4)
}

func testArrType() {
	fmt.Println("testArrType()")
	var arr = [3]int{1, 4, 7}
	fmt.Printf("arr=%v,type=%T\n", arr, arr)
	testArrTypeInt3(arr)
	//testArrTypeInt5(arr) //cannot use arr (type [3]int) as type [5]int in argument to testArrTypeInt5
	//testArrTypeInt(arr) //可以使用切片
	testArrTypeInt(arr[:])
}

func testArrTypeInt(arr []int) {
	fmt.Println("testArrTypeInt()")
	fmt.Printf("arr=%v,type=%T\n", arr, arr)
}

func testArrTypeInt3(arr [3]int) {
	fmt.Println("testArrTypeInt3()")
	fmt.Printf("arr=%v,type=%T\n", arr, arr)
}

func testArrTypeInt5(arr [5]int) {
	fmt.Println("testArrTypeInt5()")
	fmt.Printf("arr=%v,type=%T\n", arr, arr)
}

func testCallArr() {
	fmt.Println("testCallArr()")
	var arr = [3]int{1, 4, 7}
	fmt.Printf("testCallArr before arr=%v\n", arr)
	testChangeArr(arr)
	fmt.Printf("testCallArr after arr=%v\n", arr)
	fmt.Printf("%T %v %v\n", arr, &arr[0], arr)
	testChangeArrByPointer(&arr)
	fmt.Printf("testCallArr after arr=%v\n", arr)
}

func testChangeArr(arr [3]int) {
	fmt.Println("testChangeArr()")
	fmt.Printf("testChangeArr before arr=%v\n", arr)
	arr[0] = 100
	fmt.Printf("testChangeArr after arr=%v\n", arr)
}

func testChangeArrByPointer(arr *[3]int) {
	fmt.Println("testChangeArrByPointer()")
	fmt.Printf("%T %v %v %T %v %v\n", arr, &arr, arr, *arr, &((*arr)[0]), *arr)
	fmt.Printf("testChangeArrByPointer before arr=%v\n", arr)
	(*arr)[0] = 100
	fmt.Printf("testChangeArrByPointer after arr=%v\n", arr)
}
