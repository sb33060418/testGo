package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
)

func testFile(fileName string) {
	fmt.Println("testFile")
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("file open err:", err)
	} else {
		//应该包含在else中
	}
	fmt.Printf("file: %v\n", file)
	errClose := file.Close()
	if errClose != nil {
		fmt.Println("file close err:", errClose)
	}
	fmt.Println("testFile over")
}

func testReadFile(fileName string) {
	fmt.Println("testReadFile")
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("file read err:", err)
		return
	}
	fmt.Printf("file content: %v\n", content)
	fmt.Printf("file content: %v\n", string(content))
	fmt.Println(content)
	fmt.Println(string(content))
	fmt.Println("testReadFile over")
}

func testFileReader(fileName string) {
	fmt.Println("testFileReader")
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("file open err:", err)
		return
	}
	fmt.Printf("file opened: %v\n", file)
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		fmt.Println(str)
		if err == io.EOF {
			fmt.Println("file read over:", err)
			break
		}
	}
	fmt.Println("testFileReader over")
}

func testFileWriter(fileName string) {
	fmt.Println("testFileWriter")
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("file open err:", err)
		return
	}
	fmt.Println("file opened: ", file)
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		fmt.Println(str)
		if err == io.EOF {
			break
		}
	}
	writer := bufio.NewWriter(file)
	for i := 0; i < 3; i++ {
		writer.WriteString("输出：" + strconv.Itoa(i) + "\n")
	}
	writer.Flush()

	s := os.FileMode(0666).String()
	fmt.Printf("0666 file mode:%v\n", s)
	fmt.Println("testFileWriter over")
}

func testWriteFile(fileName string) {
	fmt.Println("testWriteFile")
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("file read err:", err)
		return
	}
	fmt.Println("file readed: ", fileName)
	fmt.Println("file content:", string(content))

	errWrite := ioutil.WriteFile(fileName+".bak", content, 0666)
	if err != nil {
		fmt.Println("file write err:", errWrite)
	} else {
		fmt.Println("file writed:", fileName+".bak")
	}
	fmt.Println("testWriteFile over")

}

func main() {
	fmt.Println("main")
	//testFile("D:/work/file/test.txt")
	//testFile("D:/work/file/test2.txt")
	//testReadFile("D:/work/file/test.txt")
	//testReadFile("D:/work/file/test2.txt")
	//testFileReader("D:/work/file/test.txt")
	//testFileReader("D:/work/file/test2.txt")
	//testFileWriter("D:/work/file/test.txt")
	testWriteFile("D:/work/file/test.txt")
	testWriteFile("D:/work/file/test2.txt")

	fmt.Println("main over")
}
