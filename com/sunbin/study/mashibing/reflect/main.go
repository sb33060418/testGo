package main

import (
	"fmt"
	"reflect"
)

func testReflect() {
	fmt.Println("testReflect()")
	var num int = 100
	fmt.Printf("value: %v, type: %T\n", num, num)
	reflectType := reflect.TypeOf(num)
	fmt.Printf("reflectType: %v, type: %T\n", reflectType, reflectType)
	reflecttValue := reflect.ValueOf(num)
	fmt.Printf("reflectTvalue: %v, type: %T\n", reflecttValue, reflecttValue)
	//num2 := 80 + reflectTvalue
	num2 := 80 + reflecttValue.Int()
	fmt.Println("num2: ", num2)
	i := reflecttValue.Interface()
	fmt.Printf("i: %v, type: %T\n", i, i)
	num3 := i.(int)
	fmt.Printf("num3: %v, type: %T\n", num3, num3)
}

type Student struct {
	Name string
	Age  int
}

func testReflectStruct() {
	fmt.Println("testReflectStruct()")
	student := Student{
		Name: "张三",
		Age:  18,
	}
	fmt.Printf("value: %v, type: %T\n", student, student)
	reflectType := reflect.TypeOf(student)
	fmt.Printf("reflectType: %v, type: %T\n", reflectType, reflectType)
	reflectValue := reflect.ValueOf(student)
	fmt.Printf("reflectValue: %v, type: %T\n", reflectValue, reflectValue)
	reflectType2 := reflect.TypeOf(&student)
	fmt.Printf("reflectType2: %v, type: %T\n", reflectType2, reflectType2)
	reflectValue2 := reflect.ValueOf(&student)
	fmt.Printf("reflectValue2: %v, type: %T\n", reflectValue2, reflectValue2)
	i := reflectValue.Interface()
	student2, flag := i.(Student)
	if flag {
		fmt.Printf("student2: %v, type: %T\n", student2, student2)
	} else {
		fmt.Println("类型断言失败")
	}
}

func testReflectKind() {
	fmt.Println("testReflectKind()")
	var num int = 100
	fmt.Printf("value: %v, type: %T\n", num, num)
	reflectType := reflect.TypeOf(num)
	fmt.Printf("reflectType: %v, type: %T\n", reflectType, reflectType)
	kind := reflectType.Kind()
	fmt.Printf("kind: %v, type: %T\n", kind, kind)
	reflectValue := reflect.ValueOf(num)
	fmt.Printf("reflectValue: %v, type: %T\n", reflectValue, reflectValue)
	kind2 := reflectValue.Kind()
	fmt.Printf("kind2: %v, type: %T\n", kind2, kind2)
	i := reflectValue.Interface()
	fmt.Printf("i: %v, type: %T\n", i, i)
	num2, flag := i.(int)
	if flag {
		fmt.Printf("num2: %v, type: %T\n", num2, num2)
	} else {
		fmt.Println("类型断言失败")
	}

	fmt.Println("--------------------")
	student := Student{
		Name: "张三",
		Age:  18,
	}
	fmt.Printf("value: %v, type: %T\n", student, student)
	reflectType2 := reflect.TypeOf(student)
	fmt.Printf("reflectType: %v, type: %T\n", reflectType2, reflectType2)
	kind3 := reflectType2.Kind()
	fmt.Printf("kind3: %v, type: %T\n", kind3, kind3)
	reflectValue2 := reflect.ValueOf(student)
	fmt.Printf("reflectValue: %v, type: %T\n", reflectValue2, reflectValue2)
	kind4 := reflectValue2.Kind()
	fmt.Printf("kind4: %v, type: %T\n", kind4, kind4)
	i2 := reflectValue2.Interface()
	fmt.Printf("i2: %v, type: %T\n", i2, i2)
	student2, flag2 := i2.(Student)
	if flag2 {
		fmt.Printf("student2: %v, type: %T\n", student2, student2)
	} else {
		fmt.Println("类型断言失败")
	}

}
func testReflectSet() {
	fmt.Println("testReflectSet()")
	var num int = 100
	fmt.Println("num: ", num)
	reflectValue := reflect.ValueOf(&num)
	fmt.Printf("reflectValue: %v, type: %T\n", reflectValue, reflectValue)
	kind := reflectValue.Kind()
	fmt.Printf("kind: %v, type: %T\n", kind, kind)
	reflectValue.Elem().SetInt(200)
	fmt.Println("num: ", num)
}

func (student Student) Get() (string, int) {
	fmt.Println("Student.Get(), Name: ", student.Name, ", Age: ", student.Age)
	return student.Name, student.Age
}

