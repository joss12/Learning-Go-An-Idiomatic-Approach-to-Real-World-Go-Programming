package main

import "fmt"

var pl = fmt.Println

func main(){
	x := []string{"a", "b", "c", "d"}
	y := x [:2]
	pl(cap(x), cap(y))
	y = append(y, "z")
	pl("x:", x)
	pl("y:", y)
}