package main

import "fmt"

var pl = fmt.Println

type person struct{
	age int
	name string
}

func modifyFails(i int, s string, p person){
	i = i * 2
	s = "Goodbye"
	p.name = "Eddy"
}

func main(){
	p := person{}
	i := 2
	s := "hello"
	modifyFails(i, s, p)
	pl(i, s, p)
}