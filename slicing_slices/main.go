package main

import "fmt"

var pl = fmt.Println

func main() {
	x := []string{"a", "b", "c"}
	y := x[:2]
	z := x[1:]
	x[1] = "y"
	y[0] = "x"
	z[1] = "z"
	pl("x:", x)
	pl("y:", y)
	pl("z:", z)

}
