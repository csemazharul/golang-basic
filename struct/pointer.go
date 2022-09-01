package main 

import "fmt"

func callByRef(pointer *int){
	*pointer = *pointer+1
}

func main(){
//  var p *int
  x:=10
	fmt.Println(x)
	callByRef(&x)
	fmt.Println(x)
//  p = &x
//  fmt.Println(*p)
//  fmt.Println(p)
//  fmt.Println(&x)
}