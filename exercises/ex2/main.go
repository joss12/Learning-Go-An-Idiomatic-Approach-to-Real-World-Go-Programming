package main

import (
	"fmt"
	"math/rand"
)

var pl = fmt.Println

func main() {
	mySlice := make([]int, 0, 100)
	for i := 0; i < 100; i++{
		mySlice = append(mySlice, rand.Intn(100))
	}
	for _, v := range mySlice{
		switch {
		case v%6 == 6:
			pl("Six!")
		case v%2 == 0:
			pl("Two!")
		case v%3 == 0:
			pl("Three!")
		default:
			pl("Never mind")
		}
	}
}
