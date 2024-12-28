package main

import "fmt"

var pl = fmt.Println

type MyInt int

func main() {
	var i any
	var mine MyInt = 20
	i = mine
	// i2 := i.(MyInt)
	i2, ok := i.(int)
	// i2 := i.(string)
	pl(i2)
	if !ok{
		return
	}
	// pl(i2 + 1)
	pl(i2 + 1)
}
