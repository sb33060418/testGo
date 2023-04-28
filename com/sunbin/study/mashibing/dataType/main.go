package main

import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Println("main()")
	testBool()
	testByte()
	testInt()
	testFloat()
}

func testBool() {
	fmt.Println("testBool()")
	var b bool = true
	fmt.Printf("%v %T %v\n", b, b, unsafe.Sizeof(b))
	bDefault := false
	fmt.Printf("%v %T %v\n", bDefault, bDefault, unsafe.Sizeof(bDefault))
}

func testByte() {
	fmt.Println("testByte()")
	var b byte = 'a'
	fmt.Printf("%v %T %v\n", b, b, unsafe.Sizeof(b))
	bDefault := 'b'
	fmt.Printf("%v %T %v\n", bDefault, bDefault, unsafe.Sizeof(bDefault))
}

func testInt() {
	fmt.Println("testInt()")
	var i int = 9
	// unsafe.Sizeof() returns the size of the variable in bytes
	fmt.Printf("%v %T %v\n", i, i, unsafe.Sizeof(i))
	var i8 int8 = 8
	fmt.Printf("%v %T %v\n", i8, i8, unsafe.Sizeof(i8))
	var i16 int16 = 16
	fmt.Printf("%v %T %v\n", i16, i16, unsafe.Sizeof(i16))
	var i32 int32 = 32
	fmt.Printf("%v %T %v\n", i32, i32, unsafe.Sizeof(i32))
	var i64 int64 = 64
	fmt.Printf("%v %T %v\n", i64, i64, unsafe.Sizeof(i64))
	iDefault := 100
	fmt.Printf("%v %T %v\n", iDefault, iDefault, unsafe.Sizeof(iDefault))

}

func testFloat() {
	fmt.Println("testFloat()")
	var f32 float32 = 32.32
	fmt.Printf("%v %T %v\n", f32, f32, unsafe.Sizeof(f32))
	var f64 float64 = 64.64
	fmt.Printf("%v %T %v\n", f64, f64, unsafe.Sizeof(f64))
	fDefault := 100.100
	fmt.Printf("%v %T %v\n", fDefault, fDefault, unsafe.Sizeof(fDefault))
}
