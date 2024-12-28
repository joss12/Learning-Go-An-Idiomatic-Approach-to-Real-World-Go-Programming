package main

import "fmt"

var pl = fmt.Println

var myfuncVariable func(string) int

func f1(a string) int {
	return len(a)
}
func f2(a string) int {
	total := 0
	for _, v := range a {
		total += int(v)
	}
	return total
}

func main() {

	var myfuncVariable func(string) int
	myfuncVariable = f1
	result := myfuncVariable("hello")
	pl(result)

	myfuncVariable = f2
	result = myfuncVariable("hello")
	pl(result)
}
