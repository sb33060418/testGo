package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func testClientConn() {
	fmt.Println("testClientConn()")
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("client Dial error:", err)
		return
	}
	fmt.Printf("server accept success:%v, %p, %T\n", conn, conn, conn)
	fmt.Printf("LocalAddr:%v. RemoteAddr:%v\n", conn.LocalAddr(), conn.RemoteAddr())
}

func testClientWrite() {
	fmt.Println("testClientWrite()")
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("client Dial error:", err)
		return
	}
	fmt.Printf("server accept success:%v, %p, %T\n", conn, conn, conn)
	fmt.Printf("LocalAddr:%v. RemoteAddr:%v\n", conn.LocalAddr(), conn.RemoteAddr())

	defer conn.Close()
	reader := bufio.NewReader(os.Stdin)
	sendStr, errReader := reader.ReadString('\n')
	if errReader != nil {
		fmt.Println("read input error:", errReader)
		return
	} else {
		fmt.Println("read input success:", sendStr)
	}
	nWrite, errWrite := conn.Write([]byte(sendStr))
	if errWrite != nil {
		fmt.Println("client write error:", errWrite)
		return
	} else {
		fmt.Printf("client write success: %d bytes: %v\n", nWrite, sendStr)
	}
}

func testClientWriteRead() {
	fmt.Println("testClientWriteRead()")
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("client Dial error:", err)
		return
	}
	fmt.Printf("server accept success:%v, %p, %T\n", conn, conn, conn)
	fmt.Printf("LocalAddr:%v. RemoteAddr:%v\n", conn.LocalAddr(), conn.RemoteAddr())

	defer conn.Close()

	for {
		reader := bufio.NewReader(os.Stdin)
		sendStr, errReader := reader.ReadString('\n')
		if errReader != nil {
			fmt.Println("read input error:", errReader)
			return
		} else {
			fmt.Println("read input success:", sendStr)
		}
		if sendStr == "exit\r\n" {
			fmt.Println("client exit")
			return
		}
		nWrite, errWrite := conn.Write([]byte(sendStr))
		if errWrite != nil {
			fmt.Println("client write error:", errWrite)
			return
		} else {
			fmt.Printf("client write success: %d bytes: %v\n", nWrite, sendStr)
		}
		buf := make([]byte, 1024)
		nRead, errRead := conn.Read(buf)
		if errRead != nil {
			fmt.Println("client read error:", errRead)
			return
		}
		receviceStr := string(buf[:nRead])
		fmt.Printf("client read success: %d bytes: %v\n", nRead, receviceStr)
	}
}

func main() {
	fmt.Println("main()")
	//testClientConn()
	//testClientWrite()
	testClientWriteRead()
}
