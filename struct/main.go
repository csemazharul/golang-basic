package main 

import "fmt"

type Student struct {
	id int
	name string
	age int
}


func main(){
	var student Student
	student.id = 1
	student.name = "Arif Uddin"
	student.age = 20
	// fmt.Println(student.name)
	student.getEmail()
}

func (student Student) getEmail(){
	fmt.Println(student)

}