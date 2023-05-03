package main

import (
	"fmt"
	"strconv"
	"unsafe"
)

func testChangeDataType() {
	fmt.Println("testChangeDataType()")
	var int1 int = 100
	fmt.Printf("%v %T %v\n", int1, int1, unsafe.Sizeof(int1))
	var float1 float64 = float64(int1)
	fmt.Printf("%v %T %v\n", float1, float1, unsafe.Sizeof(float1))
	var float2 float64 = 129.100
	fmt.Printf("%v %T %v\n", float2, float2, unsafe.Sizeof(float2))
	var int2 int8 = int8(float2)
	fmt.Printf("%v %T %v\n", int2, int2, unsafe.Sizeof(int2))
}
func testNumOverflow() {
	fmt.Println("testNumOverflow()")
	var num1 int8 = 127
	var num2 int8 = 12
	var num3 int8 = num1 + num2
	fmt.Printf("%v %T %v\n", num3, num3, unsafe.Sizeof(num3))
	//var num4 int8 = 128 + num2
	//fmt.Printf("%v %T %v\n", num4, num4, unsafe.Sizeof(num4))
}

func testSprintf() {
	fmt.Println("testSprintf()")
	var num1 int = 100
	var num2 float64 = 100.100
	var bool3 bool = true
	var byte4 byte = 'a'

	var str1 string = fmt.Sprintf("%d", num1)
	fmt.Printf("%T %q %v\n", num1, str1, unsafe.Sizeof(str1))
	var str2 string = fmt.Sprintf("%f", num2)
	fmt.Printf("%T %q %v\n", num2, str2, unsafe.Sizeof(str2))
	var str3 string = fmt.Sprintf("%t", bool3)
	fmt.Printf("%T %q %v\n", bool3, str3, unsafe.Sizeof(str3))
	var str4 string = fmt.Sprintf("%c", byte4)
	fmt.Printf("%T %q %v\n", byte4, str4, unsafe.Sizeof(str4))
}

func testStrconv() {
	fmt.Println("testStrconv()")
	var num1 int = 100
	var num2 float64 = 100.100
	var bool3 bool = true
	var byte4 byte = 'a'

	var str1 string = strconv.FormatInt(int64(num1), 10)
	fmt.Printf("%T %q %v\n", num1, str1, unsafe.Sizeof(str1))
	var str2 string = strconv.FormatFloat(num2, 'f', -1, 64)
	fmt.Printf("%T %q %v\n", num2, str2, unsafe.Sizeof(str2))
	var str3 string = strconv.FormatBool(bool3)
	fmt.Printf("%T %q %v\n", bool3, str3, unsafe.Sizeof(str3))
	var str4 string = strconv.FormatUint(uint64(byte4), 10)
	fmt.Printf("%T %q %v\n", byte4, str4, unsafe.Sizeof(str4))

	num11, _ := strconv.ParseInt(str1, 10, 64)
	fmt.Printf("%T %v %v\n", num11, num11, unsafe.Sizeof(num11))
	num22, _ := strconv.ParseFloat(str2, 64)
	fmt.Printf("%T %v %v\n", num22, num22, unsafe.Sizeof(num22))
	bool33, _ := strconv.ParseBool(str3)
	fmt.Printf("%T %v %v\n", bool33, bool33, unsafe.Sizeof(bool33))
	byte44, _ := strconv.ParseUint(str4, 10, 8)
	fmt.Printf("%T %v %v\n", byte44, byte44, unsafe.Sizeof(byte44))

	var str5 string = strconv.Itoa(num1)
	fmt.Printf("%T %q %v\n", str5, str5, unsafe.Sizeof(str5))
	num6, err := strconv.Atoi(str5)
	if err != nil {
		fmt.Println("strconv.Atoi() error:", err)
		return
	}
	fmt.Printf("%T %v %v\n", num6, num6, unsafe.Sizeof(num6))
}
func main() {
	fmt.Println("main()")
	//testChangeDataType()
	//testNumOverflow()
	//testSprintf()
	testStrconv()
}