func (student Student) Print() {
	fmt.Println("Student.Print(): ", student.Name, ", ", student.Age)
}
func (student Student) Set(name string, age int) {
	fmt.Println("Student.Set() before, Name: ", student.Name, ", Age: ", student.Age)
	student.Name = name
	student.Age = age
	fmt.Println("Student.Set() after, Name: ", student.Name, ", Age: ", student.Age)
}
func (student *Student) SetPointer(name string, age int) {
	fmt.Println("Student.SetPointer() before, Name: ", student.Name, ", Age: ", student.Age)
	student.Name = name
	student.Age = age
	fmt.Println("Student.SetPointer() after, Name: ", student.Name, ", Age: ", student.Age)
}

func testReflectStructSet() {
	fmt.Println("testReflectStructSet()")
	student := Student{
		Name: "张三",
		Age:  18,
	}
	fmt.Printf("value: %v, type: %T\n", student, student)
	reflectValue := reflect.ValueOf(student)
	fmt.Printf("reflectValue: %v, type: %T\n", reflectValue, reflectValue)
	// 字段
	numField := reflectValue.NumField()
	fmt.Printf("numField: %v, type: %T\n", numField, numField)
	for i := 0; i < numField; i++ {
		field := reflectValue.Field(i)
		fmt.Printf("field[%v]: %v, type: %T\n", i, field, field)
	}
	////反射结构体来修改字段失败，必须反射指针
	//reflectValue.Elem().Field(0).SetString("李四")
	//reflectValue.Elem().Field(1).SetInt(20)
	//fmt.Printf("reflectValue: %v, type: %T\n", reflectValue, reflectValue)
	//fmt.Printf("value: %v, type: %T\n", student, student)
	//方法
	numMethod := reflectValue.NumMethod()
	fmt.Printf("numMethod: %v, type: %T\n", numMethod, numMethod)
	result := reflectValue.Method(0).Call(nil)
	fmt.Printf("reflectValue.Method(0): %v, %v\n", result[0].String(), result[1].Int())
	fmt.Printf("reflectValue.Method(1): %v\n", reflectValue.Method(1).Call(nil))
	// 两种定义方法
	//params := make([]reflect.Value, 2)
	//params[0] = reflect.ValueOf("李四")
	//params[1] = reflect.ValueOf(20)
	var params []reflect.Value
	params = append(params, reflect.ValueOf("李四"))
	params = append(params, reflect.ValueOf(20))
	fmt.Println("params: ", params)
	fmt.Printf("reflect struct reflectValue.Method(2) Set(): %v\n", reflectValue.Method(2).Call(params))
	fmt.Printf("reflect struct reflectValue: %v, type: %T\n", reflectValue, reflectValue)
	fmt.Printf("reflect struct value: %v, type: %T\n", student, student)
	////反射结构体只能获取到结构体绑定方法，不能获取指针绑定方法
	//fmt.Println("params: ", params)
	//fmt.Printf("reflect struct reflectValue.Method(3) SetPointer(): %v\n", reflectValue.Method(3).Call(params))
	//fmt.Printf("reflect struct reflectValue: %v, type: %T\n", reflectValue, reflectValue)
	//fmt.Printf("reflect struct value: %v, type: %T\n", student, student)
	//反射指针来修改
	reflectValue2 := reflect.ValueOf(&student)
	fmt.Printf("reflectValue2: %v, type: %T\n", reflectValue2, reflectValue2)
	reflectValue2.Elem().Field(0).SetString("王五")
	reflectValue2.Elem().Field(1).SetInt(30)
	fmt.Printf("reflect pointer reflectValue2: %v, type: %T\n", reflectValue2, reflectValue2)
	fmt.Printf("reflect pointer value: %v, type: %T\n", student, student)
	numMethod2 := reflectValue2.NumMethod()
	fmt.Printf("reflect pointer numMethod2: %v, type: %T\n", numMethod2, numMethod2)
	params[0] = reflect.ValueOf("李四")
	params[1] = reflect.ValueOf(20)
	fmt.Println("params: ", params)
	fmt.Printf("reflect pointer reflectValue2.Method(2) Set(): %v\n", reflectValue2.Method(2).Call(params))
	fmt.Printf("reflect pointer reflectValue2: %v, type: %T\n", reflectValue2, reflectValue2)
	fmt.Printf("reflect pointer value: %v, type: %T\n", student, student)
	fmt.Println("params: ", params)
	fmt.Printf("reflect pointer reflectValue2.Method(3) SetPointer(): %v\n", reflectValue2.Method(3).Call(params))
	fmt.Printf("reflect pointer reflectValue2: %v, type: %T\n", reflectValue2, reflectValue2)
	fmt.Printf("reflect pointer value: %v, type: %T\n", student, student)

}
func main() {
	fmt.Println("main()")
	//testReflect()
	//testReflectStruct()
	//testReflectKind()
	//testReflectSet()
	testReflectStructSet()
}
