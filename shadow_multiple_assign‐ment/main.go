package main

import "fmt"

var pl = fmt.Println

func main(){
	x := 10
	if x >5{
		x,y := 5, 20
		pl(x, y)
	}
	pl(x)
}