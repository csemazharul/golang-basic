package main

import (
	"bufio"
	"os"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	println("Please enter your opinion")
	input, _ := reader.ReadString('\n')
	println("Thanks For Opinion", input)

}
