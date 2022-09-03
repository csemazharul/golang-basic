package main

import "fmt"

func main() {

	result, message := calCulate(1,2,3)
	fmt.Println(result, message)

}

func calCulate(values ...int) (int, string) {
	var sum int
	for _, value := range values {
		sum += value
	}
	return sum,"This is calculation"

}
