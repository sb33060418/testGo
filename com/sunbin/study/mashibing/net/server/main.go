package main

import (
	"fmt"
	"net"
)

func testServerConn() {
	fmt.Println("testServerConn()")
	listener, errListen := net.Listen("tcp", "127.0.0.1:8888")
	if errListen != nil {
		fmt.Println("server listen error:", errListen)
		return
	}
	fmt.Printf("server listen success:%v, %p, %T\n", listener, listener, listener)
	for {
		conn, errAccept := listener.Accept()
		if errAccept != nil {
			fmt.Println("server accept error:", errAccept)
		} else {
			fmt.Printf("server accept success:%v, %p, %T\n", conn, conn, conn)
			fmt.Printf("LocalAddr:%v. RemoteAddr:%v\n", conn.LocalAddr(), conn.RemoteAddr())
		}
	}
}
func testServerRead() {
	fmt.Println("testServerRead()")
	listener, errListen := net.Listen("tcp", "127.0.0.1:8888")
	if errListen != nil {
		fmt.Println("server listen error:", errListen)
		return
	}
	fmt.Printf("server listen success:%v, %p, %T\n", listener, listener, listener)
	for {
		conn, errAccept := listener.Accept()
		if errAccept != nil {
			fmt.Println("server accept error:", errAccept)
		} else {
			fmt.Printf("server accept success:%v, %p, %T\n", conn, conn, conn)
			fmt.Printf("LocalAddr:%v. RemoteAddr:%v\n", conn.LocalAddr(), conn.RemoteAddr())
		}
		go processRead(conn)
	}
}

func processRead(conn net.Conn) {
	defer conn.Close()
	for {
		buf := make([]byte, 1024)
		nRead, errRead := conn.Read(buf)
		if errRead != nil {
			fmt.Println("server read error:", errRead)
			return
		}
		receviceStr := string(buf[:nRead])
		fmt.Printf("server read success: %d bytes: %v\n", nRead, receviceStr)
	}
}

var flag bool = true

func testServerReadWrite() {
	fmt.Println("testServerReadWrite()")
	listener, errListen := net.Listen("tcp", "127.0.0.1:8888")
	if errListen != nil {
		fmt.Println("server listen error:", errListen)
		return
	}
	fmt.Printf("server listen success:%v, %p, %T\n", listener, listener, listener)
	for flag {
		conn, errAccept := listener.Accept()
		if errAccept != nil {
			fmt.Println("server accept error:", errAccept)
		} else {
			fmt.Printf("server accept success:%v, %p, %T\n", conn, conn, conn)
			fmt.Printf("LocalAddr:%v. RemoteAddr:%v\n", conn.LocalAddr(), conn.RemoteAddr())
		}
		go processReadWrite(conn, listener)
	}
}

func processReadWrite(conn net.Conn, listener net.Listener) {
	defer conn.Close()
	for {
		buf := make([]byte, 1024)
		nRead, errRead := conn.Read(buf)
		if errRead != nil {
			fmt.Println("server read error:", errRead)
			return
		}
		receviceStr := string(buf[:nRead])
		fmt.Printf("server read success: %d bytes: %v\n", nRead, receviceStr)
		var sendStr string
		if receviceStr == "bye\r\n" {
			fmt.Println("server bye")
			return
		} else if receviceStr == "shutdown\r\n" {
			fmt.Println("server shutdown")
			flag = false
			defer listener.Close()
			return
		}
		sendStr = "server response:" + receviceStr
		nWrite, errWrite := conn.Write([]byte(sendStr))
		if errWrite != nil {
			fmt.Println("server write error:", errWrite)
			return
		} else {
			fmt.Printf("server write success: %d bytes: %v\n", nWrite, sendStr)
		}
	}

}

func main() {
	fmt.Println("main()")
	//testServerConn()
	//testServerRead()
	testServerReadWrite()
}
