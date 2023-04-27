package array2

import (
	"fmt"
	"unsafe"
)

func TestArray2() {
	fmt.Println("array2.TestArray2()")
	testArray2Loop()
	//testDefineArray2()
	//testArray2SetValue()
	//testArray2Addr()

}
func testArray2Addr() {
	fmt.Println("testArray2Addr()")
	var arr [2][3]int
	fmt.Printf("arr=%v, type=%T, len=%v, size=%v, addr=%p\n", arr, arr, len(arr), unsafe.Sizeof(arr), &arr)
	fmt.Printf("arr[0]=%v, type=%T, len=%v, size=%v, addr=%p\n", arr[0], arr[0], len(arr[0]), unsafe.Sizeof(arr[0]), &arr[0])
	fmt.Printf("arr[0][0]=%v, type=%T, size=%v, addr=%p\n", arr[0][0], arr[0][0], unsafe.Sizeof(arr[0][0]), &arr[0][0])
	fmt.Println(&arr[0][0])
	fmt.Println(&arr[0][1])
	fmt.Println(&arr[0][2])
	fmt.Println(&arr[1][0])
}
func testArray2SetValue() {
	fmt.Println("testArray2SetValue()")
	var arr [2][3]int
	fmt.Println(arr)
	arr[0][0] = 10
	*(&arr[0][1]) = 20
	fmt.Println(arr)
}

func testDefineArray2() {
	fmt.Println("testDefineArray2()")
	var arr1 [2][3]int = [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(arr1)
	var arr2 = [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(arr2)
	var arr3 = [...][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(arr3)
	var arr4 = [...][3]int{1: {4, 5, 6}}
	fmt.Println(arr4)
}

func testArray2Loop() {
	fmt.Println("testArray2Loop()")
	var arr = [...][3]int{{1, 2, 3}, {4, 5, 6}}
	//遍历方式1
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
		for j := 0; j < len(arr[i]); j++ {
			fmt.Println(arr[i][j])
		}
	}
	//遍历方式2
	for i, v := range arr {
		fmt.Println(i, v)
		for j, vv := range v {
			fmt.Println(i, j, vv)
		}
	}
}
