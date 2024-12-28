package main

import "fmt"

var pl = fmt.Println

type Person struct{
	FirstName string
	LastName string
	Age int
}

func MakePerson(firstName, lastName string, age int)Person{
	return Person{
		FirstName: firstName,
		LastName: lastName,
		Age: age,
	}
}

func MakePersonPointer(firstName, lastName string, age int) *Person{
	return &Person{
		FirstName: firstName,
		LastName: lastName,
		Age: age,
	}
}

func main(){
	p := MakePerson("Fred", "Williamson", 25)
	pl(p)
	p2 := MakePersonPointer("Grace","Mouity", 25)
	pl(p2)
}