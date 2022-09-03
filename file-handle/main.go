package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	//defer key work reverse the order of execution
	content := "Hello Bangladesh"
	file, er := os.Create("./test.txt")

	if er != nil {
		panic(er)
	}

	length, er := io.WriteString(file, content)

	if er != nil {
		panic(er)
	}
	fmt. Println("Total length: ", length)

	defer file.Close()
	readFile("./test.txt")
}

func readFile(fileName string)  {
	dataByte, er := ioutil.ReadFile(fileName)
	if er != nil {
		panic(er)
	}
	fmt.Println(string(dataByte))

}