package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("main()")
	testTime()
}

func testTime() {
	fmt.Println("testTime()")
	now := time.Now()
	fmt.Println(now)
	fmt.Printf("%T %v\n", now, now)
	fmt.Printf("%v年%v月%v日 %v时%v分%v秒\n", now.Year(), int(now.Month()), now.Day(), now.Hour(), now.Minute(), now.Second())
	fmt.Printf("%T %v %s %d\n", now.Month(), now.Month(), now.Month(), now.Month())
	dateStr := fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	fmt.Println(dateStr)
	fmt.Println(now.Unix())
	fmt.Println(now.UnixNano())
	fmt.Println(time.Unix(now.Unix(), 0))
	fmt.Println(time.Unix(0, now.UnixNano()))
	fmt.Println(now.Format("2006-01-02 15:04:05"))
	fmt.Println(now.Format("2006-01-02 15:04:05.000"))
	fmt.Println(now.Format("2006-01-02 15:04:05.000000"))

	// 日期计算
	fmt.Println(now.Add(time.Hour))
	fmt.Println(now.AddDate(1, 1, 1))
	fmt.Println(now.AddDate(0, 0, -1))
	fmt.Println(now.AddDate(1, -2, 20))
	fmt.Println(now.AddDate(1, -2, 21))
	fmt.Println(now.AddDate(1, -2, 21).AddDate(-1, 0, 0))

}
