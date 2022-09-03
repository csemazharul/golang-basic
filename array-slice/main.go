package main

import (
	"fmt"
	"sort"
)

func main() {
	var froutList = []string{"apple", "tomato", "peach", "mango", "banana"}
	// froutList = append(froutList, "grape")
	froutList = froutList[2:4]
	fmt.Print(len(froutList))

	marks := make([]int, 4)
	marks[0] = 10
	marks[1] = 20
	marks[2] = 30
	marks[3] = 40

	marks= append(marks, 50)
	
	sort.IntsAreSorted(marks)
	fmt.Println(marks)
}
