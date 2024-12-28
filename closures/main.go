package main

import "fmt"

var pl = fmt.Println

func main() {
	a := 20
	f := func() {
		pl(a)
		a = 30
		pl(a)
	}
	f()
	pl(a)
}
