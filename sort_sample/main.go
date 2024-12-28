package main

import (
	"fmt"
	"sort"
)

var pl = fmt.Println

type Person struct{
	FirstName string
	LastName string
	Age int
}



func main(){
	people :=[]Person{
		{"Eddy", "Mouity", 32},
		{"Meyis", "Mouity", 5},
		{"Grace", "Mouity", 10},
	}
	// pl(people)

	
	sort.Slice(people,  func(i, j int)bool{
		return people[i].LastName < people[j].LastName
	})
	pl(people)
	sort.Slice(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
		})
		fmt.Println(people)
}