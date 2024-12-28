package main

import (
	"fmt"
	"math/cmplx"
)

var pl = fmt.Println

func main() {
	var x int = 10
	var b byte = 100
	var sum3 int = x + int(b)
	var sum4 byte = byte(x) + b
	pl(sum3, sum4)
}
