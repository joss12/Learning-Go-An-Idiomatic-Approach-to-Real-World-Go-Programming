package main

import "fmt"

var pl = fmt.Println

var (
	add = func(i, j int)int{return i + j}
	sub = func(i, j int)int{return i - j}
	mul = func(i, j int)int{ return i * j}
	div = func(i, j int)int{return i / j}
)


func main(){
	// f := func(j int){
	// 	pl("Printing", j, "from inside of the anonymous function")
	// }
	// for i := 0; i < 5; i++{
	// 	f(i)
	// }

	// for i := 0; i < 5; i++{
	// 	func (j int){
	// 		pl("Printing", j, "from inside of the anonymous function")
	// 	}(i)
	// }

	x := add(2, 3)
	pl(x)
	changeAdd()
	y := add(2, 3)
	pl(y)
	a := sub(2, 3)
	pl(a)
}

func changeAdd(){
	add = func(i, j int)int{return i + j + j}
}