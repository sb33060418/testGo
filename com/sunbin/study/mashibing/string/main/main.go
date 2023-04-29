package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("main()")
	testAtoi()
	testStrings()
}

func testAtoi() {
	fmt.Println("testAtoi()")
	num, err := strconv.Atoi("123")
	fmt.Println(num, err)
	str := strconv.Itoa(123)
	fmt.Println(str)
	num1, err1 := strconv.Atoi("a")
	fmt.Printf("num1=%d, err1=%v, type=%T\n", num1, err1, err1)
}

func testStrings() {
	fmt.Println("testString()")
	str := "我爱中国"
	fmt.Println(strings.Contains(str, "国"))
	fmt.Println(strings.Count(str, "国"))
	//不区分大小写比较
	fmt.Println(strings.EqualFold("aBcD中国", "AbCd中国"))
	//区分大小写比较
	fmt.Println("aBcD中国" == "AbCd中国")
	//strings.Index字节索引，中文占3个字节
	fmt.Println(strings.Index(str, "国"))
	fmt.Println(strings.Replace(str, "中国", "China", -1))
	fmt.Println(strings.Replace("aaa", "a", "b", 2))
	arr := strings.Split("1,2,3,4", ",")
	fmt.Println(arr)
	for i := 0; i < len(arr); i++ {
		fmt.Println(i, arr[i])
	}
	for i, v := range arr {
		fmt.Println(i, v)
	}
	fmt.Println(strings.ToLower("ABC中国"))
	fmt.Println(strings.ToUpper("abc中国"))
	fmt.Println(strings.TrimSpace("  abc中国  "))
	fmt.Println(strings.Trim(" 1 abc中国  1 ~ ", " ~1"))
	fmt.Println(strings.TrimLeft(" 1 abc中国  1 ~ ", " ~1"))
	fmt.Println(strings.TrimRight(" 1 abc中国  1 ~ ", " ~1"))
	fmt.Println(strings.Repeat("abc中国", 3))
	fmt.Println(strings.Fields(" 1 abc中国  1 ~ "))
	fmt.Println(strings.HasPrefix("abc中国", "ab"))
	fmt.Println(strings.HasSuffix("abc中国", "国"))
}
