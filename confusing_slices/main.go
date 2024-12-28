package main

import "fmt"

var pl = fmt.Println

func main() {
	x := make([]string, 0, 5)
	x = append(x, "a", "b", "c", "d")
	y := x[:2]
	z := x[2:]
	pl(cap(x), cap(y), cap(z))
	y = append(y, "i", "j", "k")
	x = append(x, "y")
	z = append(z, "y")
	pl("x:", x)
	pl("y:", y)
	pl("z:", z)
}